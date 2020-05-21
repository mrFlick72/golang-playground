package model

type Todo struct {
	Id      int64
	Content string
}

type TodoRepository interface {
	GetAllTodo() ([]*Todo, error)
	GetTodo(id int64) (*Todo, error)
	SaveTodo(todo *Todo) error
	RemoveTodo(id int64) error
}
