package MySendBox

import (
	"reflect"
	"testing"
	"time"
)

func TestConvertBySwitch(t *testing.T) {
	expected := "string"
	actual := ConvertBySwitch("")
	if actual != expected {
		t.Errorf("ConvertBySwitch empty string returned %v while expected %v ", actual, expected)
	}
	expected = "string"
	xst := "blabla"
	actual = ConvertBySwitch(&xst)
	if actual != expected {
		t.Errorf("ConvertBySwitch empty string returned %v while expected %v ", actual, expected)
	}

	expected = "integer"
	actual = ConvertBySwitch(5)
	if actual != expected {
		t.Errorf("ConvertBySwitch(5) returned %v while expected %v ", actual, expected)
	}
	expected = "integer"
	xin := 5
	actual = ConvertBySwitch(&xin)
	if actual != expected {
		t.Errorf("ConvertBySwitch(5) returned %v while expected %v ", actual, expected)
	}

	expected = "float"
	actual = ConvertBySwitch(float32(5.001))
	if actual != expected {
		t.Errorf("ConvertBySwitch(5.001) returned %v while expected %v ", actual, expected)
	}
	expected = "float"
	xfl := 5.0001
	actual = ConvertBySwitch(&xfl)
	if actual != expected {
		t.Errorf("ConvertBySwitch(5.001) returned %v while expected %v ", actual, expected)
	}

	expected = "time"
	actual = ConvertBySwitch(time.Now())
	if actual != expected {
		t.Errorf("ConvertBySwitch(5.001) returned %v while expected %v ", actual, expected)
	}
	expected = "time"
	xtm := time.Now()
	actual = ConvertBySwitch(&xtm)
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

func TestConvertBytesIntoSqlString(t *testing.T) {
	v := []byte("Byte Test String")
	t.Log(ConvertBytesIntoSqlString(v))
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

func BenchmarkConvertIntToStringBySprintf(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		s := ConvertIntToStringBySprintf(n)
		s += ""
	}
	b.ReportAllocs()
}

func BenchmarkConvertIntToStringByStrConv(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		s := ConvertIntToStringByStrConv(n)
		s += ""
	}
	b.ReportAllocs()
}

func BenchmarkConvertFloatToStringBySprintf(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		s := ConvertFloatToStringBySprintf(26535141592653.1415926)
		s += ""
	}
	b.ReportAllocs()
}

func BenchmarkConvertFloatToStringByStrConv(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		s := ConvertFloatToStringByStrConv(26535141592653.1415926)
		s += ""
	}
	b.ReportAllocs()
}

func BenchmarkConvertInt64ToStringBySprintf(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		s := ConvertInt64ToStringBySprintf(9223372036854775807)
		s += ""
	}
	b.ReportAllocs()
}

func BenchmarkConvertInt64ToStringByStrConv(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		s := ConvertInt64ToStringByStrConv(9223372036854775807)
		s += ""
	}
	b.ReportAllocs()
}

func BenchmarkConvertTimeToStringBySprintf(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		s := ConvertTimeToStringBySprintf(time.Now())
		s += ""
	}
	b.ReportAllocs()

}
func BenchmarkConvertTimeToStringByForma(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		s := ConvertTimeToStringByFormat(time.Now())
		s += ""
	}
	b.ReportAllocs()
}
func BenchmarkConvertTimeToStringByStringer(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		s := ConvertTimeToStringByStringer(time.Now())
		s += ""
	}
	b.ReportAllocs()
}

func BenchmarkGenerateStringsReplace(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		s := GenerateStringsReplace("dbo.TestTable", "Col1, Col2, Col3", "Value1, Value3, Value4")
		s += ""
	}
	b.ReportAllocs()
}
func BenchmarkGenerateFmtSprintf(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		s := GenerateFmtSprintf("dbo.TestTable", "Col1, Col2, Col3", "Value1, Value3, Value4")
		s += ""
	}
	b.ReportAllocs()
}
func BenchmarkGenerateCustom(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		s := GenerateCustom("dbo.TestTable", "Col1, Col2, Col3", "Value1, Value3, Value4")
		s += ""
	}
	b.ReportAllocs()
}

/*
func Benchmark(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		s :=
		s += ""
	}
	b.ReportAllocs()
}

*/
