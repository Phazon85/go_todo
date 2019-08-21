package services

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
