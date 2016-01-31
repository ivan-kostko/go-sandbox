// The file contains SqlStorage mapping functionality
package SqlStorage

import (
	"reflect"
)

type FieldMapping struct {
	StorageObjectFieldName string
	StructureFieldName     string
	StructureFieldTags     []string
}

type StructureMapping struct {
	StructureType     reflect.Type
	StorageObjectName string

	// Primary key fields mapping
	PkFieldsMappings []FieldMapping
	PkFieldsString   string

	// Business key
	BkFieldsMappings []FieldMapping
	BkFieldsString   string

	// Descriptive values
	ValFieldsMappings []FieldMapping
	ValFieldsString   string
}
