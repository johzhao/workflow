package sqlx

import (
	"encoding/json"
	"workflow/models"
	"workflow/models/DO"
	"workflow/models/PO"
)

func actionDOToActionPO(action DO.Action) (PO.Action, error) {
	context := PO.ActionContext{
		Parameters:   action.Parameters,
		Environments: action.Environments,
	}
	if context.Parameters == nil {
		context.Parameters = make([]models.ActionParameter, 0)
	}
	if context.Environments == nil {
		context.Environments = make([]models.ActionEnvironment, 0)
	}

	contextData, err := json.Marshal(context)
	if err != nil {
		return PO.Action{}, err
	}

	return PO.Action{
		Title:   action.Title,
		Content: action.Content,
		Context: string(contextData),
	}, nil
}

func actionDOToTriggerPO(action DO.Action) PO.Trigger {
	return PO.Trigger{
		Target:   action.Trigger.Target,
		Code:     action.Trigger.Code,
		Enabled:  action.Trigger.Enabled,
		HookType: PO.HookTypeAction,
	}
}

func makeActionDO(action PO.Action, trigger PO.Trigger) (DO.Action, error) {
	var context PO.ActionContext
	if err := json.Unmarshal([]byte(action.Context), &context); err != nil {
		return DO.Action{}, err
	}

	return DO.Action{
		ID:      action.ID,
		Title:   action.Title,
		Content: action.Content,
		Trigger: DO.Trigger{
			Target:  trigger.Target,
			Code:    trigger.Code,
			Enabled: trigger.Enabled,
		},
		Parameters:   context.Parameters,
		Environments: context.Environments,
	}, nil
}
