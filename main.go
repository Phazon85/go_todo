package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/phazon85/go_todo/handler"
	"github.com/phazon85/go_todo/services"

	_ "github.com/lib/pq"
)

const (
	configFile = "dev.yaml"
)

func main() {
	database := services.DBInit(configFile)
	handler := handler.NewTodoHandler(database)

	r := mux.NewRouter()
	r.HandleFunc("/todo", handler.HandleGetTodos).Methods("GET")
	r.HandleFunc("/todo/{id:[0-9]+}", handler.HandleGetTodoByID).Methods("GET")
	r.HandleFunc("/todo", handler.HandleAddTodo).Methods("POST")
	r.HandleFunc("/todo", handler.HandleUpdateTodo).Methods("PUT")
	r.HandleFunc("/todo", handler.HandleDeleteTodo).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r))
}
