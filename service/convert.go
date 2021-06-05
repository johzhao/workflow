package service

import (
	"workflow/models/DO"
	"workflow/models/DTO"
)

func createActionCmdToActionDO(cmd DTO.CreateActionCommand) DO.Action {
	return DO.Action{
		Title:   cmd.Title,
		Content: cmd.Content,
		Trigger: DO.Trigger{
			Target:  cmd.Target,
			Code:    cmd.EventCode,
			Enabled: true,
		},
		Parameters:   cmd.Parameters,
		Environments: cmd.Environments,
	}
}

func actionDOToActionDTO(actionDO DO.Action) DTO.Action {
	return DTO.Action{
		ID:           actionDO.ID,
		Title:        actionDO.Title,
		Content:      actionDO.Content,
		Target:       actionDO.Trigger.Target,
		EventCode:    actionDO.Trigger.Code,
		Enabled:      actionDO.Trigger.Enabled,
		Parameters:   actionDO.Parameters,
		Environments: actionDO.Environments,
	}
}
