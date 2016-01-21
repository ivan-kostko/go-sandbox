package MySendBox

import (
	"reflect"
	"testing"
)

func TestConvertBySwitch(t *testing.T) {
	expected := "string"
	actual := ConvertBySwitch("")
	if actual != expected {
		t.Errorf("ConvertBySwitch empty string returned %v while expected %v ", actual, expected)
	}
	expected = "integer"
	actual = ConvertBySwitch(5)
	if actual != expected {
		t.Errorf("ConvertBySwitch(5) returned %v while expected %v ", actual, expected)
	}
	expected = "float"
	actual = ConvertBySwitch(5.001)
	if actual != expected {
		t.Errorf("ConvertBySwitch(5.001) returned %v while expected %v ", actual, expected)
	}
	expected = "bit"
	actual = ConvertBySwitch(true)
	if actual != expected {
		t.Errorf("ConvertBySwitch(true) returned %v while expected %v ", actual, expected)
	}
	expected = "binary"
	actual = ConvertBySwitch([]byte(nil))
	if actual != expected {
		t.Errorf("ConvertBySwitch(true) returned %v while expected %v ", actual, expected)
	}
}

func BenchmarkMap(b *testing.B) {
	FuncMap = make(map[reflect.Type]func(interface{}) string, 5)
	FuncMap[reflect.TypeOf("")] = FuncForString
	FuncMap[reflect.TypeOf(5)] = FuncForInteger
	FuncMap[reflect.TypeOf(5.001)] = FuncForFloat
	FuncMap[reflect.TypeOf(true)] = FuncForBool
	FuncMap[reflect.TypeOf([]byte(nil))] = FuncForBytes

	b.ResetTimer()
	var s string
	for n := 0; n <= b.N; n++ {
		s = FuncMap[reflect.TypeOf("")]("")
		s = FuncMap[reflect.TypeOf(5)](5)
		s = FuncMap[reflect.TypeOf(5.001)](5.001)
		s = FuncMap[reflect.TypeOf(true)](true)
		s = FuncMap[reflect.TypeOf([]byte(nil))]([]byte(nil))
		s += ""
	}
	b.ReportAllocs()
}

func BenchmarkTypeSwitch(b *testing.B) {

	b.ResetTimer()
	for n := 0; n <= b.N; n++ {
		s := ConvertBySwitch("")
		s = ConvertBySwitch(5)
		s = ConvertBySwitch(5.001)
		s = ConvertBySwitch(true)
		s = ConvertBySwitch([]byte(nil))
		s += ""
	}
	b.ReportAllocs()
}
