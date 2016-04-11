// PointerTo project PointerTo.go
package ValueReference

import (
	"reflect"
	"testing"
	"time"
)

func TestValueReferenceImplementsValueReferencer(t *testing.T) {
	x := new(ValueReference)
	_ = ValueReferencer(x)
}

func TestNewAsValueReference(t *testing.T) {
	var xInt int
	var xIntPtr *int
	var xString string
	var xStringPtr *string
	var xInt64 int64
	var xInt64Ptr *int64
	var xTime time.Time
	var xTimePtr *time.Time
	var xFloat32 float32
	var xFloat32Ptr *float32
	var xFloat64 float64
	var xFloat64Ptr *float64
	var xByte []byte
	var xBytePtr *[]byte

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
				setReferentValue:          setReferentIntValue,
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
				setReferentValue:          setReferentIntPtrValue,
				getReferentValue:          getReferentIntPtrValue,
				reinitializeReferentValue: reinitializesReferentIntPtrValue,
			},
		},
		{
			"String",
			&xString,
			ValueReference{
				reference:                 &xString,
				referentType:              reflect.Type(nil),
				setReferentValue:          setReferentValueByReflect,
				getReferentValue:          getReferentValueByReflect,
				reinitializeReferentValue: reinitializeReferentValueByReflect,
			},
		},
		{
			"StringPtr",
			&xStringPtr,
			ValueReference{
				reference:                 &xStringPtr,
				referentType:              reflect.Type(nil),
				setReferentValue:          setReferentValueByReflect,
				getReferentValue:          getReferentValueByReflect,
				reinitializeReferentValue: reinitializeReferentValueByReflect,
			},
		},
		{
			"Int64",
			&xInt64,
			ValueReference{
				reference:                 &xInt64,
				referentType:              reflect.Type(nil),
				setReferentValue:          setReferentValueByReflect,
				getReferentValue:          getReferentValueByReflect,
				reinitializeReferentValue: reinitializeReferentValueByReflect,
			},
		},
		{
			"Int64Ptr",
			&xInt64Ptr,
			ValueReference{
				reference:                 &xInt64Ptr,
				referentType:              reflect.Type(nil),
				setReferentValue:          setReferentValueByReflect,
				getReferentValue:          getReferentValueByReflect,
				reinitializeReferentValue: reinitializeReferentValueByReflect,
			},
		},
		{
			"Time",
			&xTime,
			ValueReference{
				reference:                 &xTime,
				referentType:              reflect.Type(nil),
				setReferentValue:          setReferentValueByReflect,
				getReferentValue:          getReferentValueByReflect,
				reinitializeReferentValue: reinitializeReferentValueByReflect,
			},
		},
		{
			"TimePtr",
			&xTimePtr,
			ValueReference{
				reference:                 &xTimePtr,
				referentType:              reflect.Type(nil),
				setReferentValue:          setReferentValueByReflect,
				getReferentValue:          getReferentValueByReflect,
				reinitializeReferentValue: reinitializeReferentValueByReflect,
			},
		},
		{
			"Float32",
			&xFloat32,
			ValueReference{
				reference:                 &xFloat32,
				referentType:              reflect.Type(nil),
				setReferentValue:          setReferentValueByReflect,
				getReferentValue:          getReferentValueByReflect,
				reinitializeReferentValue: reinitializeReferentValueByReflect,
			},
		},
		{
			"Float32Ptr",
			&xFloat32Ptr,
			ValueReference{
				reference:                 &xFloat32Ptr,
				referentType:              reflect.Type(nil),
				setReferentValue:          setReferentValueByReflect,
				getReferentValue:          getReferentValueByReflect,
				reinitializeReferentValue: reinitializeReferentValueByReflect,
			},
		},
		{
			"Float64",
			&xFloat64,
			ValueReference{
				reference:                 &xFloat64,
				referentType:              reflect.Type(nil),
				setReferentValue:          setReferentValueByReflect,
				getReferentValue:          getReferentValueByReflect,
				reinitializeReferentValue: reinitializeReferentValueByReflect,
			},
		},
		{
			"Float64Ptr",
			&xFloat64Ptr,
			ValueReference{
				reference:                 &xFloat64Ptr,
				referentType:              reflect.Type(nil),
				setReferentValue:          setReferentValueByReflect,
				getReferentValue:          getReferentValueByReflect,
				reinitializeReferentValue: reinitializeReferentValueByReflect,
			},
		},
		{
			"Byte",
			&xByte,
			ValueReference{
				reference:                 &xByte,
				referentType:              reflect.Type(nil),
				setReferentValue:          setReferentValueByReflect,
				getReferentValue:          getReferentValueByReflect,
				reinitializeReferentValue: reinitializeReferentValueByReflect,
			},
		},
		{
			"BytePtr",
			&xBytePtr,
			ValueReference{
				reference:                 &xBytePtr,
				referentType:              reflect.Type(nil),
				setReferentValue:          setReferentValueByReflect,
				getReferentValue:          getReferentValueByReflect,
				reinitializeReferentValue: reinitializeReferentValueByReflect,
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
		if reflect.ValueOf(actual.setReferentValue).Pointer() != reflect.ValueOf(expected.setReferentValue).Pointer() {
			t.Errorf("Testing x as %v: New(&x)  returned setReferentValue %#v \n\t\t    while expected %#v", name, actual.setReferentValue, expected.setReferentValue)
		}
		if reflect.ValueOf(actual.getReferentValue).Pointer() != reflect.ValueOf(expected.getReferentValue).Pointer() {
			t.Errorf("Testing x as %v: New(&x)  returned getReferentValue %#v \n\t\t    while expected %#v", name, actual.getReferentValue, expected.getReferentValue)
		}
		if reflect.ValueOf(actual.reinitializeReferentValue).Pointer() != reflect.ValueOf(expected.reinitializeReferentValue).Pointer() {
			t.Errorf("Testing x as %v: New(&x)  returned reinitializeReferentValue %#v \n\t\t    while expected %#v", name, actual.reinitializeReferentValue, expected.reinitializeReferentValue)
		}

	}

}

