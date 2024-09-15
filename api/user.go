package api

import (
	"github.com/gin-gonic/gin"
	"github.com/shubhamku044/management/model"
	"github.com/shubhamku044/management/util"
)

func (api ApiRoutes) UserRoutes(routes *gin.Engine) {
	group := routes.Group("/user")

	{
		group.POST("/create", api.CreateUser)
	}
}

func (api ApiRoutes) CreateUser(ctx *gin.Context) {
	util.Log(model.LogLevelInfo, model.ApiPackage, model.CreateUser, "Creating new user", nil)
	api.Server.CreateUser(ctx)
}
