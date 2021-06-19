package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"workflow/models/DTO"
	"workflow/server"
	"workflow/service"
)

type Action struct {
	ActionService service.Action
}

func (a Action) SetupRoute(server *server.WebServer) {
	server.SetupRoute(http.MethodPost, "/actions", a.CreateAction)
	// TODO: Add more route
}

func (a Action) CreateAction(ctx *gin.Context) (interface{}, error) {
	cmd := DTO.CreateActionCommand{} // TODO: fill with request data
	return a.ActionService.Create(cmd)
}

func (a Action) UpdateAction(ctx *gin.Context) (interface{}, error) {
	cmd := DTO.UpdateActionCommand{} // TODO: fill with request data
	return nil, a.ActionService.Update(cmd)
}

func (a Action) SetActionStatus(ctx *gin.Context) (interface{}, error) {
	cmd := DTO.SetActionStatusCommand{} // TODO: fill with request data
	return nil, a.ActionService.SetActionStatus(cmd)
}

func (a Action) ListActionsByTarget(ctx *gin.Context) (interface{}, error) {
	query := DTO.ListActionsByTargetQuery{} // TODO: fill with request data
	return a.ActionService.ListActionsByTarget(query)
}

func (a Action) GetActionByID(ctx *gin.Context) (interface{}, error) {
	query := DTO.GetActionByIDQuery{} // TODO: fill with request data
	return a.ActionService.GetByID(query)
}

func (a Action) LaunchAction(ctx *gin.Context) (interface{}, error) {
	cmd := DTO.LaunchActionCommand{} // TODO: fill with request data
	return nil, a.ActionService.Launch(cmd)
}

func (a Action) DeleteActionByID(ctx *gin.Context) (interface{}, error) {
	cmd := DTO.DeleteActionByIDCommand{} // TODO: fill with request data
	return nil, a.ActionService.DeleteByID(cmd)
}
