package api

import (
	"github.com/shubhamku044/management/controller"
	"github.com/shubhamku044/management/store"
)

type ApiRoutes struct {
	Server controller.ServerOperations
}

func (api *ApiRoutes) StartApp(server controller.Server) {
	api.Server = &server
	api.Server.NewServer(store.Postgres{})
}
