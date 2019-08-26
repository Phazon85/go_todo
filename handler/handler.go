package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/phazon85/go_todo/services"
)

type TodoHandler struct {
	Service services.Actions
}

// func decodeAndValidate(r *http.Request, v services.Validation) error {
// 	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
// 		return err
// 	}
// 	defer r.Body.Close()

// 	return Validate()
// }

func encodeJSON(w http.ResponseWriter, v interface{}) {
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Printf("Error encoding JSON")
	}
}

// NewTodoHandler returns a new
func NewTodoHandler(svc services.Actions) *TodoHandler {
	return &TodoHandler{
		Service: svc,
	}
}

func (t *TodoHandler) HandleGetTodos(w http.ResponseWriter, r *http.Request) {
	res, err := t.Service.AllTodos()
	if err != nil {
		log.Printf("Error getting todos: %s", err.Error())
	}

	encodeJSON(w, res)
}

func (t *TodoHandler) HandleGetTodoByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	res, err := t.Service.GetTodoByID(vars["id"])
	if err != nil {
		log.Printf("error handling getting Todo by ID: %s", err.Error())
	}

	encodeJSON(w, res)
}

func (t *TodoHandler) HandleAddTodo(w http.ResponseWriter, r *http.Request) {
	newTodo := &services.Todo{}
	err := json.NewDecoder(r.Body).Decode(newTodo)
	if err != nil {
		log.Printf("Error decoding post Todo: %s", err.Error())
	}
	err = t.Service.AddTodo(newTodo)
	if err != nil {
		log.Printf("Error adding todo to DB: %s", err.Error())
	}

	w.WriteHeader(http.StatusCreated)
}
