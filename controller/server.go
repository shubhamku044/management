package controller

import "github.com/shubhamku044/management/store"

type Server struct {
	PostgresDB *store.Postges
}

func (s *Server) NewServer() {
	server := Server{}
	server.PostgresDB.NewStore()
}
