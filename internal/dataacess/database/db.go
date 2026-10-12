
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

func InitializeDB(
    ctx context.Context,
    cfg configs.Database.Database,
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

    if err := db.PingContext(ctx); err != nil {
        cleanup()
        return nil, nil, err
    }

    return db, cleanup, nil
}

func InitializeGoquDB(db *sql.DB) *goqu.Database {
    return goqu.New("mysql", db)
}