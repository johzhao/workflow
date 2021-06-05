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
}

func (a Action) CreateAction(ctx *gin.Context) (resp interface{}, err error) {
	cmd := DTO.CreateActionCommand{} // TODO: fill with request data
	return a.ActionService.Create(cmd)
}
