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

type SqlStorageConfiguration struct {
	DriverName  string
	ConnString  string
	DialectName string
}

type SqlStorage struct {
	log     Logger.ILogger
	conf    *SqlStorageConfiguration
	dialect SqlDialects.SqlDialect
	conn    *sqlx.DB // keeps connection pool active
}

// The interface represents embeded composition of implemented interfaces by SqlStorage struct
type ISqlStorage interface {
	Initializer
	MustInitializer
	Disposer
}

func (ss *SqlStorage) Initialize() *Error {
	var ok bool
	ss.dialect, ok = supportedSqlDialects[ss.conf.DialectName]
	if !ok {
		return NewError(Nonsupported, fmt.Sprintf("The dialect %s is not supported", ss.conf.DialectName))
	}
	var err error
	ss.conn, err = sqlx.Connect(ss.conf.DriverName, ss.conf.ConnString)
	if err != nil {
		return NewError(InvalidOperation, fmt.Sprintf("Failed to connect via driver %v to Sql %v due to error : %v", ss.conf.DriverName, ss.conf.ConnString, err))
	}
	return nil
}

func (ss *SqlStorage) Dispose() {
	if ss.conn != nil {
		ss.conn.Close()
		ss.conn = nil
	}
}
