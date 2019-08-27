package services

import "errors"

var (
	ErrInvalidTitle = errors.New("Title field cannot be empty or null")

	ErrInvalidBody = errors.New("Body field cannot be empty or null")

	ErrNoID = errors.New("ID field cannot be empty or null")
)

// Actions implements methods for getting todos
type Actions interface {
	AllTodos() ([]*Todo, error)
	GetTodoByID(id string) (*Todo, error)
	AddTodo(todo *Todo) error
	UpdateTodo(todo *Todo) error
	DeleteTodo(id string) error
}

//Todo defines what a todo entry contains
type Todo struct {
	ID    string `json:"ID"`
	Title string `json:"Title"`
	Body  string `json:"Body"`
}

type Validation interface {
	Validate() error
}

func (t Todo) Validate() error {
	if t.Title == "" {
		return ErrInvalidTitle
	}
	if t.Body == "" {
		return ErrInvalidBody
	}
	return nil
}
