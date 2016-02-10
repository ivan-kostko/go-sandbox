// The file contains SqlStorage mapping functionality tests
package SqlStorage

import (
	"reflect"
	"testing"

	. "customErrors"
)

// KEY tests
func TestNewKeySuccess(t *testing.T) {
	type MyTestType struct {
		Field1 string
		Field2 []byte
		Field3 *int
		Field4 float64
	}
	mtt := MyTestType{}
	expected := Key{Name: "TestNewKeySuccess", Type: reflect.TypeOf(mtt), fieldsIds: []int{1, 3}, fieldsNames: []string{"Field2", "Field4"}}
	actual, err := NewKey("TestNewKeySuccess", &mtt, &mtt.Field2, &mtt.Field4)
	if err != nil {
		t.Errorf("NewKey returned *Error %v while expected %v", err, nil)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("NewKey returned %v \r\n\t\t while expected %v", actual, expected)
	}
}

func TestNewKeyFail(t *testing.T) {
	type MyTestType struct {
		Field1 string
		Field2 []byte
		Field3 *int
		Field4 float64
	}
	mtt := MyTestType{}
	fake := mtt.Field4
	expected := Key{Name: "TestNewKeyFail", Type: reflect.TypeOf(mtt)}
	expectedErr := NewError(InvalidOperation, ERR_FIELDPTROUTOFRNG)
	actual, err := NewKey("TestNewKeyFail", &mtt, &mtt.Field2, &fake)
	if err == nil {
		t.Errorf("NewKey returned *Error %v while expected %v", err, expectedErr)
	}
	if *err != *expectedErr {
		t.Errorf("NewKey returned %v \r\n\t\t while expected %v", err, expectedErr)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("NewKey returned %v \r\n\t\t while expected %v", actual, expected)
	}
}

func TestNewKeyWoFields(t *testing.T) {
	type MyTestType struct {
		Field1 string
		Field2 []byte
		Field3 *int
		Field4 float64
	}
	mtt := MyTestType{}
	expected := Key{Name: "TestNewKeyWoFields", Type: reflect.TypeOf(mtt)}
	actual, err := NewKey("TestNewKeyWoFields", &mtt)
	if err != nil {
		t.Errorf("NewKey returned *Error %v while expected %v", err, nil)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("NewKey returned %v \r\n\t\t while expected %v", actual, expected)
	}
}

func TestKeyExtractFieldsByPtr(t *testing.T) {
	type MyTestType struct {
		Field1 string
		Field2 []byte
		Field3 *int
		Field4 float64
	}
	sample := MyTestType{}
	mtt := MyTestType{Field1: "Field1", Field2: []byte("Field2"), Field3: nil, Field4: float64(2345.98765)}
	expectedStr := []string{"Field1", "Field2", "Field4"}
	expectedVal := []interface{}{"Field1", []byte("Field2"), float64(2345.98765)}
	k, err := NewKey("TestKeyExtractFields", &sample, &sample.Field1, &sample.Field2, &sample.Field4)
	if err != nil {
		t.Errorf("NewKey returned *Error %v while expected %v", err, nil)
	}
	actualStr, actualVal, err := k.Extract(&mtt)
	if !(reflect.DeepEqual(actualStr, expectedStr) &&
		reflect.DeepEqual(actualVal, expectedVal) &&
		err == nil) {
		t.Errorf("NewKey returned Names: %v Vals: %v Error %v while expected Names: %v Vals: %v Error %v", actualStr, actualVal, err, expectedStr, expectedVal, nil)
	}

}

func TestKeyExtractFieldsByVal(t *testing.T) {
	type MyTestType struct {
		Field1 string
		Field2 []byte
		Field3 *int
		Field4 float64
	}
	sample := MyTestType{}
	mtt := MyTestType{Field1: "Field1", Field2: []byte("Field2"), Field3: nil, Field4: float64(2345.98765)}
	expectedStr := []string{"Field3", "Field2", "Field4"}
	expectedVal := []interface{}{((*int)(nil)), []byte("Field2"), float64(2345.98765)}
	k, err := NewKey("TestKeyExtractFields", &sample, &sample.Field3, &sample.Field2, &sample.Field4)
	if err != nil {
		t.Errorf("NewKey returned *Error %v while expected %v", err, nil)
	}
	actualStr, actualVal, err := k.Extract(mtt)
	if !(reflect.DeepEqual(actualStr, expectedStr) &&
		reflect.DeepEqual(actualVal, expectedVal) &&
		err == nil) {
		t.Errorf("NewKey returned Names: %v Vals: %v Error %v \r\n\t\t\twhile expected Names: %v Vals: %v Error %v", actualStr, actualVal, err, expectedStr, expectedVal, nil)
	}

}

//-------------//
// Benchmarks  //
//-------------//

