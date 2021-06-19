package repository

import "workflow/models/DO"

type Action interface {
	Insert(action DO.Action) (int, error)

	Delete(actionID int) error

	Update() error
	SetStatus(actionID int, status bool) error

	ListByTarget(target string) ([]*DO.Action, error)
	FindByID(actionID int) (*DO.Action, error)
}
