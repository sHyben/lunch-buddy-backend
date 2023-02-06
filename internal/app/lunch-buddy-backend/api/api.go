package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sHyben/lunch-buddy-backend/internal/app/lunch-buddy-backend/api/router"
	"github.com/sHyben/lunch-buddy-backend/internal/pkg/private/config"
	"github.com/sHyben/lunch-buddy-backend/internal/pkg/private/db"
)

// setConfiguration sets up the configuration and the database
// It is called by Run
// It is not intended to be called by the user
// It panics if the configuration or the database could not be set up
func setConfiguration(configPath string) {
	config.Setup(configPath)
	db.SetupDB()
	gin.SetMode(config.GetConfig().Server.Mode)
}

// Run sets up the configuration and the database
// It starts the web server
func Run(configPath string) {
	if configPath == "" {
		configPath = "data/config.yml"
	}
	setConfiguration(configPath)
	conf := config.GetConfig()
	web := router.Setup()
	fmt.Println("Go API REST Running on port " + conf.Server.Port)
	fmt.Println("==================>")
	_ = web.Run(":" + conf.Server.Port)
}
