package sqlx

import (
	"github.com/jmoiron/sqlx"
	"workflow/models/PO"
)

func newActionDAO(ext sqlx.Ext) actionDAO {
	return actionDAO{
		ext: ext,
	}
}

type actionDAO struct {
	ext sqlx.Ext
}

const (
	actionTableName = "actions"
	actionColumns   = "id, title, content, context"
)

func (a actionDAO) Insert(action PO.Action) (int, error) {
	sql := "INSERT INTO " + actionTableName + " (title, content, context) VALUES (?, ?, ?)"
	result, err := a.ext.Exec(sql, action.Title, action.Content, action.Context)
	if err != nil {
		return 0, err
	}

	actionID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(actionID), nil
}

func (a actionDAO) FindByID(actionID int) (*PO.Action, error) {
	sql := "SELECT " + actionColumns + " FROM " + actionTableName + " WHERE id = ?"
	row := a.ext.QueryRowx(sql, actionID)
	var result PO.Action
	if err := row.StructScan(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
