package MySendBox

import (
	"encoding/json"
	"reflect"
	"strings"
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

func TestReflectFieldTags(t *testing.T) {
	type MyTestType struct {
		Field1 *int `db:"{'ColName':'Column1','ColName':'Column2', 'Keys':['PK', 'BK'], 'ResolvedByDb':true}"`
		Field2 *string
	}
	type TagJsonStruct struct {
		ColName         string
		Keys            []string
		ResolvedByDb    bool
		ConvertByDriver bool
	}
	mtt := MyTestType{}
	c := reflect.TypeOf(mtt).NumField()
	for fi := 0; fi < c; fi++ {
		tag := reflect.TypeOf(mtt).Field(fi).Tag.Get("db")
		t.Log(tag)
		var tgs *TagJsonStruct
		tgs = new(TagJsonStruct)
		err := json.Unmarshal([]byte(strings.Replace(tag, "'", string('"'), -1)), tgs)
		if err != nil {
			t.Log(err)
		}
		t.Log(*tgs)
	}
}

func TestReflectFields(t *testing.T) {
	type MyTestType struct {
		Field1 string
		Field2 []byte
		Field3 *int
		Field4 float64
	}
	type TagJsonStruct struct {
		ColName         string
		Keys            []string
		ResolvedByDb    bool
		ConvertByDriver bool
	}
	mtt := MyTestType{}
	c := reflect.TypeOf(mtt).NumField()
	//val := reflect.ValueOf(mtt)
	for fi := 0; fi < c; fi++ {
		t.Logf("The field pointer is %v and calculated is %v", reflect.ValueOf(&mtt.Field3).Pointer(), reflect.TypeOf(mtt).Field(fi).Offset+reflect.ValueOf(&mtt).Pointer())
	}
}

func TestNewFieldsSubset(t *testing.T) {
	type MyTestType struct {
		Field1 string
		Field2 []byte
		Field3 *int
		Field4 float64
	}
	mtt := MyTestType{}
	expected := FieldSubset{Name: "TestSubSet", Type: reflect.TypeOf(mtt), fieldsIds: []int{1, 3}}
	actual, err := NewFieldsSubset("TestSubSet", &mtt, &mtt.Field2, &mtt.Field4)
	if err != nil {
		t.Log(err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("NewFieldsSubset returned %v \r\n while expected %v", actual, expected)
	}
}

func TestGetNewByReflect(t *testing.T) {
	{
		expected := new(int)
		actual := GetNewByReflect((*int)(nil))
		if _, ok := actual.(*int); !ok {
			t.Errorf("GetNewByReflect((*int)(nil)) returned %#v whileexpected %#v", actual, expected)
		}
	}
	{
		expected := (int)(0)
		actual := GetNewByReflect((int)(0))
		if _, ok := actual.(int); !ok {
			t.Errorf("GetNewByReflect((int)(0)) returned %#v whileexpected %#v", actual, expected)
		}
	}
	_ = GetNewByReflect((*string)(nil))
	_ = GetNewByReflect((*float32)(nil))
	_ = GetNewByReflect((*float64)(nil))
	_ = GetNewByReflect((*bool)(nil))
	_ = GetNewByReflect((*time.Time)(nil))
	_ = GetNewByReflect((*[]byte)(nil))
	_ = GetNewByReflect((*int64)(nil))
}

func BenchmarkMap(b *testing.B) {
	b.Skip()
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
	b.Skip()

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
	b.Skip()
	for n := 0; n <= b.N; n++ {
		s := ConvertIntToStringBySprintf(n)
		s += ""
	}
	b.ReportAllocs()
}

func BenchmarkConvertIntToStringByStrConv(b *testing.B) {
	b.Skip()
	for n := 0; n <= b.N; n++ {
		s := ConvertIntToStringByStrConv(n)
		s += ""
	}
	b.ReportAllocs()
}

func BenchmarkConvertFloatToStringBySprintf(b *testing.B) {
	b.Skip()
	for n := 0; n <= b.N; n++ {
		s := ConvertFloatToStringBySprintf(26535141592653.1415926)
		s += ""
	}
	b.ReportAllocs()
}

func BenchmarkConvertFloatToStringByStrConv(b *testing.B) {
	b.Skip()
	for n := 0; n <= b.N; n++ {
		s := ConvertFloatToStringByStrConv(26535141592653.1415926)
		s += ""
	}
	b.ReportAllocs()
}

func BenchmarkConvertInt64ToStringBySprintf(b *testing.B) {
	b.Skip()
	for n := 0; n <= b.N; n++ {
		s := ConvertInt64ToStringBySprintf(9223372036854775807)
		s += ""
	}
	b.ReportAllocs()
}

func BenchmarkConvertInt64ToStringByStrConv(b *testing.B) {
	b.Skip()
	for n := 0; n <= b.N; n++ {
		s := ConvertInt64ToStringByStrConv(9223372036854775807)
		s += ""
	}
	b.ReportAllocs()
}

func BenchmarkConvertTimeToStringBySprintf(b *testing.B) {
	b.Skip()
	for n := 0; n <= b.N; n++ {
		s := ConvertTimeToStringBySprintf(time.Now())
		s += ""
	}
	b.ReportAllocs()

}
func BenchmarkConvertTimeToStringByForma(b *testing.B) {
	b.Skip()
	for n := 0; n <= b.N; n++ {
		s := ConvertTimeToStringByFormat(time.Now())
		s += ""
	}
	b.ReportAllocs()
}
func BenchmarkConvertTimeToStringByStringer(b *testing.B) {
	b.Skip()
	for n := 0; n <= b.N; n++ {
		s := ConvertTimeToStringByStringer(time.Now())
		s += ""
	}
	b.ReportAllocs()
}

func BenchmarkGenerateStringsReplace(b *testing.B) {
	b.Skip()
	for n := 0; n <= b.N; n++ {
		s := GenerateStringsReplace("dbo.TestTable", "Col1, Col2, Col3", "Value1, Value3, Value4")
		s += ""
	}
	b.ReportAllocs()
}
func BenchmarkGenerateFmtSprintf(b *testing.B) {
	b.Skip()
	for n := 0; n <= b.N; n++ {
		s := GenerateFmtSprintf("dbo.TestTable", "Col1, Col2, Col3", "Value1, Value3, Value4")
		s += ""
	}
	b.ReportAllocs()
}
func BenchmarkGenerateCustom(b *testing.B) {
	b.Skip()
	for n := 0; n <= b.N; n++ {
		s := GenerateCustom("dbo.TestTable", "Col1, Col2, Col3", "Value1, Value3, Value4")
		s += ""
	}
	b.ReportAllocs()
}

func BenchmarkCustomJoinStrings(b *testing.B) {
	b.Skip()
	for n := 0; n <= b.N; n++ {
		s := CustomJoinStrings("a", "b", "c", "d", "e", "f", "j", "h", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "1", "22", "3", "5", "777", "8", "9", "0", "1234567890qwertyuiopasdfghjkzxcvbnm")
		_ = s
	}
	b.ReportAllocs()
}

func BenchmarkStringsJoinStrings(b *testing.B) {
	b.Skip()
	for n := 0; n <= b.N; n++ {
		s := strings.Join([]string{"a", "b", "c", "d", "e", "f", "j", "h", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "1", "22", "3", "5", "777", "8", "9", "0", "1234567890qwertyuiopasdfghjkzxcvbnm"}, ",")
		_ = s
	}
	b.ReportAllocs()
}

func BenchmarkReflectFieldById(b *testing.B) {
	b.Skip()
	type MyTestType struct {
		Field1 int
		Field2 int
		Field3 int
	}
	m := MyTestType{}
	t := reflect.TypeOf(m)
	b.Log(t.Field(1).Name)

	b.ResetTimer()
	for n := 0; n <= b.N; n++ {
		v := reflect.ValueOf(reflect.ValueOf(m).Field(1))
		_ = v
	}
	b.ReportAllocs()
}

func BenchmarkReflectFieldByIndex(b *testing.B) {
	b.Skip()
	type MyTestType struct {
		Field1 int
		Field2 int
		Field3 int
	}
	m := MyTestType{}
	t := reflect.TypeOf(m)
	b.Log(t.FieldByIndex([]int{1}).Name)

	b.ResetTimer()
	for n := 0; n <= b.N; n++ {
		v := reflect.ValueOf(reflect.ValueOf(m).FieldByIndex([]int{1}))
		_ = v
	}
	b.ReportAllocs()
}

func BenchmarkReflectFieldByName(b *testing.B) {
	b.Skip()
	type MyTestType struct {
		Field1 int
		Field2 int
		Field3 int
	}
	m := MyTestType{}
	t := reflect.TypeOf(m)
	x, _ := t.FieldByName("Field2")
	b.Log(x.Name)

	b.ResetTimer()
	for n := 0; n <= b.N; n++ {
		v := reflect.ValueOf(reflect.ValueOf(m).FieldByName("Field2"))
		_ = v
	}
	b.ReportAllocs()
}

func BenchmarkGetPointerValueInterface(b *testing.B) {
	b.Skip()
	s := "TestString"
	mptt := MyTestPtrType{S: &s}
	b.ResetTimer()
	for n := 0; n <= b.N; n++ {
		_ = GetPointerValueInterface(&mptt.S)
	}
	b.ReportAllocs()
}

func BenchmarkGetPointerValueString(b *testing.B) {
	b.Skip()
	s := "TestString"
	mptt := MyTestPtrType{S: &s}
	b.ResetTimer()
	for n := 0; n <= b.N; n++ {
		_ = GetPointerValueString(&mptt.S)
	}
	b.ReportAllocs()
}

func BenchmarkGetNewBySwitchNew(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		_ = GetNewBySwitchNew((*int)(nil))
		_ = GetNewBySwitchNew((*string)(nil))
		_ = GetNewBySwitchNew((*float32)(nil))
		_ = GetNewBySwitchNew((*float64)(nil))
		_ = GetNewBySwitchNew((*bool)(nil))
		_ = GetNewBySwitchNew((*time.Time)(nil))
		_ = GetNewBySwitchNew((*[]byte)(nil))
		_ = GetNewBySwitchNew((*int64)(nil))
	}
	b.ReportAllocs()
}

func BenchmarkGetNewBySwitch(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		_ = GetNewBySwitch((*int)(nil))
		_ = GetNewBySwitch((*string)(nil))
		_ = GetNewBySwitch((*float32)(nil))
		_ = GetNewBySwitch((*float64)(nil))
		_ = GetNewBySwitch((*bool)(nil))
		_ = GetNewBySwitch((*time.Time)(nil))
		_ = GetNewBySwitch((*[]byte)(nil))
		_ = GetNewBySwitch((*int64)(nil))
	}
	b.ReportAllocs()
}

func BenchmarkGetNewByReflect(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		_ = GetNewByReflect((*int)(nil))
		_ = GetNewByReflect((*string)(nil))
		_ = GetNewByReflect((*float32)(nil))
		_ = GetNewByReflect((*float64)(nil))
		_ = GetNewByReflect((*bool)(nil))
		_ = GetNewByReflect((*time.Time)(nil))
		_ = GetNewByReflect((*[]byte)(nil))
		_ = GetNewByReflect((*int64)(nil))
	}
	b.ReportAllocs()
}

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
