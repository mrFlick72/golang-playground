package adapter

import (
	"errors"
	"githab/mrflick72/go-playground/src/model"
)

type InMemoryTodoRepository struct {
	database []*model.Todo
}

func (repository *InMemoryTodoRepository) GetAllTodo() ([]*model.Todo, error) {
	return repository.database, nil
}

func (repository *InMemoryTodoRepository) GetTodo(id int64) (*model.Todo, error) {
	database := repository.database

	for _, todo := range database {
		if todo.Id == id {
			return todo, nil
		}
	}
	return &model.Todo{}, errors.New("todo entry do not found")
}

func (repository *InMemoryTodoRepository) SaveTodo(todo *model.Todo) error {
	repository.database = append(repository.database, todo)
	return nil
}

func (repository *InMemoryTodoRepository) RemoveTodo(id int64) error {
	database := repository.database

	for index, todo := range database {
		if todo.Id == id {
			repository.database = append(database[:index], database[index+1:]...)
			return nil
		}
	}
	return nil
}
