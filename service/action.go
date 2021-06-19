package service

import (
	"github.com/go-playground/validator"
	"workflow/database/database"
	"workflow/models/DTO"
)

type Action struct {
	Validate        *validator.Validate
	DB              database.Database
	WorkflowService Workflow
	ActivityService Activity
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
		if err != nil {
			_ = db.Rollback()
		} else {
			err = db.Commit()
		}
	}()

	var actionID int
	actionID, err = db.GetActionRepository().Insert(actionDO)

	return actionID, nil
}

func (a Action) Update(cmd DTO.UpdateActionCommand) error {
	err := a.Validate.Struct(cmd)
	if err != nil {
		return err
	}

	db, err := a.DB.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = db.Rollback()
		} else {
			err = db.Commit()
		}
	}()

	err = db.GetActionRepository().Update() // TODO: add parameters

	return err
}

func (a Action) SetActionStatus(cmd DTO.SetActionStatusCommand) error {
	err := a.Validate.Struct(cmd)
	if err != nil {
		return err
	}

	db, err := a.DB.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = db.Rollback()
		} else {
			err = db.Commit()
		}
	}()

	err = a.DB.GetActionRepository().SetStatus(cmd.ActionID, cmd.Enabled)

	return err
}

func (a Action) ListActionsByTarget(query DTO.ListActionsByTargetQuery) ([]DTO.Action, error) {
	err := a.Validate.Struct(query)
	if err != nil {
		return nil, err
	}

	actionDOs, err := a.DB.GetActionRepository().ListByTarget(query.Target)
	if err != nil {
		return nil, err
	}

	result := make([]DTO.Action, 0, len(actionDOs))
	for _, actionDO := range actionDOs {
		actionDTO := actionDOToActionDTO(*actionDO)
		result = append(result, actionDTO)
	}

	return result, nil
}

func (a Action) GetByID(query DTO.GetActionByIDQuery) (DTO.Action, error) {
	err := a.Validate.Struct(query)
	if err != nil {
		return DTO.Action{}, err
	}

	actionDO, err := a.DB.GetActionRepository().FindByID(query.ActionID)
	if err != nil {
		return DTO.Action{}, err
	}

	return actionDOToActionDTO(*actionDO), nil
}

func (a Action) Launch(cmd DTO.LaunchActionCommand) error {
	err := a.Validate.Struct(cmd)
	if err != nil {
		return err
	}

	actionDO, err := a.DB.GetActionRepository().FindByID(cmd.ActionID)
	if err != nil {
		return err
	}

	workflowID, err := a.WorkflowService.ExecuteAction(*actionDO)
	if err != nil {
		return err
	}

	return a.ActivityService.Insert(*actionDO, workflowID)
}

func (a Action) DeleteByID(cmd DTO.DeleteActionByIDCommand) error {
	err := a.Validate.Struct(cmd)
	if err != nil {
		return err
	}

	db, err := a.DB.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = db.Rollback()
		} else {
			err = db.Commit()
		}
	}()

	err = a.DB.GetActionRepository().Delete(cmd.ActionID)

	return err
}
