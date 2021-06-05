package main

import (
	"github.com/go-playground/validator"
	"os"
	"workflow/database"
	"workflow/service"
)

func main() {
	dbDriverName := ""
	dbDataSourceName := ""
	db, err := database.NewDatabase(dbDriverName, dbDataSourceName)
	if err != nil {
		os.Exit(1)
	}
	defer func() {
		_ = db.Close()
	}()

	validate := validator.New()

	svc := service.Action{
		Validate: validate,
		DB:       db,
	}

	_, _ = svc.GetByID(0)
}
