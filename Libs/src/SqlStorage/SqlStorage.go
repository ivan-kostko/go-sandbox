// SqlStorage project SqlStorage.go
package SqlStorage

import (
	"fmt"
	"reflect"

	"Logger"
	"SqlStorage/SqlDialects"

	. "commonInterfaces"
	. "customErrors"
)

const (
	ERR_FAILEDTOINITIALIZE_SQLSTORAGE = "Failed to initialize SqlStorage due to error "
	ERR_FAILEDTOGETSTORAGEFIELDS      = "Failed to get storage object foelds"
	ERR_FUNCNOTYETIMPLEMENTED         = "The function is not yet implemented"
)

type SqlStorageConfiguration struct {
	DriverName   string
	ConnString   string
	DialectAlias string
	MappingTag   string
}

// Represents generic abstract Sql storage
//
type SqlStorage struct {
	log     Logger.ILogger
	conf    *SqlStorageConfiguration
	dialect SqlDialects.ISqlDialect
	db      *SqlDatabase // keeps connection pool active

	structureMappings      map[reflect.Type]StructureMapping
	namesMatch             func(string, string) bool
	getStorageObjectFields func(*SqlStorage, string) ([]StorageObjectField, *Error)
}

// The interface represents embeded composition of implemented interfaces by SqlStorage struct
// and general SqlStorage functionality
type ISqlStorage interface {
	// Implements:
	Initializer
	MustInitializer
	Disposer

	IStorage
}

// Generic (I)SqlStorage factory.
// Tries to get default Logger if log is nil
// Returns *Error in case of nil configuration or problems on obtaining Logger or error on Initialization
func GetNewISqlStorage(conf SqlStorageConfiguration, log Logger.ILogger) (ISqlStorage, *Error) {
	if log == nil {
		// try to get default logger by providing empty LoggerConfig
		log = Logger.GetILogger(Logger.LoggerConfig{})
	}
	iss := &SqlStorage{log: log, conf: &conf, getStorageObjectFields: getStorageObjectFields}
	err := iss.Initialize()
	if err != nil {
		log.Critical(ERR_FAILEDTOINITIALIZE_SQLSTORAGE, err)
		return nil, err
	}
	return iss, nil
}

// Tries to set up SqlStorage instance according to conf. Returns Error on failure
// Implementatio of Initializer interface
func (ss *SqlStorage) Initialize() *Error {
	var derr *Error
	ss.dialect, derr = SqlDialects.GetDialectByAlias(ss.conf.DialectAlias)
	if derr != nil {
		return derr
	}
	var err error
	ss.db, err = GetNewSqlDatabase(ss.conf.DriverName, ss.conf.ConnString)
	if err != nil {
		return NewError(InvalidOperation, fmt.Sprintf("Failed to connect via driver %v to Sql %v due to error : %v", ss.conf.DriverName, ss.conf.ConnString, err))
	}
	return nil
}

// Tries to set up SqlStorage instance according to conf. Panics on failure
// Implementatio of MustInitializer interface
func (ss *SqlStorage) MustInitialize() {
	err := ss.Initialize()
	if err != nil {
		ss.log.Critical(ERR_FAILEDTOINITIALIZE_SQLSTORAGE, err)
		panic(ERR_FAILEDTOINITIALIZE_SQLSTORAGE)
	}
}

// Prepare SqlStorage to be garbage collected
func (ss *SqlStorage) Dispose() {
	if ss.db != nil {
		ss.db.Close()
		ss.db = nil
	}
	ss.conf = nil
	// The dialect shouldn't be disposed by its own, cause it could be in use by other objects. Just remove reference to it
	ss.dialect = nil
	// The log shouldn't be disposed by its own, cause it could be in use by other objects. Just remove reference to it
	ss.log = nil
}

func (ss *SqlStorage) Get(m interface{}) *Error {
	// TODO : Implement
	return NewError(Nonsupported, ERR_FUNCNOTYETIMPLEMENTED)
}

func (ss *SqlStorage) GetKeyByKey(m interface{}, getKeyName, byKeyName string) *Error {
	typ := reflect.TypeOf(m)
	sm := ss.structureMappings[typ]
	getKeyMapping := sm.KeyMappings[getKeyName]
	byKeyMapping := sm.KeyMappings[byKeyName]
	whereCols := byKeyMapping.SOFieldsNames
	whereColumnList := make([]SqlDialects.SqlScriptString, len(whereCols))
	for i, c := range whereCols {
		whereColumnList[i] = SqlDialects.SqlScriptString(c)
	}
	_, whereVals, err := byKeyMapping.Extract(m)
	if err != nil {
		return err
	}
	whereValueList := make([]SqlDialects.SqlScriptString, len(whereVals))
	for i, v := range whereVals {
		whereValueList[i], err = ss.dialect.ConvertIntoSqlScriptString(v)
		if err != nil {
			return err
		}
	}
	whereCondition := ss.dialect.BuildWhereSqlScriptString(whereColumnList, whereValueList)
	getColumnList := ss.dialect.BuildColumnsListSqlScriptString(getKeyMapping.SOFieldsNames)
	_, getVals, err := getKeyMapping.Extract(m)
	selectQuery := ss.dialect.BuildSelectSqlScriptString(SqlDialects.SqlScriptString(sm.StorageObjectName), getColumnList, whereCondition, 1)
	// TODO : Scan Rows
	queryErr := ss.db.QueryIntoSlice(string(selectQuery), getVals)
	if queryErr != nil {
		ss.log.Criticalf("Query %v failed with error: %v", selectQuery, queryErr)
	}

	return nil
}

func (ss *SqlStorage) Put(m interface{}) *Error {
	// TODO : Implement
	return NewError(Nonsupported, ERR_FUNCNOTYETIMPLEMENTED)
}

func (ss *SqlStorage) Resolve(m interface{}) *Error {
	// TODO : Implement
	return NewError(Nonsupported, ERR_FUNCNOTYETIMPLEMENTED)
}

func (ss *SqlStorage) Del(m interface{}) *Error {
	// TODO : Implement
	return NewError(Nonsupported, ERR_FUNCNOTYETIMPLEMENTED)
}

type StorageObjectField struct {
	Name string
}

// Gets the storage object fields
func (ss *SqlStorage) GetStorageObjectFields(stObjName string) ([]StorageObjectField, *Error) {
	return ss.getStorageObjectFields(ss, stObjName)
}

func getStorageObjectFields(ss *SqlStorage, stObjName string) ([]StorageObjectField, *Error) {
	r, err := ss.db.Query(string(ss.dialect.BuildSelectAllColumnsSqlScriptString(stObjName)), nil)
	if err != nil {
		ss.log.Warning(err.Error())
		return nil, NewError(InvalidOperation, ERR_FAILEDTOGETSTORAGEFIELDS)
	}
	if r == nil {
		ss.log.Info("SelectAllColumns returned NIL result")
		return nil, NewError(InvalidOperation, ERR_FAILEDTOGETSTORAGEFIELDS)
	}
	cols, colErr := r.Columns()
	if colErr != nil {
		ss.log.Warning(colErr.Error())
		return nil, NewError(InvalidOperation, ERR_FAILEDTOGETSTORAGEFIELDS)
	}
	res := make([]StorageObjectField, len(cols))
	for i, col := range cols {
		res[i] = StorageObjectField{Name: col}
	}
	return res, nil
}
