package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/salawatbro/chat-app/internal/config"
	"github.com/salawatbro/chat-app/internal/middlewares"
	"github.com/salawatbro/chat-app/internal/routes"
	"github.com/salawatbro/chat-app/pkg/database"
	"github.com/salawatbro/chat-app/pkg/utils"
	"go.uber.org/zap"
)

func main() {
	// load config
	configData, err := config.LoadConfig()
	if err != nil {
		utils.Logger.Panic("Error while loading config", zap.Error(err))
	}
	// initialize zap logger
	utils.ZapLogger(configData.App.Env)
	// initialize fiber app
	app := fiber.New()
	// connect to database
	db, err := database.Connect(&configData)
	if err != nil {
		utils.Logger.Panic("Error while connecting to database", zap.Error(err))
	}
	// disconnect from a database when app is closed
	defer func() {
		if err = db.Disconnect(); err != nil {
			utils.Logger.Error("Error while disconnecting from database", zap.Error(err))
		}
	}()
	//middleware setup
	middlewares.Setup(app, &configData)
	// setup routes
	routes.Setup(app, db, &configData)
	// listen to port 3000
	err = app.Listen(":" + configData.App.Port)
	if err != nil {
		utils.Logger.Panic("Error while listening to port", zap.Error(err))
	}
}
