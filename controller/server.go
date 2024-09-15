package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/shubhamku044/management/model"
	"github.com/shubhamku044/management/store"
	"github.com/shubhamku044/management/util"
)

type Server struct {
	PostgresDB store.StoreOperations
}

func (s *Server) NewServer(pgstore store.Postgres) {
	util.SetLogger()
	util.Logger.Infof("Logger setup is done\n")
	s.PostgresDB = &pgstore
	err := s.PostgresDB.NewStore()
	if err != nil {
		util.Log(model.LogLevelError, model.ControllerPackage, model.NewServer, "error in creating new store", err)
	} else {
		util.Log(model.LogLevelInfo, model.ControllerPackage, model.NewServer, "new store created", nil)
	}
}

type ServerOperations interface {
	NewServer(pgstore store.Postgres)
	CreateUser(ctx *gin.Context)
}
