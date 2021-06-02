package sqlx

import (
	"encoding/json"
	_interface "workflow/database/interface"
)

func actionDOToActionPO(action _interface.ActionDO) (actionPO, error) {
	context := actionContext{
		Parameters:   action.Parameters,
		Environments: action.Environments,
	}
	if context.Parameters == nil {
		context.Parameters = make([]_interface.ActionParameter, 0)
	}
	if context.Environments == nil {
		context.Environments = make([]_interface.ActionEnvironment, 0)
	}

	contextData, err := json.Marshal(context)
	if err != nil {
		return actionPO{}, err
	}

	return actionPO{
		Title:   action.Title,
		Content: action.Content,
		Context: string(contextData),
	}, nil
}

func actionDOToTriggerPO(action _interface.ActionDO) triggerPO {
	return triggerPO{
		Target:   action.Trigger.Target,
		Code:     action.Trigger.Code,
		Enabled:  action.Trigger.Enabled,
		HookType: hookTypeAction,
	}
}

func makeActionDO(action actionPO, trigger triggerPO) (_interface.ActionDO, error) {
	var context actionContext
	if err := json.Unmarshal([]byte(action.Context), &context); err != nil {
		return _interface.ActionDO{}, err
	}

	return _interface.ActionDO{
		ID:      action.ID,
		Title:   action.Title,
		Content: action.Content,
		Trigger: _interface.Trigger{
			Target:  trigger.Target,
			Code:    trigger.Code,
			Enabled: trigger.Enabled,
		},
		Parameters:   context.Parameters,
		Environments: context.Environments,
	}, nil
}
