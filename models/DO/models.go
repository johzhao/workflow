package DO

import (
	"workflow/models"
)

type Trigger struct {
	Target  string
	Code    int
	Enabled bool
}

type Action struct {
	ID           int
	Title        string
	Content      string
	Trigger      Trigger
	Parameters   []models.ActionParameter
	Environments []models.ActionEnvironment
}
