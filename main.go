package main

import (
	"fmt"

	"gitlab.com/url-builder/go-admin/src/controllers"

	"github.com/gin-gonic/gin"

	"gitlab.com/url-builder/go-admin/src/config"
	"gitlab.com/url-builder/go-admin/src/database/connection"
	"gitlab.com/url-builder/go-admin/src/database/repositories"
)

func main() {
	// Load configs
	configs := config.LoadConfigs()

	r := gin.New()
	db, err := connection.Connect(configs.DB)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Database successfully connected")
	}

	controllers.Register(r, repositories.GetRepositoriesPool(db), configs)

	if err = r.Run(); err != nil {
		fmt.Println("Exception occurred: " + err.Error())
	}
}
