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
	ERR_GNRATEDNILMAPING = "Generated mapping is <nil>"
	ERR_FAILEDTOPARSETAG = "Failed to parse JSON value at field tag"
)

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
type TagJsonContent struct {
	ColName          string
	Keys             []string
	AssignedByDb     bool
	ConvertViaDriver bool
}

// Registers the structure mapping in the storage
// TODO : keys registration
func (ss *SqlStorage) RegisterType(storageObjectName string, st interface{}, keys ...Key) *Error {
	typ := reflect.TypeOf(st)
	sm, err := ss.generateStructureMapping(storageObjectName, typ)
	if err != nil {
		//ss.log.Critical(err)
		return err
	}
	if sm == nil {
		ss.log.Alert(ERR_GNRATEDNILMAPING)
		return NewError(BasicError, ERR_GNRATEDNILMAPING)
	}
	ss.structureMappings[typ] = *sm
	return nil
}

// Generates structure mapping.
// If typ is not a kind of struct - returns nil, ERR_NONSTRUCTTYPE
// if failed to generate mapping  - returns nil, ERR_FAILEDTOGENMAP
func (ss *SqlStorage) generateStructureMapping(storageObjectName string, typ reflect.Type) (*StructureMapping, *Error) {
	// TODO : Register input keys
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
	//ss.log.Debug("storageObjectFields", storageObjectFields)

	// loop over fields generating field mapping for eachone
	for fi := 0; fi < c; fi++ {
		// Get field tag
		fieldTagString := typ.Field(fi).Tag.Get(ss.conf.MappingTag)
		fieldTagValue := new(TagJsonContent)

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
