package service

import (
	"github.com/go-playground/validator"
	"workflow/database/database"
	"workflow/errors"
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

	if err := db.Commit(); err != nil {
		return 0, err
	}

	return actionID, nil
}

func (a Action) Update(cmd DTO.UpdateActionCommand) error {
	err := a.Validate.Struct(cmd)
	if err != nil {
		return err
	}

	return errors.ErrNeedImplement.New("")
}

func (a Action) SetActionStatus(cmd DTO.SetActionStatusCommand) error {
	err := a.Validate.Struct(cmd)
	if err != nil {
		return err
	}

	return errors.ErrNeedImplement.New("")
}

func (a Action) ListActionsByTarget(query DTO.ListActionsByTargetQuery) ([]DTO.Action, error) {
	err := a.Validate.Struct(query)
	if err != nil {
		return nil, err
	}

	return nil, errors.ErrNeedImplement.New("")
}

func (a Action) GetByID(query DTO.GetActionByIDQuery) (DTO.Action, error) {
	err := a.Validate.Struct(query)
	if err != nil {
		return DTO.Action{}, err
	}

	return a.getByID(query.ActionID)
}

func (a Action) Launch(cmd DTO.LaunchActionCommand) error {
	err := a.Validate.Struct(cmd)
	if err != nil {
		return err
	}

	action, err := a.getByID(cmd.ActionID)
	if err != nil {
		return err
	}

	// TODO: launch the action
	_ = action
	return errors.ErrNeedImplement.New("")
}

func (a Action) getByID(actionID int) (DTO.Action, error) {
	actionDO, err := a.DB.GetActionRepository().FindByID(actionID)
	if err != nil {
		return DTO.Action{}, err
	}
	return actionDOToActionDTO(*actionDO), nil
}
