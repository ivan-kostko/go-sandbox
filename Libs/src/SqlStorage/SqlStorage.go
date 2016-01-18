// SqlStorage project SqlStorage.go
package SqlStorage

import (
	"fmt"

	"Logger"
	"SqlStorage/SqlDialects"

	. "commonInterfaces"
	. "customErrors"

	"github.com/jmoiron/sqlx" // opensource Sql Extentions lib.

	_ "github.com/alexbrainman/odbc" // mssql or freetds. Registred as odbc
	_ "github.com/lib/pq"            // postgreSql. Registred as postrgres
)

const (
	ERR_FAILEDTOINITIALIZE_SQLSTORAGE = "Failed to initialize SqlStorage due to error "
)

type SqlStorageConfiguration struct {
	DriverName   string
	ConnString   string
	DialectAlias string
}

// Represents generic abstract Sql storage
//
type SqlStorage struct {
	log     Logger.ILogger
	conf    *SqlStorageConfiguration
	dialect SqlDialects.ISqlDialect
	conn    *sqlx.DB // keeps connection pool active
}

// The interface represents embeded composition of implemented interfaces by SqlStorage struct
// and general SqlStorage functionality
type ISqlStorage interface {
	Initializer
	MustInitializer
	Disposer
}

// Generic (I)SqlStorage factory.
// Tries to get default Logger if log is nil
// Returns *Error in case of nil configuration or problems on obtaining Logger or error on Initialization
func GetNewISqlStorage(conf SqlStorageConfiguration, log Logger.ILogger) (ISqlStorage, *Error) {
	if log == nil {
		// try to get default logger
		log = Logger.GetILogger(Logger.LoggerConfig{})
	}
	iss := &SqlStorage{log: log, conf: &conf}
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
	ss.conn, err = sqlx.Connect(ss.conf.DriverName, ss.conf.ConnString)
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
	if ss.conn != nil {
		ss.conn.Close()
		ss.conn = nil
	}
	ss.conf = nil
	// The dialect shouldn't be disposed by its own, cause it could be in use by other objects. Just remove reference to it
	ss.dialect = nil
	// The log shouldn't be disposed by its own, cause it could be in use by other objects. Just remove reference to it
	ss.log = nil
}
