package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/phazon85/go_todo/services"
)

type TodoHandler struct {
	Service services.Actions
}

func decodeAndValidate(r *http.Request, v services.Validation) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}
	defer r.Body.Close()

	return nil
}

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
