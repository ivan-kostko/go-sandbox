// // +build integration

// keep empty line between build tag and comment
// to run integration tests execute 'go test -tags=integration'

package SqlStorage

import (
	"Logger"
	"reflect"
	"testing"
	"time"
)

const TEST_DRIVER = "odbc"
const TEST_DSN = `driver=SQL Server;server=.\MSSQL_2014_DEV;uid=godev;pwd=godev;database=Test`
const TEST_FREETDS_CONNSTR = `driver=FreeTDS;Server=127.0.0.1;Port=1433;UID=godev;PWD=godev;DB=Test;`
const TEST_STORAGE_OBJ_NAME = "dbo.TypesTest"

type MSSQLTestType struct {
	Field1  *int       `db:"{'ColName':'Col_INT','Keys':['INT']}"`
	Field2  *int64     `db:"{'ColName':'Col_BIGINT','Keys':['BIGINT']}"`
	Field3  *float64   `db:"{'ColName':'Col_DECIMAL','Keys':['DECIMAL']}"`
	Field4  *bool      `db:"{'ColName':'Col_BIT','Keys':['VAL']}"`
	Field5  *float32   `db:"{'ColName':'Col_NUMERIC','Keys':['VAL']}"`
	Field6  *float64   `db:"{'ColName':'Col_FLOAT','Keys':['VAL']}"`
	Field7  *time.Time `db:"{'ColName':'Col_DATE','Keys':['TIME']}"`
	Field8  *time.Time `db:"{'ColName':'Col_TIME','Keys':['TIME']}"`
	Field9  *time.Time `db:"{'ColName':'Col_DATETIME','Keys':['TIME']}"`
	Field10 *time.Time `db:"{'ColName':'Col_DATETIME2','Keys':['TIME']}"`
	Field11 *time.Time `db:"{'ColName':'Col_SMALLDATETIME','Keys':['TIME']}"`
	Field12 *time.Time `db:"{'ColName':'Col_DATETIMEOFFSET','Keys':['TIME']}"`
	Field14 *string    `db:"{'ColName':'Col_NVARCHAR','Keys':['STRING']}"`
	Field15 *[]byte    `db:"{'ColName':'Col_VARBINARY','Keys':['BYTE']}"`
}

func GetTestSqlStorageConfiguration() SqlStorageConfiguration {
	return SqlStorageConfiguration{
		DriverName:   TEST_DRIVER,
		ConnString:   TEST_FREETDS_CONNSTR,
		DialectAlias: "mssql2014",
		MappingTag:   "db",
	}
}

func GetNewMSSQL2014Storage() *SqlStorage {
	log := Logger.GetStdTerminalLogger()
	ss, err := GetNewSqlStorage(GetTestSqlStorageConfiguration(), log)
	if err != nil {
		log.Critical(err)
	}
	return ss
}

func TestSqlDatabaseConnect(t *testing.T) {
	sdb, err := GetNewSqlDatabase(TEST_DRIVER, TEST_FREETDS_CONNSTR)
	if err != nil {
		t.Log(err)
	}
	sdb.Ping()
	t.Log(*sdb)
}

