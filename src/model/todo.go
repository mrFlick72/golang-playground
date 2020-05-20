package model

import "errors"

type Todo struct {
	Id      string
	Content string
}

type TodoRepository interface {
	GetAllTodo() ([]*Todo, error)
	GetTodo(id string) (*Todo, error)
	SaveTodo(todo Todo) error
	RemoveTodo(todo Todo) error
}

type InMemoryTodoRepository struct {
	database []*Todo
}

func (repository *InMemoryTodoRepository) GetAllTodo() ([]*Todo, error) {
	return repository.database, nil
}

func (repository *InMemoryTodoRepository) GetTodo(id string) (*Todo, error) {
	database := repository.database

	for _, todo := range database {
		if todo.Id == id {
			return todo, nil
		}
	}
	return &Todo{}, errors.New("todo entry do not found")
}

func (repository *InMemoryTodoRepository) SaveTodo(todo Todo) error {
	repository.database = append(repository.database, &todo)
	return nil
}

func (repository *InMemoryTodoRepository) RemoveTodo(todo Todo) error {

	return nil
}