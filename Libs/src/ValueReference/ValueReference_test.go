// PointerTo project PointerTo.go
package ValueReference

import (
	"reflect"
	"testing"
	//"time"
)

func TestValueReferenceImplementsValueReferencer(t *testing.T) {
	x := new(ValueReference)
	_ = ValueReferencer(x)
}

func TestNewValueReference(t *testing.T) {
	var xInt int
	var xIntPtr *int
	//	var xString string
	//	var xStringPtr *string
	//	var xInt64 int64
	//	var xInt64Ptr *int64
	//	var xTime time.Time
	//	var xTimePtr *time.Time
	//	var xFloat32 float32
	//	var xFloat32Ptr *float32
	//	var xFloat64 float64
	//	var xFloat64Ptr *float64
	//	var xByte []byte

	testCases := []struct {
		Name     string
		X        interface{}
		Expected ValueReference
	}{
		{
			"Int",
			&xInt,
			ValueReference{
				reference:                 &xInt,
				referentType:              reflect.Type(nil),
				assignReferentValue:       assignReferentIntValue,
				getReferentValue:          getReferentIntValue,
				reinitializeReferentValue: nil,
			},
		},
		{
			"IntPtr",
			&xIntPtr,
			ValueReference{
				reference:                 &xIntPtr,
				referentType:              reflect.Type(nil),
				assignReferentValue:       assignReferentIntPtrValue,
				getReferentValue:          getReferentIntPtrValue,
				reinitializeReferentValue: reinitializesReferentIntPtrValue,
			},
		},
	}

	for _, testCase := range testCases {
		name := testCase.Name
		x := testCase.X
		expected := testCase.Expected

		actual := *(New(x)).(*ValueReference)

		if actual.reference != expected.reference {
			t.Errorf("Testing x as %v: New(&x)  returned reference %#v \n\t\t    while expected %#v", name, actual.reference, expected.reference)
		}
		if actual.referentType != expected.referentType {
			t.Errorf("Testing x as %v: New(&x)  returned refrentType %#v \n\t\t    while expected %#v", name, actual.referentType, expected.referentType)
		}
		if reflect.ValueOf(actual.assignReferentValue).Pointer() != reflect.ValueOf(expected.assignReferentValue).Pointer() {
			t.Errorf("Testing x as %v: New(&x)  returned assignReferentValue %#v \n\t\t    while expected %#v", name, actual.assignReferentValue, expected.assignReferentValue)
		}
		if reflect.ValueOf(actual.getReferentValue).Pointer() != reflect.ValueOf(expected.getReferentValue).Pointer() {
			t.Errorf("Testing x as %v: New(&x)  returned getReferentValue %#v \n\t\t    while expected %#v", name, actual.getReferentValue, expected.getReferentValue)
		}
		if reflect.ValueOf(actual.reinitializeReferentValue).Pointer() != reflect.ValueOf(expected.reinitializeReferentValue).Pointer() {
			t.Errorf("Testing x as %v: New(&x)  returned reinitializeReferentValue %#v \n\t\t    while expected %#v", name, actual.reinitializeReferentValue, expected.reinitializeReferentValue)
		}

	}

}

func TestAssignReferentValueInt(t *testing.T) {
	px := new(int)
	expected := 4

	cp := New(px)
	cp.AssignReferentValue(4)
	actual := *px
	if actual != expected {
		t.Errorf(" returned %#v \n\t\t\twhile expected %#v", actual, expected)
	}
}

func TestGetReferentValueInt(t *testing.T) {
	px := new(int)
	expected := 5

	cp := New(px)
	*px = 5
	actual := cp.GetReferentValue()
	if actual != expected {
		t.Errorf(" returned %#v while expected %#v", actual, expected)
	}
}

func TestReinitializeReferentValueByReflect(t *testing.T) {

	x := (*int)(nil)
	expected := 0

	vr := *(New(&x)).(*ValueReference)

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
