package adapter

import (
	"database/sql"
	"githab/mrflick72/go-playground/src/model"
)

type MySqlTodoRepository struct {
	connectionString string
}

func (repository *MySqlTodoRepository) GetAllTodo() ([]*model.Todo, error) {
	sql.Open("mysql", repository.connectionString)

	return nil, nil
}

func (repository *MySqlTodoRepository) GetTodo(id int64) (*model.Todo, error) {
	return nil, nil
}

func (repository *MySqlTodoRepository) SaveTodo(todo *model.Todo) error {
	return nil
}

func (repository *MySqlTodoRepository) RemoveTodo(id int64) error {
	return nil
}
