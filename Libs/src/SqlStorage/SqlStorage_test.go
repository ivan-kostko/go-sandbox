package SqlStorage

import (
	"Logger"
	"SqlStorage/SqlDialects"
	"customErrors"
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
	expectedError := customErrors.NewError(customErrors.InvalidOperation, `Failed to connect via driver  to Sql  due to error : sql: unknown driver "" (forgotten import?)`)
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
	if ss.conn == nil {
		t.Errorf("SqlStorage.Initialize sets up connection as nil")
	}
	if pingErr := ss.conn.Ping(); err != nil {
		t.Errorf("Ping connection to initialized SqlStorage failed with error: %v", pingErr)
	}
}
