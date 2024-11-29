package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/platinumscatter/simple_api/internal/database"
	"github.com/platinumscatter/simple_api/internal/handlers"
	"github.com/platinumscatter/simple_api/internal/taskService"
	"github.com/platinumscatter/simple_api/internal/web/tasks"
)

func main() {
	database.InitDB()
	if err := database.DB.AutoMigrate(&taskService.Task{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewService(*repo)

	handler := handlers.NewHandler(service)

	e := echo.New()

	e.Use(middleware.Logger()) 
	e.Use(middleware.Recover())

	strictHandler := tasks.NewStrictHandler(handler, nil)
	tasks.RegisterHandlers(e, strictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
