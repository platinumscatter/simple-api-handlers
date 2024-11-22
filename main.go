package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type requestBody struct {
	Message string `json:"message"`
}

var task requestBody

func TaskRecorder(w http.ResponseWriter, r *http.Request) {
	var body requestBody
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid request body: %v", err)
		return
	}
	task = body
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Task recorded successfully")

}

func GreetTask(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %v", task.Message)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/task", GreetTask).Methods("GET")
	router.HandleFunc("/api/task", TaskRecorder).Methods("POST")

	http.ListenAndServe(":8080", router)
}
