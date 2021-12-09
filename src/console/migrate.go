package main

import (
	"fmt"

	"gitlab.com/url-builder/go-admin/src/config"
	"gitlab.com/url-builder/go-admin/src/database/connection"
)

func main() {
	fmt.Println("Migration script start.")
	config.Setup()
	_, err := connection.Connect()
	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println("Database successfully connected")
	}

	err = connection.Migrate()
	if err != nil {
		fmt.Println("Migration error")
		return
	}
	fmt.Println("Migration script end.")
}
