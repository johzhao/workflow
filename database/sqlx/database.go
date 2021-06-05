package sqlx

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"workflow/database/database"
	"workflow/database/database/repository"
)

func NewSQLXDatabase(driverName string, dataSourceName string) (database.Database, error) {
	db, err := sqlx.Connect(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}

	return &sqlxDB{
		db: db,
	}, nil
}

type sqlxDB struct {
	db *sqlx.DB
	tx *sqlx.Tx
}

func (d *sqlxDB) Close() error {
	if d.tx != nil {
		return fmt.Errorf("in transaction")
	}
	return d.db.Close()
}

func (d *sqlxDB) Begin() (database.Database, error) {
	if d.tx != nil {
		return d, nil
	}

	tx, err := d.db.Beginx()
	if err != nil {
		return nil, err
	}

	return &sqlxDB{
		db: d.db,
		tx: tx,
	}, nil
}

func (d *sqlxDB) Rollback() error {
	if d.tx == nil {
		return fmt.Errorf("not in transaction")
	}

	if err := d.tx.Rollback(); err != nil {
		return err
	}

	d.tx = nil
	return nil
}

func (d *sqlxDB) Commit() error {
	if d.tx == nil {
		return fmt.Errorf("not in transaction")
	}

	if err := d.tx.Commit(); err != nil {
		return err
	}

	d.tx = nil
	return nil
}

func (d sqlxDB) GetActionRepository() repository.Action {
	return newActionRepository(d.getExt())
}

func (d sqlxDB) getExt() sqlx.Ext {
	if d.tx != nil {
		return d.tx
	}
	return d.db
}