func TestSetReferentValueInt(t *testing.T) {
	var xInt int
	var xIntPtr *int
	var xString string
	var xStringPtr *string
	var xInt64 int64
	var xInt64Ptr *int64
	var xTime time.Time
	var xTimePtr *time.Time
	var xFloat32 float32
	var xFloat32Ptr *float32
	var xFloat64 float64
	var xFloat64Ptr *float64
	var xByte []byte
	var xBytePtr *[]byte

	testCases := []struct {
		Name     string
		X        interface{}
		Expected interface{}
	}{
		{
			"xInt",
			&xInt,
			interface{}(5),
		},
		{
			"xIntPtr",
			&xIntPtr,
			interface{}(new(int)),
		},
		{
			"xString",
			&xString,
			interface{}("Test string"),
		},
		{
			"xStringPtr",
			&xStringPtr,
			interface{}(new(string)),
		},
		{
			"xInt64",
			&xInt64,
			interface{}(int64(9999999999999)),
		},
		{
			"xInt64Ptr",
			&xInt64Ptr,
			interface{}(new(int64)),
		},
		{
			"xTime",
			&xTime,
			interface{}(time.Now()),
		},
		{
			"xTimePtr",
			&xTimePtr,
			interface{}(new(time.Time)),
		},
		{
			"xFloat32",
			&xFloat32,
			interface{}(float32(567.987651012)),
		},
		{
			"xFloat32Ptr",
			&xFloat32Ptr,
			interface{}(new(float32)),
		},
		{
			"xFloat64",
			&xFloat64,
			interface{}(float64(56789101112.987651012)),
		},
		{
			"xFloat64Ptr",
			&xFloat64Ptr,
			interface{}(new(float64)),
		},
		{
			"xByte",
			&xByte,
			interface{}([]byte(`Testing set refernet value for []byte`)),
		},
		{
			"xBytePtr",
			&xBytePtr,
			interface{}(new([]byte)),
		},
	}

	for _, testCase := range testCases {
		name := testCase.Name
		x := testCase.X
		expected := testCase.Expected

		vr := New(x)
		vr.SetReferentValue(expected)
		actual := reflect.ValueOf((vr).(*ValueReference).reference).Elem().Interface()
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Testing x as %v: SetReferentValue() assigned %#v \n\t\t\twhile expected %#v", name, actual, expected)
		}
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
