package DTO

type ListActionsByTargetQuery struct {
	Target string `validate:"required"`
	// TODO:
}

type GetActionByIDQuery struct {
	ActionID int `validate:"gt=0"`
}
