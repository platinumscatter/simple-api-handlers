package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/platinumscatter/simple_api/internal/database"
	"github.com/platinumscatter/simple_api/internal/handlers"
	"github.com/platinumscatter/simple_api/internal/taskService"
	"github.com/platinumscatter/simple_api/internal/userService"
	"github.com/platinumscatter/simple_api/internal/web/tasks"
	"github.com/platinumscatter/simple_api/internal/web/users"
)

func main() {
	database.InitDB()
	if err := database.DB.AutoMigrate(&taskService.Task{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	database.InitDB()
	if err := database.DB.AutoMigrate(&userService.User{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	UserRepo := userService.NewUserRepository(database.DB)
	UserService := userService.NewService(*UserRepo)
	UserHandler := handlers.NewUserHandler(UserService)

	TaskRepo := taskService.NewTaskRepository(database.DB)
	TaskService := taskService.NewService(*TaskRepo)
	TaskHandler := handlers.NewHandler(TaskService)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictHandler := tasks.NewStrictHandler(TaskHandler, nil)
	tasks.RegisterHandlers(e, strictHandler)

	strictUserHandler := users.NewStrictHandler(UserHandler, nil)
	users.RegisterHandlers(e, strictUserHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
