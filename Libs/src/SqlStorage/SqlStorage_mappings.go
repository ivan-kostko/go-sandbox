// The file contains SqlStorage mapping functionality
package SqlStorage

import (
	. "customErrors"
	"encoding/json"
	"reflect"
	"strings"
)

const (
	EMPTY_STRING = ""
)

const (
	ERR_NONSTRUCTTYPE    = "Wont generate structure mapping for non-struct type"
	ERR_FAILEDTOGENMAP   = "Failed to generate mapping"
	ERR_FAILEDTOPARSETAG = "Failed to parse JSON value at field tag"
	ERR_FIELDPTROUTOFRNG = "sampleFields pointer is out of sample range"
	ERR_KEYTYPENOTMATCH  = "Key type is differ from provided"
	ERR_NONPTRARGUMENT   = "The acceptor is not a pointer"
)

// Represents subset of structure fields
type Key struct {
	Name        string
	Type        reflect.Type
	fieldsIds   []int
	fieldsNames []string
}

// Creates new FieldsSubset based on the sample pointer and the samples fields pointers
// It returns InvalidOperation Error if any of fields does not belong to the sample
func NewKey(name string, sample interface{}, sampleFields ...interface{}) (Key, *Error) {
	typ := reflect.TypeOf(sample).Elem()
	ret := Key{Name: name, Type: typ}
	c := len(sampleFields)
	// if no fields provided return empty Key
	if c == 0 {
		return ret, nil
	}
	fids := make([]int, c, c)
	fnms := make([]string, c, c)
	sampleFirstPtr := reflect.ValueOf(sample).Pointer()
	sampleLeastPtr := sampleFirstPtr + typ.Size()

	for i := 0; i < c; i++ {
		sfPtr := reflect.ValueOf(sampleFields[i]).Pointer()
		if sfPtr < sampleFirstPtr || sampleLeastPtr <= sfPtr {
			return ret, NewError(InvalidOperation, ERR_FIELDPTROUTOFRNG)
		}
		for fi := 0; fi < typ.NumField(); fi++ {
			if sfPtr == sampleFirstPtr+typ.Field(fi).Offset {
				fids[i] = fi
				fnms[i] = typ.Field(fi).Name
				break
			}
		}
	}
	ret.fieldsIds = fids
	ret.fieldsNames = fnms
	return ret, nil
}

// Extracts fields names and values for given instance as arrays
// NB: returns Error InvalidArgument if type registered for key is not the same as for given instance
func (k *Key) Extract(i interface{}) ([]string, []interface{}, *Error) {
	var typ reflect.Type
	var val reflect.Value
	if reflect.TypeOf(i).Kind() == reflect.Ptr {
		typ = reflect.TypeOf(i).Elem()
		val = reflect.ValueOf(i).Elem()
	} else {
		typ = reflect.TypeOf(i)
		val = reflect.ValueOf(i)
	}

	if k.Type != typ {
		return nil, nil, NewError(InvalidArgument, ERR_KEYTYPENOTMATCH)
	}

	vals := make([]interface{}, len(k.fieldsIds))
	for fi := 0; fi < len(k.fieldsIds); fi++ {
		vals[fi] = (val.Field(k.fieldsIds[fi])).Interface()
	}
	return k.fieldsNames, vals, nil
}

// Assigns values from array into given instance fields
// NB: returns Error InvalidArgument if type registered for key is not the same as for given instance
//     returns Error InvalidArgument if i is not a kind of pointer
func (k *Key) Assign(i interface{}, vals []interface{}) *Error {
	var typ reflect.Type
	var val reflect.Value
	if reflect.TypeOf(i).Kind() == reflect.Ptr {
		typ = reflect.TypeOf(i).Elem()
		val = reflect.ValueOf(i).Elem()
	} else {
		return NewError(InvalidArgument, ERR_NONPTRARGUMENT)
	}

	if k.Type != typ {
		return NewError(InvalidArgument, ERR_KEYTYPENOTMATCH)
	}

	//vals := make([]interface{}, len(k.fieldsIds))
	for fi := 0; fi < len(k.fieldsIds); fi++ {
		val.Field(k.fieldsIds[fi]).Set(reflect.ValueOf(vals[fi]))
	}
	return nil
}

// Contains all needed data for mapping between Model(struct) and StorageObject(table/view/collection) field/column
type FieldMapping struct {
	StorageObjectFieldName string   // corresponding storage field
	StructureFieldName     string   // just for testing, cause access field by name almost 2 times slower than by index
	StructureFieldId       int      // structure FieldId. Assigned on registration
	ParticipateInKeys      []string // the list of real or virtual keys
	AssignedByDb           bool     // flags whether value is assigned by db. ForEx: sequence, default or calculated value
	ConvertViaDriver       bool     // flags whether value should be converted via driver(prepared query) or directly by storage query builder
}

// Contains all needed data for mapping between Model(struct) and StorageObject(table/view/collection) subset of fields/columns participating in the KEY
type KeyMapping struct {
	Key
	SOFieldsNames []string
}

