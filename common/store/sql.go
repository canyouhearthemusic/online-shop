package store

import (
	_ "database/sql"
	"errors"
	"strings"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jmoiron/sqlx"
)

// postgres://username:password@localhost:5432/dbname?sslmode=disable&search_path=public
type SQLX struct {
	Client *sqlx.DB
}

func NewSQL(dataSourceName string) (store SQLX, err error) {
	if !strings.Contains(dataSourceName, "://") {
		err = errors.New("store: undefined data source name " + dataSourceName)
		return
	}
	driverName := strings.ToLower(strings.Split(dataSourceName, "://")[0])

	store.Client, err = sqlx.Connect(driverName, dataSourceName)
	if err != nil {
		return
	}
	store.Client.SetMaxOpenConns(20)

	return
}
