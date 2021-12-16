package src

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gitlab.com/url-builder/go-admin/src/config"
	"gitlab.com/url-builder/go-admin/src/controllers"
	"gitlab.com/url-builder/go-admin/src/database/connection"
	"gitlab.com/url-builder/go-admin/src/database/repositories"
)

func BootStrap(envPath string) *gin.Engine {
	// Load configs
	configs := config.LoadConfigs(envPath)

	r := gin.New()
	db, err := connection.Connect(configs.DB)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Database successfully connected")
	}

	controllers.Register(r, repositories.GetRepositoriesPool(db), configs)

	return r
}
