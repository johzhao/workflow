package sqlx

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_interface "workflow/database/interface"
)

func NewSQLXDatabase(driverName string, dataSourceName string) (_interface.Database, error) {
	db, err := sqlx.Connect(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}

	return &database{
		db: db,
	}, nil
}

type database struct {
	db *sqlx.DB
	tx *sqlx.Tx
}

func (d *database) Close() error {
	if d.tx != nil {
		return fmt.Errorf("in transaction")
	}
	return d.db.Close()
}

func (d *database) Begin() (_interface.Database, error) {
	if d.tx != nil {
		return d, nil
	}

	tx, err := d.db.Beginx()
	if err != nil {
		return nil, err
	}

	return &database{
		db: d.db,
		tx: tx,
	}, nil
}

func (d *database) Rollback() error {
	if d.tx == nil {
		return fmt.Errorf("not in transaction")
	}

	if err := d.tx.Rollback(); err != nil {
		return err
	}

	d.tx = nil
	return nil
}

func (d *database) Commit() error {
	if d.tx == nil {
		return fmt.Errorf("not in transaction")
	}

	if err := d.tx.Commit(); err != nil {
		return err
	}

	d.tx = nil
	return nil
}

func (d database) GetActionRepository() _interface.ActionRepository {
	return newActionRepository(d.getExt())
}

func (d database) getExt() sqlx.Ext {
	if d.tx != nil {
		return d.tx
	}
	return d.db
}
