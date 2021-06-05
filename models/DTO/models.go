package DTO

import "workflow/models"

type Action struct {
	ID           int                        `json:"id"`
	Title        string                     `json:"title"`
	Content      string                     `json:"content"`
	Target       string                     `json:"target"`
	EventCode    int                        `json:"event_code"`
	Enabled      bool                       `json:"enabled"`
	Parameters   []models.ActionParameter   `json:"parameters"`
	Environments []models.ActionEnvironment `json:"environments"`
}