func TestMSSQLStorageGetKeyByKey(t *testing.T) {
	//t.Skip()
	ctime, _ := time.Parse("20060102 15:04:05.99999999 Z07:00", "20160131 09:05:33.1100000 +00:00")

	//t.Log(ctime)

	var _uint32 int = 2147483646
	var _big_Int int64 = 9223372036854775806
	var _big_Float float64 = 12345.1234567
	var _bool bool = true
	var _float32 float32 = 12345.1234567
	var _float64 float64 = 123456789.123456789
	var _DATE time.Time = ctime.Round(time.Hour)
	var _TIME time.Time = ctime.Round(100 * time.Nanosecond)
	var _DATETIME time.Time = ctime.Round(1000000 * time.Nanosecond)
	var _DATETIME2 time.Time = ctime.Round(100 * time.Nanosecond)
	var _SMALLDATETIME time.Time = ctime.Round(time.Second)
	var _DATETIMEOFFSET time.Time = ctime.Round(100 * time.Nanosecond)
	var _NVARCHAR string = `driver=SQL Server;server=.\MSSQL_2014_DEV;uid=godev;pwd=godev;database=Test`
	var _slbyte []byte = ([]byte)(`driver=SQL Server;server=.\MSSQL_2014_DEV;uid=godev;pwd=godev;database=Test`)
	expectedMtt := MSSQLTestType{
		Field1:  &_uint32,
		Field2:  &_big_Int,
		Field3:  &_big_Float,
		Field4:  &_bool,
		Field5:  &_float32,
		Field6:  &_float64,
		Field7:  &_DATE,
		Field8:  &_TIME,
		Field9:  &_DATETIME,
		Field10: &_DATETIME2,
		Field11: &_SMALLDATETIME,
		Field12: &_DATETIMEOFFSET,
		Field14: &_NVARCHAR,
		Field15: &_slbyte,
	}

	sample := MSSQLTestType{}
	ss := GetNewMSSQL2014Storage()
	defer ss.Dispose()
	err := ss.RegisterType(TEST_STORAGE_OBJ_NAME, sample)
	if err != nil {
		t.Logf("ss.RegisterType returned error %v", err)
	}

	actualMtt := MSSQLTestType{
		Field1:  new(int),
		Field2:  &_big_Int,
		Field3:  new(float64),
		Field4:  new(bool),
		Field5:  new(float32),
		Field6:  new(float64),
		Field7:  new(time.Time),
		Field8:  new(time.Time),
		Field9:  new(time.Time),
		Field10: new(time.Time),
		Field11: new(time.Time),
		Field12: new(time.Time),
		Field14: new(string),
		Field15: new([]byte),
	}

	err = ss.GetKeyByKey(actualMtt, "BYTE", "BIGINT")
	if err != nil {
		t.Logf("ss.GetKeyByKey returned error %v", err)
	}
	if (string)(*actualMtt.Field15) != (string)(*expectedMtt.Field15) {
		t.Errorf("ss.GetKeyByKey returned %v \r\n\t\t\t\t while expected %v", (string)(*actualMtt.Field15), (string)(*expectedMtt.Field15))
	}

	err = ss.GetKeyByKey(actualMtt, "STRING", "BIGINT")
	if err != nil {
		t.Logf("ss.GetKeyByKey returned error %v", err)
	}
	if (string)(*actualMtt.Field14) != (string)(*expectedMtt.Field14) {
		t.Errorf("ss.GetKeyByKey returned %v \r\n\t\t\t\t while expected %v", (string)(*actualMtt.Field14), (string)(*expectedMtt.Field14))
	}

	err = ss.GetKeyByKey(actualMtt, "INT", "BIGINT")
	if err != nil {
		t.Logf("ss.GetKeyByKey returned error %v", err)
	}
	if (*actualMtt.Field1) != (*expectedMtt.Field1) {
		t.Errorf("ss.GetKeyByKey returned %v \r\n\t\t\t\t while expected %v", (*actualMtt.Field1), (*expectedMtt.Field1))
	}

	err = ss.GetKeyByKey(actualMtt, "DECIMAL", "BIGINT")
	if err != nil {
		t.Logf("ss.GetKeyByKey returned error %v", err)
	}
	if (*actualMtt.Field3) != (*expectedMtt.Field3) {
		t.Errorf("ss.GetKeyByKey returned %v \r\n\t\t\t\t while expected %v", (*actualMtt.Field3), (*expectedMtt.Field3))
	}

	err = ss.GetKeyByKey(actualMtt, "VAL", "BIGINT")
	if err != nil {
		t.Logf("ss.GetKeyByKey returned error %v", err)
	}
	if (*actualMtt.Field4) != (*expectedMtt.Field4) {
		t.Errorf("ss.GetKeyByKey returned %v \r\n\t\t\t\t while expected %v", (*actualMtt.Field4), (*expectedMtt.Field4))
	}
	if (*actualMtt.Field5) != (*expectedMtt.Field5) {
		t.Errorf("ss.GetKeyByKey returned %v \r\n\t\t\t\t while expected %v", (*actualMtt.Field5), (*expectedMtt.Field5))
	}
	if (*actualMtt.Field6) != (*expectedMtt.Field6) {
		t.Errorf("ss.GetKeyByKey returned %v \r\n\t\t\t\t while expected %v", (*actualMtt.Field6), (*expectedMtt.Field6))
	}

	err = ss.GetKeyByKey(actualMtt, "TIME", "BIGINT")
	if err != nil {
		t.Logf("ss.GetKeyByKey returned error %v", err)
	}
	if (*actualMtt.Field7) != (*expectedMtt.Field7) {
		t.Errorf("ss.GetKeyByKey returned %v \r\n\t\t\t\t while expected %v", (*actualMtt.Field7), (*expectedMtt.Field7))
	}
	if (*actualMtt.Field8) != (*expectedMtt.Field8) {
		t.Errorf("ss.GetKeyByKey returned %v \r\n\t\t\t\t while expected %v", (*actualMtt.Field8), (*expectedMtt.Field8))
	}
	if (*actualMtt.Field9) != (*expectedMtt.Field9) {
		t.Errorf("ss.GetKeyByKey returned %v \r\n\t\t\t\t while expected %v", (*actualMtt.Field9), (*expectedMtt.Field9))
	}
	if (*actualMtt.Field10) != (*expectedMtt.Field10) {
		t.Errorf("ss.GetKeyByKey returned %v \r\n\t\t\t\t while expected %v", (*actualMtt.Field10), (*expectedMtt.Field10))
	}
	if (*actualMtt.Field11) != (*expectedMtt.Field11) {
		t.Errorf("ss.GetKeyByKey returned %v \r\n\t\t\t\t while expected %v", (*actualMtt.Field11), (*expectedMtt.Field11))
	}
	if (*actualMtt.Field12) != (*expectedMtt.Field12) {
		t.Errorf("ss.GetKeyByKey returned %v \r\n\t\t\t\t while expected %v", (*actualMtt.Field12), (*expectedMtt.Field12))
	}

	err = ss.GetKeyByKey(actualMtt, "", "BIGINT")
	if err != nil {
		t.Logf("ss.GetKeyByKey returned error %v", err)
	}
	if reflect.DeepEqual(actualMtt, expectedMtt) {
		t.Errorf("ss.GetKeyByKey returned %v \r\n\t\t\t\t while expected %v", (actualMtt), (expectedMtt))
	}

	//	ss.Dispose()
}
