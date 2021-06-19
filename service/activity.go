package service

import (
	"workflow/database/database"
	"workflow/errors"
	"workflow/models/DO"
)

type Activity struct {
	DB database.Database
}

func (a Activity) Insert(action DO.Action, workflowID string) error {
	return errors.ErrNeedImplement.New("")
}
