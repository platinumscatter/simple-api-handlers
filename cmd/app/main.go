package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/platinumscatter/simple_api/internal/database"
	"github.com/platinumscatter/simple_api/internal/handlers"
	"github.com/platinumscatter/simple_api/internal/taskService"
)

func main() {
	database.InitDB()
	database.DB.AutoMigrate(&taskService.Task{})

	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewService(*repo)

	handler := handlers.NewHandler(service)

	router := mux.NewRouter()
	router.HandleFunc("/api/tasks", handler.GetTaskHandler).Methods("GET")
	router.HandleFunc("/api/tasks", handler.PostTaskHandler).Methods("POST")
	router.HandleFunc("/api/tasks/{id}", handler.UpdateTaskHandler).Methods("PATCH")
	router.HandleFunc("/api/tasks/{id}", handler.DeleteTaskHandler).Methods("DELETE")
	http.ListenAndServe(":8080", router)
}
