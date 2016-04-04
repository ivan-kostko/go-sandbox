// PointerTo project PointerTo.go
package ValueReference

import (
	"reflect"
	//"unsafe"
)

// Declares generic ValueReference interface functionality
type IValueReference interface {
	AssignReferentValue(i interface{})
	GetReferentValue() interface{}
	ReInitializeReferentValue()
}

type ValueReference struct {
	ptr                       interface{}
	refrentType               reflect.Type
	isReferentPtr             bool
	assignReferentValue       func(vr *ValueReference, val interface{})
	getReferentValue          func(vr *ValueReference) interface{}
	reinitializeReferentValue func(vr *ValueReference)
}

// Generic constructor
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
		break
	case **int:
		vr.assignReferentValue = assignReferentIntPtrValue
		vr.getReferentValue = getReferentIntPtrValue
		break
	default:
		vr.assignReferentValue = assignReferentValueByReflect
		vr.getReferentValue = getReferentValueByReflect
		break
	}
	return vr
}

func (vr *ValueReference) AssignReferentValue(i interface{}) {
	vr.assignReferentValue(vr, i)
}

func (vr *ValueReference) GetReferentValue() interface{} {
	return vr.getReferentValue(vr)
}

func (vr *ValueReference) ReInitializeReferentValue() {
	vr.reinitializeReferentValue(vr)
}

func assignReferentIntValue(vr *ValueReference, val interface{}) {
	*((vr.ptr).(*int)) = val.(int)
}

func getReferentIntValue(vr *ValueReference) interface{} {
	return interface{}(*((vr.ptr).(*int)))
}

func assignReferentIntPtrValue(vr *ValueReference, val interface{}) {
	*((vr.ptr).(**int)) = val.(*int)
}

func getReferentIntPtrValue(vr *ValueReference) interface{} {
	return interface{}(*((vr.ptr).(**int)))
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

func TestHelperPurePointer(px *int, x interface{}) {
	*px = (x).(int)
	_ = *px
}

func TestHelperCleverPointer(vr ValueReference, val interface{}) {
	vr.AssignReferentValue(val)
	_ = vr.GetReferentValue()
}

/*
func assignReferent<Type>Value(vr *ValueReference,val interface{}) {
	*((vr.ptr).(*<Type>)) = val.(<Type>)
}

func getReferent<Type>Value(vr *ValueReference) interface{} {
	return *((vr.ptr).(*<Type>))
}

func assignReferent<Type>PtrValue(vr *ValueReference,val interface{}) {
	*((vr.ptr).(**<Type>)) = val.(*<Type>)
}

func getReferent<Type>PtrValue(vr *ValueReference) interface{} {
	return *((vr.ptr).(**<Type>))
}

*/
