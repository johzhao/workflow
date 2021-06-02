package database

import (
	_interface "workflow/database/interface"
	"workflow/database/sqlx"
)

func NewDatabase(driverName string, dataSourceName string) (_interface.Database, error) {
	return sqlx.NewSQLXDatabase(driverName, dataSourceName)
}
