package DTO

import "workflow/models"

type CreateActionCommand struct {
	Title        string                     `json:"title" validate:"required,alphanum"`
	Content      string                     `json:"content" validate:"required"`
	Target       string                     `json:"target"`
	EventCode    int                        `json:"event_code" validate:"gt=0"`
	Parameters   []models.ActionParameter   `json:"parameters"`
	Environments []models.ActionEnvironment `json:"environments"`
}

type UpdateActionCommand struct {
	ActionID int `validate:"gt=0"`
	// TODO:
}

type SetActionStatusCommand struct {
	ActionID int `validate:"gt=0"`
	Enabled  bool
}

type LaunchActionCommand struct {
	ActionID int `validate:"gt=0"`
	// TODO:
}

type DeleteActionByIDCommand struct {
	ActionID int `validate:"gt=0"`
	// TODO:
}
