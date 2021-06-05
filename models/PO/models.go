package PO

import (
	"workflow/models"
)

const (
	HookTypeAction = "action"
)

type Trigger struct {
	ID       int    `db:"id"`
	Target   string `db:"target"`
	Code     int    `db:"code"`
	Enabled  bool   `db:"enabled"`
	HookType string `db:"hook_type"`
	HookID   int    `db:"hook_id"`
}

type ActionContext struct {
	Parameters   []models.ActionParameter   `json:"parameters"`
	Environments []models.ActionEnvironment `json:"environments"`
}

type Action struct {
	ID      int    `db:"id"`
	Title   string `db:"title"`
	Content string `db:"content"`
	Context string `db:"context"`
}
