package main

import (
	"github.com/shubhamku044/management/api"
	"github.com/shubhamku044/management/controller"
)

func main() {
	api := api.ApiRoutes{}
	api.StartApp(controller.Server{})
}
