package sqlx

import (
	"github.com/jmoiron/sqlx"
	"workflow/models/PO"
)

func newTriggerDAO(ext sqlx.Ext) triggerDAO {
	return triggerDAO{
		ext: ext,
	}
}

type triggerDAO struct {
	ext sqlx.Ext
}

const (
	triggerTableName = "triggers"
	triggerColumns   = "id, target, code, enabled, hook_type, hook_id"
)

func (t triggerDAO) Insert(trigger PO.Trigger) (int, error) {
	sql := "INSERT INTO " + triggerTableName + " (target, code, enabled, hook_type, hook_id) VALUES (?, ?, ?, ?, ?)"
	result, err := t.ext.Exec(sql, trigger.Target, trigger.Code, trigger.Enabled, trigger.HookType, trigger.HookID)
	if err != nil {
		return 0, err
	}

	triggerID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(triggerID), nil
}

func (t triggerDAO) FindByActionID(actionID int) (*PO.Trigger, error) {
	sql := "SELECT " + triggerColumns + " FROM " + triggerTableName + " WHERE hook_type = ? AND hook_id = ?"
	row := t.ext.QueryRowx(sql, PO.HookTypeAction, actionID)
	var result PO.Trigger
	if err := row.StructScan(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
