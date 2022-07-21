package main

import (
	"fmt"
	"todo/database/migration"
	"todo/server"
)

func main() {
	err := migration.ConnectAndMigrate("localhost", "5433", "tod", "local", "local", migration.SSLModeDisable)
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected")
	svr := server.SetUpRoutes()
	err = svr.Run(":8080")
	if err != nil {
		panic(err)
	}
}
