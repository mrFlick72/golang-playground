package adapter

import (
	"database/sql"
	"fmt"
	"githab/mrflick72/go-playground/src/model"
	_ "github.com/go-sql-driver/mysql"
)

type MySqlTodoRepository struct {
	ConnectionString string
}

func (repository *MySqlTodoRepository) GetAllTodo() ([]*model.Todo, error) {
	var result []*model.Todo

	database, err := sql.Open("mysql", repository.ConnectionString)
	errorLog(err)

	query, _ := database.Prepare("SELECT id, content FROM TODO")
	rows, err := query.Query()
	errorLog(err)

	result = buildTodo(rows, result)

	closeResources(rows, query, database)

	return result, nil
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

func closeResources(rows *sql.Rows, query *sql.Stmt, database *sql.DB) {
	defer rows.Close()
	defer query.Close()
	defer database.Close()
}

func buildTodo(rows *sql.Rows, result []*model.Todo) []*model.Todo {
	for rows.Next() {
		var id int64
		var content string
		rows.Scan(&id, &content)
		todo := model.Todo{Id: id, Content: content}
		result = append(result, &todo)
	}
	return result
}

func errorLog(err error) {
	if err != nil {
		fmt.Println(err, "\n")
	}
}
