// The file contains Key functionality
package SqlStorage

import (
	. "customErrors"
	"reflect"
)

const (
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
func (k *Key) ExtractFrom(i interface{}) ([]string, []interface{}, *Error) {
	var typ reflect.Type
	var val reflect.Value
	var isPtr bool
	if reflect.TypeOf(i).Kind() == reflect.Ptr {
		isPtr = true
	}

	if isPtr {
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
	if isPtr {
		for fi := 0; fi < len(k.fieldsIds); fi++ {
			vals[fi] = (val.Field(k.fieldsIds[fi])).Addr().Interface()
		}
	} else {
		for fi := 0; fi < len(k.fieldsIds); fi++ {
			vals[fi] = (val.Field(k.fieldsIds[fi])).Interface()
		}
	}
	return k.fieldsNames, vals, nil
}

// Assigns values from array into given instance fields
// NB: returns Error InvalidArgument if type registered for key is not the same as for given instance
//     returns Error InvalidArgument if i is not a kind of pointer
func (k *Key) AssignTo(i interface{}, vals []interface{}) *Error {
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
