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
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid request body: %v", err)
		return
	}

	task = bodyHolder

	result := DB.Create(&task)
	if result.Error != nil {
		panic(result)
	}
	w.WriteHeader(http.StatusOK)
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

func main() {
	InitDB()

	DB.AutoMigrate(&Message{})

	router := mux.NewRouter()
	router.HandleFunc("/api/tasks", GetAllTasks).Methods("GET")
	router.HandleFunc("/api/task", CreateTask).Methods("POST")

	http.ListenAndServe(":8080", router)
}
