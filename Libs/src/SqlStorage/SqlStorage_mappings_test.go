// The file contains SqlStorage mapping functionality tests
package SqlStorage

import (
	"reflect"
	"strings"
	"testing"

	"Logger"
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
	//expectedVal := []interface{}{"Field1", []byte("Field2"), float64(2345.98765)}
	expectedPtr := []interface{}{&mtt.Field1, &mtt.Field2, &mtt.Field4}
	k, err := NewKey("TestKeyExtractFields", &sample, &sample.Field1, &sample.Field2, &sample.Field4)
	if err != nil {
		t.Errorf("NewKey returned *Error %v while expected %v", err, nil)
	}
	actualStr, actualVal, err := k.ExtractFrom(&mtt)
	if !(reflect.DeepEqual(actualStr, expectedStr) &&
		reflect.DeepEqual(actualVal, expectedPtr) &&
		err == nil) {
		t.Errorf("k.Extract(&mtt) returned Names: %v Vals: %v Error %v while expected Names: %v Vals: %v Error %v", actualStr, actualVal, err, expectedStr, expectedPtr, nil)
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
	expectedVal := []interface{}{(*int)(nil), []byte("Field2"), float64(2345.98765)}
	//expectedPtr := []interface{}{&mtt.Field3, &mtt.Field2, &mtt.Field4}
	k, err := NewKey("TestKeyExtractFields", &sample, &sample.Field3, &sample.Field2, &sample.Field4)
	if err != nil {
		t.Errorf("NewKey returned *Error %v while expected %v", err, nil)
	}
	actualStr, actualVal, err := k.ExtractFrom(mtt)
	if !(reflect.DeepEqual(actualStr, expectedStr) &&
		reflect.DeepEqual(actualVal, expectedVal) &&
		err == nil) {
		t.Errorf("k.Extract(mtt) returned Names: %v Vals: %v Error %v \r\n\t\t\twhile expected Names: %v Vals: %v Error %v", actualStr, actualVal, err, expectedStr, expectedVal, nil)
	}

}

func TestKeyAssignFieldsByVals(t *testing.T) {
	type MyTestType struct {
		Field1 string
		Field2 []byte
		Field3 *int
		Field4 float64
	}
	sample := MyTestType{}
	actualStruct := MyTestType{Field1: "Field1", Field2: []byte("Field2"), Field3: nil, Field4: float64(2345.98765)}
	expectedStruct := MyTestType{Field1: "NewField1Value", Field2: []byte("AKind Of New Field2"), Field3: nil, Field4: float64(54321.56789)}
	k, err := NewKey("TestKeyExtractFields", &sample, &sample.Field1, &sample.Field2, &sample.Field4)
	if err != nil {
		t.Errorf("NewKey returned *Error %v while expected %v", err, nil)
	}
	err = k.AssignTo(&actualStruct, []interface{}{"NewField1Value", []byte("AKind Of New Field2"), float64(54321.56789)})
	if !(reflect.DeepEqual(actualStruct, expectedStruct) && err == nil) {
		t.Errorf("k.Assign(&mtt,...) assigned : %v with Error %v \r\n\t\t\t while expected %v with Error %v", actualStruct, err, expectedStruct, nil)
	}

}

func TestSqlStorageGenerateStructureMapping(t *testing.T) {
	type MyTestType struct {
		Id     *int    `db:"{'ColName':'Id', 'Keys':['PK'], 'ResolvedByDb':true}"`
		Field1 *string `db:"{'ColName':'field_1', 'Keys':['BK'], 'ResolvedByDb':false}"`
		Field2 *int    `db:"{'ColName':'field_2', 'Keys':['BK'], 'ResolvedByDb':false}"`
		Field3 *int    `db:"{'ColName':'field_3', 'Keys':['BK', 'Val'], 'ResolvedByDb':false}"`
		Field4 float64 `db:"{'ColName':'field_4', 'Keys':['Val'], 'ResolvedByDb':false}"`
		Field5 float64 `db:"{'ColName':'field_5', 'Keys':['Val'], 'ResolvedByDb':false}"`
		Field6 float64 `db:"{'ColName':'field_6', 'Keys':['Val'], 'ResolvedByDb':false}"`
	}
	typ := reflect.TypeOf(MyTestType{})
	ssc := SqlStorageConfiguration{
		MappingTag: "db",
	}
	ss := SqlStorage{
		conf: &ssc,
		log:  Logger.GetStdTerminalLogger(),
		namesMatch: func(a, b string) bool {
			return strings.ToLower(a) == strings.ToLower(b)
		},
		getStorageObjectFields: func(ss *SqlStorage, name string) ([]StorageObjectField, *Error) {
			return []StorageObjectField{{Name: "Id"}, {Name: "field_1"}, {Name: "field_2"}, {Name: "field_3"}, {Name: "field_4"}, {Name: "field_5"}, {Name: "field_6"}}, nil
		},
	}
	initStorObjName := "TestStorageObject"
	expectedStructMap := StructureMapping{
		StorageObjectName: initStorObjName,
		FieldsMappings: []FieldMapping{
			{
				StorageObjectFieldName: "Id",
				StructureFieldName:     "Id",
				StructureFieldId:       0,
				ParticipateInKeys:      []string{EMPTY_STRING, "PK"},
			},
			{
				StorageObjectFieldName: "field_1",
				StructureFieldName:     "Field1",
				StructureFieldId:       1,
				ParticipateInKeys:      []string{EMPTY_STRING, "BK"},
			},
			{
				StorageObjectFieldName: "field_2",
				StructureFieldName:     "Field2",
				StructureFieldId:       2,
				ParticipateInKeys:      []string{EMPTY_STRING, "BK"},
			},

			{
				StorageObjectFieldName: "field_3",
				StructureFieldName:     "Field3",
				StructureFieldId:       3,
				ParticipateInKeys:      []string{EMPTY_STRING, "BK", "Val"},
			},

			{
				StorageObjectFieldName: "field_4",
				StructureFieldName:     "Field4",
				StructureFieldId:       4,
				ParticipateInKeys:      []string{EMPTY_STRING, "Val"},
			},

			{
				StorageObjectFieldName: "field_5",
				StructureFieldName:     "Field5",
				StructureFieldId:       5,
				ParticipateInKeys:      []string{EMPTY_STRING, "Val"},
			},

			{
				StorageObjectFieldName: "field_6",
				StructureFieldName:     "Field6",
				StructureFieldId:       6,
				ParticipateInKeys:      []string{EMPTY_STRING, "Val"},
			},
		},
		KeyMappings: map[string]KeyMapping{
			"PK": KeyMapping{
				Key: Key{
					Name:        "PK",
					Type:        typ,
					fieldsIds:   []int{0},
					fieldsNames: []string{"Id"},
				},
				SOFieldsNames: []string{"Id"},
			},
			"BK": KeyMapping{
				Key: Key{
					Name:        "BK",
					Type:        typ,
					fieldsIds:   []int{1, 2, 3},
					fieldsNames: []string{"Field1", "Field2", "Field3"},
				},
				SOFieldsNames: []string{"field_1", "field_2", "field_3"},
			},
			"Val": KeyMapping{
				Key: Key{
					Name:        "Val",
					Type:        typ,
					fieldsIds:   []int{3, 4, 5, 6},
					fieldsNames: []string{"Field3", "Field4", "Field5", "Field6"},
				},
				SOFieldsNames: []string{"field_3", "field_4", "field_5", "field_6"},
			},
			"": KeyMapping{
				Key: Key{
					Name:        "",
					Type:        typ,
					fieldsIds:   []int{0, 1, 2, 3, 4, 5, 6},
					fieldsNames: []string{"Id", "Field1", "Field2", "Field3", "Field4", "Field5", "Field6"},
				},
				SOFieldsNames: []string{"Id", "field_1", "field_2", "field_3", "field_4", "field_5", "field_6"},
			},
		},
	}
	actualStructMap, actualErr := ss.generateStructureMapping(initStorObjName, typ)
	if actualErr != nil {
		t.Fatalf("ss.generateStructureMapping(initStorObjName, typ) returned unexpected error %v", actualErr)
	}
	if !(reflect.DeepEqual(*actualStructMap, expectedStructMap)) {
		t.Errorf("ss.generateStructureMapping(initStorObjName, typ) returned mapping: \r\n%v \r\n while expected \r\n%v", *actualStructMap, expectedStructMap)

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
		_, _, _ = k.ExtractFrom(&mtt)
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
		_, _, _ = k.ExtractFrom(&mtt)
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
		_, _, _ = k.ExtractFrom(&mtt)
	}
	b.ReportAllocs()
}
