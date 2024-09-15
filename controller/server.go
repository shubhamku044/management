package controller

import (
	"fmt"

	"github.com/shubhamku044/management/store"
)

type Server struct {
	PostgresDB store.StoreOperations
}

func (s *Server) NewServer(pgstore store.Postgres) {
	s.PostgresDB = &pgstore
	s.PostgresDB.NewStore()
	fmt.Printf("Server: %v\n", s)
}

type ServerOperations interface {
	NewServer(pgstore store.Postgres)
}
