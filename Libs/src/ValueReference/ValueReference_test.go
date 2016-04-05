// PointerTo project PointerTo.go
package ValueReference

import (
	"reflect"
	"testing"
)

func TestNewValueReferenceIntPtr(t *testing.T) {
	x := new(int)
	expected := ValueReference{
		ptr:                       &x,
		refrentType:               reflect.TypeOf(new(int)),
		isReferentPtr:             true,
		assignReferentValue:       assignReferentIntPtrValue,
		getReferentValue:          getReferentIntPtrValue,
		reinitializeReferentValue: reinitializesReferentIntPtrValue,
	}
	actual := NewValueReference(&x)
	if actual.ptr != expected.ptr {
		t.Errorf("NewValueReference returned ptr %#v \n\t\t    while expected %#v", actual.ptr, expected.ptr)
	}
	if actual.refrentType != expected.refrentType {
		t.Errorf("NewValueReference returned refrentType %#v \n\t\t    while expected %#v", actual.refrentType, expected.refrentType)
	}
	if actual.isReferentPtr != expected.isReferentPtr {
		t.Errorf("NewValueReference returned isReferentPtr %#v \n\t\t    while expected %#v", actual.isReferentPtr, expected.isReferentPtr)
	}
	if reflect.ValueOf(actual.assignReferentValue).Pointer() != reflect.ValueOf(expected.assignReferentValue).Pointer() {
		t.Errorf("NewValueReference returned assignReferentValue %#v \n\t\t    while expected %#v", actual.assignReferentValue, expected.assignReferentValue)
	}
	if reflect.ValueOf(actual.getReferentValue).Pointer() != reflect.ValueOf(expected.getReferentValue).Pointer() {
		t.Errorf("NewValueReference returned getReferentValue %#v \n\t\t    while expected %#v", actual.getReferentValue, expected.getReferentValue)
	}
	if reflect.ValueOf(actual.reinitializeReferentValue).Pointer() != reflect.ValueOf(expected.reinitializeReferentValue).Pointer() {
		t.Errorf("NewValueReference returned getReferentValue %#v \n\t\t    while expected %#v", actual.reinitializeReferentValue, expected.reinitializeReferentValue)
	}
}

func TestNewValueReferenceStringPtr(t *testing.T) {
	x := new(string)
	expected := ValueReference{
		ptr:                       &x,
		refrentType:               reflect.TypeOf(new(string)),
		isReferentPtr:             true,
		assignReferentValue:       assignReferentValueByReflect,
		getReferentValue:          getReferentValueByReflect,
		reinitializeReferentValue: reinitializeReferentValueByReflect,
	}
	actual := NewValueReference(&x)
	if actual.ptr != expected.ptr {
		t.Errorf("NewValueReference returned ptr %#v \n\t\t    while expected %#v", actual.ptr, expected.ptr)
	}
	if actual.refrentType != expected.refrentType {
		t.Errorf("NewValueReference returned refrentType %#v \n\t\t    while expected %#v", actual.refrentType, expected.refrentType)
	}
	if actual.isReferentPtr != expected.isReferentPtr {
		t.Errorf("NewValueReference returned isReferentPtr %#v \n\t\t    while expected %#v", actual.isReferentPtr, expected.isReferentPtr)
	}
	if reflect.ValueOf(actual.assignReferentValue).Pointer() != reflect.ValueOf(expected.assignReferentValue).Pointer() {
		t.Errorf("NewValueReference returned assignReferentValue %#v \n\t\t    while expected %#v", actual.assignReferentValue, expected.assignReferentValue)
	}
	if reflect.ValueOf(actual.getReferentValue).Pointer() != reflect.ValueOf(expected.getReferentValue).Pointer() {
		t.Errorf("NewValueReference returned getReferentValue %#v \n\t\t    while expected %#v", actual.getReferentValue, expected.getReferentValue)
	}
	if reflect.ValueOf(actual.reinitializeReferentValue).Pointer() != reflect.ValueOf(expected.reinitializeReferentValue).Pointer() {
		t.Errorf("NewValueReference returned getReferentValue %#v \n\t\t    while expected %#v", actual.reinitializeReferentValue, expected.reinitializeReferentValue)
	}
}

func TestAssignReferentValue(t *testing.T) {
	px := new(int)
	expected := 4

	cp := NewValueReference(px)
	cp.AssignReferentValue(4)
	actual := *px
	if actual != expected {
		t.Errorf(" returned %#v \n\t\t\twhile expected %#v", actual, expected)
	}
}

func TestGetReferentValue(t *testing.T) {
	px := new(int)
	expected := 5

	cp := NewValueReference(px)
	*px = 5
	actual := cp.GetReferentValue()
	if actual != expected {
		t.Errorf(" returned %#v while expected %#v", actual, expected)
	}
}

func TestReinitializeReferentValueByReflect(t *testing.T) {

	x := (*int)(nil)
	expected := 0

	vr := NewValueReference(&x)

	initialPtr := getReferentIntPtrValue(&vr).(*int)
	if initialPtr != (*int)(nil) {
		t.Logf("Initial value is %#v while expected %#v", initialPtr, (*int)(nil))
	}

	reinitializeReferentValueByReflect(&vr)
	actualPtr := getReferentIntPtrValue(&vr).(*int)
	if actualPtr == (*int)(nil) {
		t.Errorf("actualPtr value is %#v while expected NOT NIL", actualPtr)
	}
	actual := *(actualPtr)
	if actual != expected {
		t.Errorf("reinitializeReferentValueByReflect returned %#v while expected %#v", actual, expected)
	}
}

/*
func Test(t *testing.T) {
    expected :=

    actual :=
	if actual != expected {
		t.Errorf(" returned %#v while expected %#v", actual, expected)
	}
}
*/

//--------------------------------------//
//             Benchmarks               //
//--------------------------------------//

/*
func Benchmark(b *testing.B) {
	b.ResetTimer()
	for n := 0; n <= b.N; n++ {
		s :=
		s += ""
	}
	b.ReportAllocs()
}
*/
