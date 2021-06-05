package repository

import "workflow/models/DO"

type Action interface {
	Insert(action DO.Action) (int, error)
	FindByID(actionID int) (*DO.Action, error)
}
