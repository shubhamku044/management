package main

import (
	"fmt"

	"github.com/shubhamku044/management/store"
)

type Server struct {
	PostgresDB *store.Postges
}

func main() {
	fmt.Println("Tring to connect to the database...")
	server := &Server{}
	server.PostgresDB = &store.Postges{}
	err := server.PostgresDB.NewStore()
	if err != nil {
		fmt.Println("Error while connecting to the database: ", err)
	}
	fmt.Println("Connected to the database successfully!")
}
