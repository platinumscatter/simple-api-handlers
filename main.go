package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var task Message

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var bodyHolder Message
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&bodyHolder)
	if err != nil {
		fmt.Fprintf(w, "Invalid request body: %v", err)
		return
	}

	task = bodyHolder

	result := DB.Create(&task)
	if result.Error != nil {
		panic(result)
	}
	fmt.Fprintf(w, "Task recorded successfully")
}

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []Message
	result := DB.Find(&tasks)
	if result.Error != nil {
		panic(result)
	}

	jsonData, err := json.Marshal(tasks)
	if err != nil {
		panic(err)
	}
	w.Write(jsonData)
}

func UpdateTaskById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var updatedTask Message
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&updatedTask)
	if err != nil {
		fmt.Fprintf(w, "Invalid request body: %v", err)
		return
	}

	jsonData, err := json.Marshal(updatedTask)
	if err != nil {
		panic(err)
	}

	result := DB.Model(&Message{}).Where("id = ?", id).Updates(updatedTask)
	if result.Error != nil {
		fmt.Fprintf(w, "Error updating task: %v", result.Error)
		return
	}

	fmt.Fprintf(w, "Task updated successfully:\n")
	w.Write(jsonData)
}

func DeleteTaskById(w http.ResponseWriter, r *http.Request) {
	var taskToDelete Message
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&taskToDelete)
	if err != nil {
		fmt.Fprintf(w, "Invalid request body: %v", err)
		return
	}

	result := DB.Delete(&Message{}, taskToDelete.ID)
	if result.Error != nil {
		fmt.Fprintf(w, "Error deleting task: %v", result.Error)
		return
	}

	fmt.Fprintf(w, "Task deleted successfully")
}

func main() {
	InitDB()

	DB.AutoMigrate(&Message{})

	router := mux.NewRouter()
	router.HandleFunc("/api/tasks", GetAllTasks).Methods("GET")
	router.HandleFunc("/api/tasks", CreateTask).Methods("POST")
	router.HandleFunc("/api/tasks/{id}", UpdateTaskById).Methods("PATCH")
	router.HandleFunc("/api/tasks/{id}", DeleteTaskById).Methods("DELETE")

	http.ListenAndServe(":8080", router)
}
