package main

import (
	"fmt"
	"webservermod/api"
	"webservermod/dao"
)

func main() {
	addr := ":8080"
	dbConnStr := "postgres://postgres:postgres@192.168.0.107:5432/users"

	apiServer := api.NewServer(addr, dao.NewUserDB(dbConnStr))
	if apiServer != nil {
		fmt.Println("DB connection successful ")
	}

	if err := apiServer.Start(); err != nil {
		fmt.Printf("error while starting the server...%v", err)
		return
	}

}
