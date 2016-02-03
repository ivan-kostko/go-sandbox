// The file contains SqlStorage mapping functionality
package SqlStorage

import (
	"reflect"
)

const (

)

const(
    ERR_NONSTRUCTTYPE = 'Wont generate structure mapping for non-struct type'
)

// Represents all needed data for mapping between Model(struct) and StorageObject(table/view/collection) field
type FieldMapping struct {
	StorageObjectFieldName string     // corresponding storage field
	StructureFieldName     string     // just for testing, cause access field by name almost 2 times slower than by index
    StructureFieldId       int        // structure FieldId. Assigned on registration
	ParticipateInKeys      []string   // the list of real or virtual keys
    AssignedByDb           bool       // flags whether value is assigned by db. ForEx: sequence, default or calculated value
    ConvertViaDriver       bool       // flags whether value should be converted via driver(prepared query) or directly by storage query builder
}

type KeyMapping struct {
    FieldsMappings []FieldMapping
}

type StructureMapping struct {
	StorageObjectName string

	// All fields mapping
	FieldsMappings []FieldMapping

    // Map of preset key data
    Keys           map[string]KeyMapping
}

// Generates structure mapping.If typ is not a kind of struct - returns nil, ERR_NONSTRUCTTYPE
func (ss *SqlStorage) generateStructureMapping(storageObjectName string, typ reflect.Type) (*StructureMapping, *Error) {
    if typ.Kind() != reflect.Struct{
        return nil, NewError(InvalidArgument, ERR_NONSTRUCTTYPE)
    }

    // get maximum possible capacity for arrays
    c := typ.NumField()

    pks := make([]FieldMapping, 0, c)
    bks := make([]FieldMapping, 0, c)
    vals := make([]FieldMapping, 0, c)

    // loop over fields generating field mapping for eachone
    for fi := 0; fi < c; fi++ {

    }

	return &StructureMapping{
		StructureType:     typ,
		StorageObjectName: storageObjectName,
		FieldMappings:     sfm,
	}, nil

}

func (ss *SqlStorage) generateStructureFieldMappings(storageObjectName string, typ reflect.Type) ([]FieldMapping, error) {
	// checks first
	if typ.Kind() != reflect.Struct {
		return nil, NewError(InvalidArgument, ERR_NONSTUCTTYPE, ERR_DEFAULT_SEVERITY)
	}
	var ret = make([]FieldMapping, 0, typ.NumField())
	var storageObjectFields = s.getStorageObjectFields(storageObjectName)

	for i := 0; i < typ.NumField(); i++ {
		fieldTagValue := typ.Field(i).Tag.Get(s.config.MappingTag)
		fieldName := typ.Field(i).Name
		// If no tag provided, take field name
		comparisonName := fieldName
		if fieldTagValue != "" {
			comparisonName = fieldTagValue
		}
		// Now looking for matching field at storageObjectFields
		for x := 0; x < len(storageObjectFields); x++ {
			if s.namesMatch(comparisonName, storageObjectFields[x]) {
				fm := FieldMapping{
					StorageObjectFieldName: storageObjectFields[x],
					StructureFieldName:     fieldName,
					// TODO : implement
					// StructureFieldTags:
				}
				ret = append(ret, fm)
				// StructureField found, go to next field
				break
			}
		}

	}
	if len(ret) == 0 {
		return nil, nil
	}
	return ret, nil
}
