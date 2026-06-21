package main

import (
	"log"
	"net/http"
	"tasks-api/internal/handlers"
	"tasks-api/internal/storage/memory"
)

func main() {
	storage := memory.New()
	handler := handlers.NewTaskHandler(storage)
	mux := http.NewServeMux()
	mux.HandleFunc("POST /tasks", handler.CreateTask)
	mux.HandleFunc("GET /tasks", handler.ListTasks)
	mux.HandleFunc("GET /tasks/{id}", handler.GetTask)
	mux.HandleFunc("PUT /tasks/{id}", handler.UpdateTask)
	mux.HandleFunc("DELETE /tasks/{id}", handler.DeleteTask)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
