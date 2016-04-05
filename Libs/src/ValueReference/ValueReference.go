// PointerTo project PointerTo.go
package ValueReference

import (
	. "database/sql"
	"errors"
	"reflect"
)

// Declares generic ValueReference interface functionality
type ValueReferencer interface {
	AssignReferentValue(i interface{})
	GetReferentValue() interface{}
	ReInitializeReferentValue()
	Scanner
}

//ValueReference container
type ValueReference struct {
	ptr                       interface{}
	refrentType               reflect.Type
	isReferentPtr             bool
	assignReferentValue       func(vr *ValueReference, val interface{})
	getReferentValue          func(vr *ValueReference) interface{}
	reinitializeReferentValue func(vr *ValueReference)
}

// Generic factory. Fills up container depending on type.
// NB: In case of performance issues just implement custom functions and extend the type switch to avoid reflection.
func NewValueReference(iPtr interface{}) ValueReference {
	vr := ValueReference{
		ptr:           iPtr,
		refrentType:   (reflect.ValueOf(iPtr).Elem()).Type(),
		isReferentPtr: ((reflect.ValueOf(iPtr).Elem()).Type().Kind() == reflect.Ptr),
	}
	switch iPtr.(type) {
	case *int:
		vr.assignReferentValue = assignReferentIntValue
		vr.getReferentValue = getReferentIntValue
		vr.reinitializeReferentValue = func(vr *ValueReference) { return } // int wont be nil
		break
	case **int:
		vr.assignReferentValue = assignReferentIntPtrValue
		vr.getReferentValue = getReferentIntPtrValue
		vr.reinitializeReferentValue = reinitializesReferentIntPtrValue
		break
	default:
		vr.assignReferentValue = assignReferentValueByReflect
		vr.getReferentValue = getReferentValueByReflect
		vr.reinitializeReferentValue = reinitializeReferentValueByReflect
		break
	}
	return vr
}

// Assigns provided i to referent
func (vr *ValueReference) AssignReferentValue(i interface{}) {
	vr.assignReferentValue(vr, i)
}

// Gets referent value.
func (vr *ValueReference) GetReferentValue() interface{} {
	return vr.getReferentValue(vr)
}

// Reinitializes referent with new pointer if it is nil
func (vr *ValueReference) ReInitializeReferentValue() {
	vr.reinitializeReferentValue(vr)
}

// Scanner implementation
func (vr *ValueReference) Scan(src interface{}) error {
	srcType := reflect.TypeOf(src)
	if srcType == vr.refrentType {
		vr.AssignReferentValue(src)
		return nil
	}

	return errors.New("Unsupported type")
}

// Assigns provided i to int referent
func assignReferentIntValue(vr *ValueReference, val interface{}) {
	*((vr.ptr).(*int)) = val.(int)
}

// Gets int referent value.
func getReferentIntValue(vr *ValueReference) interface{} {
	return interface{}(*((vr.ptr).(*int)))
}

// Assigns provided i to *int referent
func assignReferentIntPtrValue(vr *ValueReference, val interface{}) {
	*((vr.ptr).(**int)) = val.(*int)
}

// Gets *int referent value.
func getReferentIntPtrValue(vr *ValueReference) interface{} {
	return interface{}(*((vr.ptr).(**int)))
}

// ReInitializes *int referent value if it is nil.
func reinitializesReferentIntPtrValue(vr *ValueReference) {
	if (*((vr.ptr).(**int))) == (*int)(nil) {
		(*((vr.ptr).(**int))) = new(int)
	}
	return
}

func assignReferentValueByReflect(vr *ValueReference, i interface{}) {
	reflect.ValueOf(vr.ptr).Elem().Set(reflect.ValueOf(i))
}

func getReferentValueByReflect(vr *ValueReference) interface{} {
	return reflect.ValueOf(vr.ptr).Elem().Interface()
}

func reinitializeReferentValueByReflect(vr *ValueReference) {
	if !vr.isReferentPtr {
		return
	}

	val := reflect.ValueOf(vr.ptr).Elem()
	if !val.IsNil() {
		return
	}

	val.Set(reflect.New(vr.refrentType.Elem()))
}

/*
//   The template for a custom type functions

// Assigns provided i to <Type> referent
func assignReferentIntValue(vr *ValueReference, val <Type>erface{}) {
	*((vr.ptr).(*<Type>)) = val.(<Type>)
}

// Gets <Type> referent value.
func getReferentIntValue(vr *ValueReference) <Type>erface{} {
	return <Type>erface{}(*((vr.ptr).(*<Type>)))
}

// Assigns provided i to *<Type> referent
func assignReferentIntPtrValue(vr *ValueReference, val <Type>erface{}) {
	*((vr.ptr).(**<Type>)) = val.(*<Type>)
}

// Gets *<Type> referent value.
func getReferentIntPtrValue(vr *ValueReference) <Type>erface{} {
	return <Type>erface{}(*((vr.ptr).(**<Type>)))
}

// Reinitializes *<Type> referent value if it is nil.
func reinitializesReferentIntPtrValue(vr *ValueReference) {
	if (*((vr.ptr).(**<Type>))) == (*<Type>)(nil) {
		(*((vr.ptr).(**<Type>))) = new(<Type>)
	}
	return
}

*/
