package service

import (
	"github.com/go-playground/validator"
	"workflow/database/database"
	"workflow/models/DTO"
)

type Action struct {
	Validate *validator.Validate
	DB       database.Database
}

func (a Action) Create(cmd DTO.CreateActionCommand) (int, error) {
	err := a.Validate.Struct(cmd)
	if err != nil {
		return 0, err
	}

	actionDO := createActionCmdToActionDO(cmd)

	db, err := a.DB.Begin()
	if err != nil {
		return 0, err
	}
	defer func() {
		_ = db.Rollback()
	}()

	actionID, err := db.GetActionRepository().Insert(actionDO)
	if err != nil {
		return 0, err
	}

	if err := db.Rollback(); err != nil {
		return 0, err
	}

	return actionID, nil
}

func (a Action) GetByID(actionID int) (DTO.Action, error) {
	actionDO, err := a.DB.GetActionRepository().FindByID(actionID)
	if err != nil {
		return DTO.Action{}, err
	}
	return actionDOToActionDTO(*actionDO), nil
}
