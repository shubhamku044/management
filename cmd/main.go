package main

import (
	"fmt"

	"github.com/shubhamku044/management/api"
	"github.com/shubhamku044/management/controller"
)

func main() {
	api := api.ApiRoutes{}
	api.StartApp(controller.Server{})
	fmt.Println("Server is running")
	fmt.Printf("Main Server: %v\n", api)
}
