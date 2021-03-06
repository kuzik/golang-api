package main

import (
	"fmt"

	"gitlab.com/url-builder/go-admin/src/config"
	"gitlab.com/url-builder/go-admin/src/database/connection"
)

func main() {
	fmt.Println("Migration script start.")
	configs := config.LoadConfigs(".env")
	dbConnection, err := connection.Connect(configs.DB)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Database successfully connected")
	}

	err = connection.Migrate(dbConnection)
	if err != nil {
		fmt.Println("Migration error")
		return
	}
	fmt.Println("Migration script end.")
}