func BenchmarkKeyExtractCustom6fields(b *testing.B) {
	type MyTestType struct {
		Field1  string
		Field2  []byte
		Field3  *int
		Field4  float64
		Field01 string
		Field02 []byte
		Field03 *int
		Field04 float64
	}
	mtt := MyTestType{Field1: "Field1", Field2: []byte("Field2"), Field3: nil, Field4: float64(2345.98765), Field01: "Field1", Field02: []byte("Field2"), Field03: nil, Field04: float64(2345.98765)}
	fn := func(i *MyTestType) ([]string, []interface{}, *Error) {
		return []string{"Field1", "Field2", "Field4", "Field01", "Field02", "Field04"}, []interface{}{i.Field1, i.Field2, i.Field4, i.Field01, i.Field02, i.Field04}, nil
	}
	b.ResetTimer()

	for n := 0; n <= b.N; n++ {
		_, _, _ = fn(&mtt)
	}
	b.ReportAllocs()
}

func BenchmarkKeyExtract6fields(b *testing.B) {
	type MyTestType struct {
		Field1  string
		Field2  []byte
		Field3  *int
		Field4  float64
		Field01 string
		Field02 []byte
		Field03 *int
		Field04 float64
	}
	sample := MyTestType{}
	mtt := MyTestType{Field1: "Field1", Field2: []byte("Field2"), Field3: nil, Field4: float64(2345.98765), Field01: "Field1", Field02: []byte("Field2"), Field03: nil, Field04: float64(2345.98765)}
	k, err := NewKey("BenchmarkKeyExtract", &sample, &sample.Field1, &sample.Field2, &sample.Field4, &sample.Field01, &sample.Field02, &sample.Field04)
	if err != nil {
		b.Errorf("NewKey returned *Error %v while expected %v", err, nil)
	}
	b.ResetTimer()

	for n := 0; n <= b.N; n++ {
		_, _, _ = k.Extract(&mtt)
	}
	b.ReportAllocs()
}

func BenchmarkKeyExtractCustom4fields(b *testing.B) {
	type MyTestType struct {
		Field1  string
		Field2  []byte
		Field3  *int
		Field4  float64
		Field01 string
		Field02 []byte
		Field03 *int
		Field04 float64
	}
	mtt := MyTestType{Field1: "Field1", Field2: []byte("Field2"), Field3: nil, Field4: float64(2345.98765), Field01: "Field1", Field02: []byte("Field2"), Field03: nil, Field04: float64(2345.98765)}
	fn := func(i *MyTestType) ([]string, []interface{}, *Error) {
		return []string{"Field1", "Field4", "Field01", "Field02"}, []interface{}{i.Field1, i.Field4, i.Field01, i.Field02}, nil
	}
	b.ResetTimer()

	for n := 0; n <= b.N; n++ {
		_, _, _ = fn(&mtt)
	}
	b.ReportAllocs()
}

func BenchmarkKeyExtract4fields(b *testing.B) {
	type MyTestType struct {
		Field1  string
		Field2  []byte
		Field3  *int
		Field4  float64
		Field01 string
		Field02 []byte
		Field03 *int
		Field04 float64
	}
	sample := MyTestType{}
	mtt := MyTestType{Field1: "Field1", Field2: []byte("Field2"), Field3: nil, Field4: float64(2345.98765), Field01: "Field1", Field02: []byte("Field2"), Field03: nil, Field04: float64(2345.98765)}
	k, err := NewKey("BenchmarkKeyExtract", &sample, &sample.Field1, &sample.Field4, &sample.Field01, &sample.Field02)
	if err != nil {
		b.Errorf("NewKey returned *Error %v while expected %v", err, nil)
	}
	b.ResetTimer()

	for n := 0; n <= b.N; n++ {
		_, _, _ = k.Extract(&mtt)
	}
	b.ReportAllocs()
}

func BenchmarkKeyExtractCustom2fields(b *testing.B) {
	type MyTestType struct {
		Field1  string
		Field2  []byte
		Field3  *int
		Field4  float64
		Field01 string
		Field02 []byte
		Field03 *int
		Field04 float64
	}
	mtt := MyTestType{Field1: "Field1", Field2: []byte("Field2"), Field3: nil, Field4: float64(2345.98765), Field01: "Field1", Field02: []byte("Field2"), Field03: nil, Field04: float64(2345.98765)}
	fn := func(i *MyTestType) ([]string, []interface{}, *Error) {
		return []string{"Field1", "Field02"}, []interface{}{i.Field1, i.Field02}, nil
	}
	b.ResetTimer()

	for n := 0; n <= b.N; n++ {
		_, _, _ = fn(&mtt)
	}
	b.ReportAllocs()
}

func BenchmarkKeyExtract2fields(b *testing.B) {
	type MyTestType struct {
		Field1  string
		Field2  []byte
		Field3  *int
		Field4  float64
		Field01 string
		Field02 []byte
		Field03 *int
		Field04 float64
	}
	sample := MyTestType{}
	mtt := MyTestType{Field1: "Field1", Field2: []byte("Field2"), Field3: nil, Field4: float64(2345.98765), Field01: "Field1", Field02: []byte("Field2"), Field03: nil, Field04: float64(2345.98765)}
	k, err := NewKey("BenchmarkKeyExtract", &sample, &sample.Field1, &sample.Field02)
	if err != nil {
		b.Errorf("NewKey returned *Error %v while expected %v", err, nil)
	}
	b.ResetTimer()

	for n := 0; n <= b.N; n++ {
		_, _, _ = k.Extract(&mtt)
	}
	b.ReportAllocs()
}
