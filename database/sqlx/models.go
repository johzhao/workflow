package sqlx

import _interface "workflow/database/interface"

const (
	hookTypeAction = "action"
)

type triggerPO struct {
	ID       int    `db:"id"`
	Target   string `db:"target"`
	Code     int    `db:"code"`
	Enabled  bool   `db:"enabled"`
	HookType string `db:"hook_type"`
	HookID   int    `db:"hook_id"`
}

type actionContext struct {
	Parameters   []_interface.ActionParameter   `json:"parameters"`
	Environments []_interface.ActionEnvironment `json:"environments"`
}

type actionPO struct {
	ID      int    `db:"id"`
	Title   string `db:"title"`
	Content string `db:"content"`
	Context string `db:"context"`
}
