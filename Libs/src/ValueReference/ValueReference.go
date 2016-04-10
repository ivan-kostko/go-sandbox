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
	reference                 interface{}
	referentType              reflect.Type
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
// TODO(me):
// Extend package with custom dereferecing for at least the following types:
//  &string, &(*string)
//  &int64, &(*int64)
//  &time.Time, &(*time.Time)
//  &[]byte
//  &float32,&float64, &(*float32), &(*float64)
func New(iPtr interface{}) ValueReferencer {
	vr := ValueReference{
		reference: iPtr,
	}
	switch iPtr.(type) {
	case *int:
		vr.assignReferentValue = assignReferentIntValue
		vr.getReferentValue = getReferentIntValue
		vr.reinitializeReferentValue = nil // int wont be nil
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

// Returns type of referent.
// It is implemented as Lazy initialization, so on a first call could be slow
// If ValueReferencer represents *T then the method returns reflect.TypeOf(T)
func (vr *ValueReference) GetReferentType() reflect.Type {
	if vr.referentType == nil {
		vr.referentType = (reflect.ValueOf(vr.reference).Elem()).Type()
	}
	return vr.referentType
}

// Returns whether the referent is pointer
// If ValueReferencer represents **T then the method returns true
// If ValueReferencer represents *T where T is not pointer type, then the method returns false
func (vr *ValueReference) IsReferentPtr() bool {
	return vr.GetReferentType().Kind() == reflect.Ptr
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
	if vr.reinitializeReferentValue != nil {
		vr.reinitializeReferentValue(vr)
	}
}

// Assigns provided i to int referent
func assignReferentIntValue(vr *ValueReference, val interface{}) {
	*((vr.reference).(*int)) = val.(int)
}

// Gets int referent value.
func getReferentIntValue(vr *ValueReference) interface{} {
	return interface{}(*((vr.reference).(*int)))
}

// Assigns provided i to *int referent
func assignReferentIntPtrValue(vr *ValueReference, val interface{}) {
	*((vr.reference).(**int)) = val.(*int)
}

// Gets *int referent value.
func getReferentIntPtrValue(vr *ValueReference) interface{} {
	return interface{}(*((vr.reference).(**int)))
}

// ReInitializes *int referent value if it is nil.
func reinitializesReferentIntPtrValue(vr *ValueReference) {
	if (*((vr.reference).(**int))) == (*int)(nil) {
		(*((vr.reference).(**int))) = new(int)
	}
	return
}

func assignReferentValueByReflect(vr *ValueReference, i interface{}) {
	reflect.ValueOf(vr.reference).Elem().Set(reflect.ValueOf(i))
}

func getReferentValueByReflect(vr *ValueReference) interface{} {
	return reflect.ValueOf(vr.reference).Elem().Interface()
}

func reinitializeReferentValueByReflect(vr *ValueReference) {
	if !vr.IsReferentPtr() {
		return
	}

	val := reflect.ValueOf(vr.reference).Elem()
	if !val.IsNil() {
		return
	}

	val.Set(reflect.New(vr.referentType.Elem()))
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
