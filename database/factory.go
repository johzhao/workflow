package database

import (
	"workflow/database/database"
	"workflow/database/sqlx"
)

func NewDatabase(driverName string, dataSourceName string) (database.Database, error) {
	return sqlx.NewSQLXDatabase(driverName, dataSourceName)
}
