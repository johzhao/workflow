package main

import (
	"fmt"
	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"os"
	"workflow/controller"
	"workflow/database"
	"workflow/logger"
	"workflow/server"
	"workflow/service"
)

func main() {
	myEnvs, err := godotenv.Read()
	if err != nil {
		fmt.Printf("failed to load .env file")
		os.Exit(1)
	}

	zapLogger := logger.SetupLogger()

	dbDriverName := myEnvs["DB_DRIVER_NAME"]
	dbDataSourceName := myEnvs["DB_DATA_SOURCE_NAME"]
	db, err := database.NewDatabase(dbDriverName, dbDataSourceName)
	if err != nil {
		zapLogger.Error("failed to create database object",
			zap.Error(err))
		os.Exit(1)
	}
	defer func() {
		_ = db.Close()
	}()

	validate := validator.New()

	actionService := service.Action{
		Validate: validate,
		DB:       db,
	}

	actionController := controller.Action{
		ActionService: actionService,
	}

	webServer := server.NewWebServer(zapLogger, actionController)
	if err := webServer.SetupServer(); err != nil {
		zapLogger.Info("setup server failed", zap.Error(err))
		os.Exit(1)
	}

	zapLogger.Info("start server")

	g := server.MakeGroup()
	g.Add(webServer.RunServer, webServer.StopServer)

	if err := g.Run(); err != nil {
		zapLogger.Info("run failed", zap.Error(err))
		os.Exit(1)
	}
}
