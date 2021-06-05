package sqlx

import (
	"github.com/jmoiron/sqlx"
	"workflow/database/database/repository"
	"workflow/models/DO"
)

func newActionRepository(ext sqlx.Ext) repository.Action {
	return &actionRepository{
		actionDAO:  newActionDAO(ext),
		triggerDAO: newTriggerDAO(ext),
	}
}

type actionRepository struct {
	actionDAO  actionDAO
	triggerDAO triggerDAO
}

func (a actionRepository) Insert(action DO.Action) (int, error) {
	actionPO, err := actionDOToActionPO(action)
	if err != nil {
		return 0, err
	}

	actionID, err := a.actionDAO.Insert(actionPO)
	if err != nil {
		return 0, err
	}

	triggerPO := actionDOToTriggerPO(action)
	triggerPO.HookID = actionID
	if _, err := a.triggerDAO.Insert(triggerPO); err != nil {
		return 0, err
	}

	return actionID, nil
}

func (a actionRepository) FindByID(actionID int) (*DO.Action, error) {
	actionPO, err := a.actionDAO.FindByID(actionID)
	if err != nil {
		return nil, err
	}

	triggerPO, err := a.triggerDAO.FindByActionID(actionID)
	if err != nil {
		return nil, err
	}

	actionDO, err := makeActionDO(*actionPO, *triggerPO)
	if err != nil {
		return nil, err
	}

	return &actionDO, nil
}
