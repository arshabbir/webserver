package main

import (
	"fmt"
	"webservermod/api"
)

func main() {
	addr := ":8080"
	apiServer := api.NewServer(addr)

	if err := apiServer.Start(); err != nil {
		fmt.Printf("error while starting the server...%v", err)
		return
	}

}
