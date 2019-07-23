package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type todo struct {
	Message string
	ID      int
}

var todos = []*todo{
	&todo{
		Message: "Hello",
		ID:      0,
	},
}

func getTodo(w http.ResponseWriter, r *http.Request) {
	log.Printf("Incoming GET request on: %s", r.URL.Path)
	json.NewEncoder(w).Encode(todos)
}

func postTodo(w http.ResponseWriter, r *http.Request) {
	log.Printf("Incoming POST request on: %s", r.URL.Path)
	newTodo := &todo{}
	err := json.NewDecoder(r.Body).Decode(newTodo)
	if err != nil {
		log.Printf("Error decoding post Todo: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
	}
	newTodo.ID = len(todos)
	todos = append(todos, newTodo)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todos)
}

func delTodo(w http.ResponseWriter, r *http.Request) {
	log.Printf("Incoming DELETE request on: %s", r.URL.Path)
	for i, v := range todos {
		if r.Header.Get("ID") == strconv.Itoa(v.ID) {
			todos = todos[:i+copy(todos[i:], todos[i+1:])]
		}
	}
	w.WriteHeader(http.StatusOK)
}

func putTodo(w http.ResponseWriter, r *http.Request) {
	log.Printf("Incoming PUT request on: %s", r.URL.Path)
	newTodo := &todo{}
	err := json.NewDecoder(r.Body).Decode(newTodo)
	if err != nil {
		log.Printf("Error decoding put Todo: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
	}
	for i, v := range todos {
		if r.Header.Get("ID") == strconv.Itoa(v.ID) {
			todos[i].Message = newTodo.Message
		}
	}
	w.WriteHeader(http.StatusOK)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getTodo(w, r)
	case "POST":
		postTodo(w, r)
	case "DELETE":
		delTodo(w, r)
	case "PUT":
		putTodo(w, r)
	default:
		fmt.Fprintf(w, "Unknown method")
	}
}

func main() {
	http.HandleFunc("/todo", rootHandler)
	http.ListenAndServe(":8080", nil)
}
