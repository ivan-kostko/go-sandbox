// ValueReference project ValueReference.go
package ValueReference

import (
	"reflect"
)

// Declares generic ValueReference interface functionality
type ValueReferencer interface {

	// Returns type of referent.
	// If ValueReferencer represents *T then the method returns reflect.TypeOf(T)
	GetReferentType() reflect.Type

	// Returns whether the referent is pointer
	// If ValueReferencer represents **T then the method returns true
	// If ValueReferencer represents *T where T is not pointer type, then the method returns false
	IsReferentPtr() bool

	// Assigns provided value to underlying referent
	// For.ex. if ValueReferencer represents v := &x , then it assigns i to x
	// It panics if ValueReferencer represents *T i is not assertable to i.(*T)
	AssignReferentValue(i interface{})

	// Gets underlying referent value
	// For.ex. if ValueReferencer represents v := &x , then it returns x as interface{}
	GetReferentValue() interface{}

	// Reinitializes underlying referent
	// Applicable if ValueReferencer represents v of type **T, and *v == nil
	// Then the method instantiates a new(T) reference and assigns it to *v
	ReInitializeReferentValue()
}

// ValueReference is basic implementation of ValueReferencer interface
// It holds reference to underlying referent value
type ValueReference struct {
	ptr                       interface{}
	refrentType               reflect.Type
	isReferentPtr             bool
	assignReferentValue       func(vr *ValueReference, val interface{})
	getReferentValue          func(vr *ValueReference) interface{}
	reinitializeReferentValue func(vr *ValueReference)
}

// Generic ValueReferencer factory.
// It initializes new ValueReference and returns it as ValueReferencer interface.
//
// Currently it supports as custom(optimized) dereferencing :
//     * &int and &(*int)
//
// All other type are dereferenced via reflection.
//
// TODO : Extend package with custom dereferecing for at least the following types:
//    * &string, &(*string)
//    * &int64, &(*int64)
//    * &time.Time, &(*time.Time)
//    * &[]byte
//    * &float32,&float64, &(*float32), &(*float64)
func New(iPtr interface{}) ValueReferencer {
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
		// Evrything by Reflection
		vr.assignReferentValue = assignReferentValueByReflect
		vr.getReferentValue = getReferentValueByReflect
		vr.reinitializeReferentValue = reinitializeReferentValueByReflect
		break
	}
	return &vr
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