// Contains all needed data for mapping between Model(struct) and StorageObject(table/view/collection) all fields/columns
type StructureMapping struct {
	StorageObjectName string

	// All fields mapping
	FieldsMappings []FieldMapping

	// Map of preset keys data
	KeyMappings map[string]KeyMapping
}

// Field tag JSON structure
// Helper structure to parse mapping tag's JSON
type TagJsonStruct struct {
	ColName          string
	Keys             []string
	AssignedByDb     bool
	ConvertViaDriver bool
}

// Generates structure mapping.
// If typ is not a kind of struct - returns nil, ERR_NONSTRUCTTYPE
// if failed to generate mapping  - returns nil, ERR_FAILEDTOGENMAP
func (ss *SqlStorage) generateStructureMapping(storageObjectName string, typ reflect.Type) (*StructureMapping, *Error) {
	if typ.Kind() != reflect.Struct {
		return nil, NewError(InvalidArgument, ERR_NONSTRUCTTYPE)
	}

	// get maximum possible capacity for arrays
	c := typ.NumField()

	// make slice to hold field mappings
	fieldMappings := make([]FieldMapping, 0, c)

	// Get storage object fields(columns)
	storageObjectFields, err := ss.GetStorageObjectFields(storageObjectName)
	if err != nil {
		// any furthem mapping would be invalid
		ss.log.Error(ERR_FAILEDTOGENMAP, err)
		return nil, NewError(InvalidOperation, ERR_FAILEDTOGENMAP)
	}

	// loop over fields generating field mapping for eachone
	for fi := 0; fi < c; fi++ {
		// Get field tag
		fieldTagString := typ.Field(fi).Tag.Get(ss.conf.MappingTag)
		fieldTagValue := new(TagJsonStruct)

		// if Tag is not empty - parse as Json
		if fieldTagString != EMPTY_STRING {
			err := json.Unmarshal([]byte(strings.Replace(fieldTagString, "'", string('"'), -1)), fieldTagValue)
			if err != nil || fieldTagValue == nil {
				ss.log.Error(ERR_FAILEDTOPARSETAG, err)
			}
		}

		fieldName := typ.Field(fi).Name
		comparisonFieldName := fieldName
		if fieldTagValue != nil && fieldTagValue.ColName != EMPTY_STRING {
			comparisonFieldName = fieldTagValue.ColName
		}

		var correspondingStorageObjectFieldName string
		// Now looking for matching field at storageObjectFields
		for x := 0; x < len(storageObjectFields); x++ {
			if ss.namesMatch(comparisonFieldName, storageObjectFields[x].Name) {
				correspondingStorageObjectFieldName = storageObjectFields[x].Name
				break
			}
		}

		if correspondingStorageObjectFieldName != EMPTY_STRING {
			fm := FieldMapping{
				StorageObjectFieldName: correspondingStorageObjectFieldName,
				StructureFieldName:     fieldName,
				StructureFieldId:       fi,
				// Defaults
				ParticipateInKeys: []string{EMPTY_STRING},
				AssignedByDb:      false,
				ConvertViaDriver:  false,
			}

			if fieldTagValue != nil {
				fm.ParticipateInKeys = append(fm.ParticipateInKeys, fieldTagValue.Keys...)
				fm.AssignedByDb = fieldTagValue.AssignedByDb
				fm.ConvertViaDriver = fieldTagValue.ConvertViaDriver
			}
			fieldMappings = append(fieldMappings, fm)
		}
	}

	if len(fieldMappings) == 0 {
		ss.log.Error(ERR_FAILEDTOGENMAP, "Generated mapping is empty")
		return nil, NewError(InvalidOperation, ERR_FAILEDTOGENMAP)
	}

	keyMappings := make(map[string]KeyMapping)
	// fillup of  Keys  map[string]KeyMapping
	for _, fieldMapping := range fieldMappings {
		if len(fieldMapping.ParticipateInKeys) > 0 {
			for _, keyName := range fieldMapping.ParticipateInKeys {
				keyMapping, ok := keyMappings[keyName]
				if ok {
					keyMapping.Key.fieldsIds = append(keyMapping.fieldsIds, fieldMapping.StructureFieldId)
					keyMapping.Key.fieldsNames = append(keyMapping.fieldsNames, fieldMapping.StructureFieldName)
					keyMapping.SOFieldsNames = append(keyMapping.SOFieldsNames, fieldMapping.StorageObjectFieldName)
				} else {
					keyMapping = KeyMapping{Key: Key{Name: keyName, Type: typ, fieldsIds: []int{fieldMapping.StructureFieldId}, fieldsNames: []string{fieldMapping.StructureFieldName}}, SOFieldsNames: []string{fieldMapping.StorageObjectFieldName}}
				}
				keyMappings[keyName] = keyMapping

			}
		}
	}

	return &StructureMapping{
		StorageObjectName: storageObjectName,
		FieldsMappings:    fieldMappings,
		KeyMappings:       keyMappings,
	}, nil

}
