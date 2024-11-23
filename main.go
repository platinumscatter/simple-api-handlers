package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var task Message

func TaskRecorder(w http.ResponseWriter, r *http.Request) {
	var body Message
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid request body: %v", err)
		return
	}

	task = body

	result := DB.Create(&task)
	if result.Error != nil {
		panic(result)
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Task recorded successfully")
}

func GreetTask(w http.ResponseWriter, r *http.Request) {
	DB.Find(&task)
	fmt.Fprintf(w, "Fields: %v", task)
}

func main() {
	InitDB()

	DB.AutoMigrate(&Message{})

	router := mux.NewRouter()
	router.HandleFunc("/api/task", GreetTask).Methods("GET")
	router.HandleFunc("/api/task", TaskRecorder).Methods("POST")

	http.ListenAndServe(":8080", router)
}
