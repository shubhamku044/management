package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shubhamku044/management/api"
	"github.com/shubhamku044/management/controller"
)

func main() {
	api := api.ApiRoutes{}
	controller := controller.Server{}
	routes := gin.Default()
	api.StartApp(routes, controller)

	routes.Run(":8080")
}
