package api

import (
	"github.com/gin-gonic/gin"
	"github.com/shubhamku044/management/controller"
	_ "github.com/shubhamku044/management/docs"
	"github.com/shubhamku044/management/store"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

type ApiRoutes struct {
	Server controller.ServerOperations
}

func (api *ApiRoutes) StartApp(router *gin.Engine, server controller.Server) {
	api.Server = &server
	api.Server.NewServer(store.Postgres{})

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api.UserRoutes(router)
}
