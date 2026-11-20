
package database

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"

	"go_load/internal/configs"
)

type Database interface {
    Delete(table interface{}) *goqu.DeleteDataset
    Dialect() string
    Exec(query string, args ...interface{}) (sql.Result, error)
    ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
    From(from ...interface{}) *goqu.SelectDataset
    Insert(table interface{}) *goqu.InsertDataset
    Logger(logger goqu.Logger)
    Prepare(query string) (*sql.Stmt, error)
    PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
    Query(query string, args ...interface{}) (*sql.Rows, error)
    QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
    QueryRow(query string, args ...interface{}) *sql.Row
    QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
    ScanStruct(i interface{}, query string, args ...interface{}) (bool, error)
    ScanStructContext(ctx context.Context, i interface{}, query string, args ...interface{}) (bool, error)
    ScanStructs(i interface{}, query string, args ...interface{}) error
    ScanStructsContext(ctx context.Context, i interface{}, query string, args ...interface{}) error
    ScanVal(i interface{}, query string, args ...interface{}) (bool, error)
    ScanValContext(ctx context.Context, i interface{}, query string, args ...interface{}) (bool, error)
    ScanVals(i interface{}, query string, args ...interface{}) error
    ScanValsContext(ctx context.Context, i interface{}, query string, args ...interface{}) error
    Select(cols ...interface{}) *goqu.SelectDataset
    Trace(op string, sqlString string, args ...interface{})
    Truncate(table ...interface{}) *goqu.TruncateDataset
    Update(table interface{}) *goqu.UpdateDataset
    

}


func InitializeDB(
    cfg configs.Database,
) (*sql.DB, func(), error) {
    dsn := fmt.Sprintf(
        "%s:%s@tcp(%s:%s)/%s?parseTime=true",
        cfg.Username,
        cfg.Password,
        cfg.Host,
        cfg.Port,
        cfg.Database,
    )

    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, nil, err
    }

    cleanup := func() {
        _ = db.Close()
    }

    if err := db.Ping(); err != nil {
        cleanup()
        return nil, nil, err
    }

    return db, cleanup, nil
}

func InitializeGoquDB(db *sql.DB) *goqu.Database {
    return goqu.New("mysql", db)
}