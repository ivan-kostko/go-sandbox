package SqlStorage

import (
	"Logger"
	"SqlStorage/SqlDialects"
	"customErrors"
	"reflect"
	"strings"
	"testing"
)

const (
	TEST_CONN_STR = `driver=Sql Server Native Client 11.0;server=.\MSSQL_2014_DEV;uid=godev;pwd=godev;database=Test`
)

func TestSqlStorageInitializeWithFakeDialect(t *testing.T) {
	ssc := SqlStorageConfiguration{
		DriverName:   "odbc",
		ConnString:   TEST_CONN_STR,
		DialectAlias: "FakeDialect",
	}
	l := Logger.GetStdTerminalLogger()
	ss := SqlStorage{conf: &ssc, log: l}
	expectedError := customErrors.NewError(customErrors.Nonsupported, SqlDialects.ERR_DIALECTNOTFOUND)
	err := ss.Initialize()
	if *err != *expectedError {
		t.Errorf("SqlStorage.Initialize failed with error %s when expected %s", err.Error(), expectedError.Error())
	}
}

func TestSqlStorageInitializeWithEmptyConnection(t *testing.T) {
	ssc := SqlStorageConfiguration{
		DriverName: "",
		ConnString: "",
		// TODO : replace by default dialect when implemented
		DialectAlias: SqlDialects.MSSQL2014_DIALECT_REGISTER_NAME,
	}
	l := Logger.GetStdTerminalLogger()
	ss := SqlStorage{conf: &ssc, log: l}
	expectedError := customErrors.NewError(customErrors.InvalidOperation, `Failed to connect via driver  to Sql  due to error : customErrors.Error{Type:InvalidOperation, Message:sql: unknown driver "" (forgotten import?)}`)
	err := ss.Initialize()
	if *err != *expectedError {
		t.Errorf("SqlStorage.Initialize failed with error %s when expected %s", err.Error(), expectedError.Error())
	}
}

// NB: Does integration - tries to connect to TEST_CONN_STR via odbc driver
func TestSqlStorageInitializeForSuccess(t *testing.T) {
	t.Skip()
	ssc := SqlStorageConfiguration{
		DriverName: "odbc",
		ConnString: TEST_CONN_STR,
		// TODO : replace by default dialect when implemented
		DialectAlias: SqlDialects.MSSQL2014_DIALECT_REGISTER_NAME,
	}
	l := Logger.GetStdTerminalLogger()
	ss := SqlStorage{conf: &ssc, log: l}
	err := ss.Initialize()
	if err != nil {
		t.Errorf("SqlStorage.Initialize failed with error %s when expected success", err.Error())
	}
	// Check initialized values
	expectedDialect, _ := SqlDialects.GetDialectByAlias(ssc.DialectAlias)
	if ss.dialect != expectedDialect {
		t.Errorf("SqlStorage.Initialize sets up dialect %#v when expected %#v", ss.dialect, expectedDialect)
	}
	if ss.db == nil {
		t.Errorf("SqlStorage.Initialize sets up connection as nil")
	}
	if pingErr := ss.db.Ping(); err != nil {
		t.Errorf("Ping connection to initialized SqlStorage failed with error: %v", pingErr)
	}
}

// Func tests

type MyTestType struct {
	Id     *int    `db:"{'ColName':'Id', 'Keys':['PK'], 'ResolvedByDb':true}"`
	Field1 *string `db:"{'ColName':'field_1', 'Keys':['BK'], 'ResolvedByDb':false}"`
	Field2 *int    `db:"{'ColName':'field_2', 'Keys':['BK'], 'ResolvedByDb':false}"`
	Field3 *int    `db:"{'ColName':'field_3', 'Keys':['BK', 'Val'], 'ResolvedByDb':false}"`
	Field4 float64 `db:"{'ColName':'field_4', 'Keys':['Val'], 'ResolvedByDb':false}"`
	Field5 float64 `db:"{'ColName':'field_5', 'Keys':['Val'], 'ResolvedByDb':false}"`
	Field6 float64 `db:"{'ColName':'field_6', 'Keys':['Val'], 'ResolvedByDb':false}"`
}

// Test helper function
func GetTestStorage() SqlStorage {
	typ := reflect.TypeOf(MyTestType{})
	ssc := SqlStorageConfiguration{
		MappingTag: "db",
	}
	ss := SqlStorage{
		conf:    &ssc,
		log:     Logger.GetStdTerminalLogger(),
		dialect: SqlDialects.GetMsSql2014Dialect(),
		namesMatch: func(a, b string) bool {
			return strings.ToLower(a) == strings.ToLower(b)
		},
		getStorageObjectFields: func(ss *SqlStorage, name string) ([]StorageObjectField, *customErrors.Error) {
			return []StorageObjectField{{Name: "Id"}, {Name: "field_1"}, {Name: "field_2"}, {Name: "field_3"}, {Name: "field_4"}, {Name: "field_5"}, {Name: "field_6"}}, nil
		},
		structureMappings: make(map[reflect.Type]StructureMapping),
	}
	initStorObjName := "TestStorageObject"
	sm := StructureMapping{
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
	ss.structureMappings[typ] = sm
	return ss
}

func TestGetKeyByKey(t *testing.T) {
	ss := GetTestStorage()
	field1Val := "Field1"
	field2Val := 45678
	mtt := MyTestType{
		Id:     nil,
		Field1: &field1Val,
		Field2: &field2Val,
		Field3: nil,
		Field4: 123.456,
		Field5: 654.321,
		Field6: 0,
	}
	ss.GetKeyByKey(mtt, "", "BK")
}
