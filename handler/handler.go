package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/phazon85/go_todo/services"
)

//TodoHandler implements the Actions Interface
type TodoHandler struct {
	Service services.Actions
}

func decodeAndValidate(r *http.Request, v services.Validation) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}
	defer r.Body.Close()

	return v.Validate()
}

func encodeJSON(w http.ResponseWriter, v interface{}) {
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Printf("Error encoding JSON")
	}
}

// NewTodoHandler returns a new TodoHandler service
func NewTodoHandler(svc services.Actions) *TodoHandler {
	return &TodoHandler{
		Service: svc,
	}
}

//HandleGetTodos gets all todos in DB
func (t *TodoHandler) HandleGetTodos(w http.ResponseWriter, r *http.Request) {
	res, err := t.Service.AllTodos()
	if err != nil {
		log.Printf("Error getting todos: %s", err.Error())
	}

	encodeJSON(w, res)
}

//HandleGetTodoByID will queries a Todo by ID
func (t *TodoHandler) HandleGetTodoByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	res, err := t.Service.GetTodoByID(vars["id"])
	if err != nil {
		log.Printf("error handling getting Todo by ID: %s", err.Error())
	}

	encodeJSON(w, res)
}

//HandleAddTodo takes a new todo and inserts into DB
func (t *TodoHandler) HandleAddTodo(w http.ResponseWriter, r *http.Request) {
	newTodo := &services.Todo{}
	err := decodeAndValidate(r, newTodo)
	if err != nil {
		log.Printf("Error decoding post Todo: %s", err.Error())
	}
	err = t.Service.AddTodo(newTodo)
	if err != nil {
		log.Printf("Error adding todo to DB: %s", err.Error())
	}

	w.WriteHeader(http.StatusCreated)
}

//HandleUpdateTodo replaces values in todo by ID
func (t *TodoHandler) HandleUpdateTodo(w http.ResponseWriter, r *http.Request) {
	newTodo := &services.Todo{}
	err := decodeAndValidate(r, newTodo)
	if err != nil {
		log.Printf("Error decoding post Todo: %s", err.Error())
	}
	err = t.Service.UpdateTodo(newTodo)
	if err != nil {
		log.Printf("Error Updating Todo: %s", err.Error())
	}

	w.WriteHeader(http.StatusOK)
}

//HandleDeleteTodo will delete a todo by ID
func (t *TodoHandler) HandleDeleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err := t.Service.DeleteTodo(vars["id"])
	if err != nil {
		log.Printf("Error deleteing Todo: %s", err.Error())
	}

	w.WriteHeader(http.StatusOK)
}
