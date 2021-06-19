package service

import (
	"workflow/errors"
	"workflow/models/DO"
)

type Workflow struct {
}

func (w Workflow) ExecuteAction(action DO.Action) (string, error) {
	return "", errors.ErrNeedImplement.New("")
}
