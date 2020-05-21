package main

import (
	"fmt"
	"githab/mrflick72/go-playground/src/adapter"
	"githab/mrflick72/go-playground/src/model"
	"githab/mrflick72/go-playground/src/web"
	"github.com/labstack/echo"
)

func main() {
	inMemoryRepository := adapter.InMemoryTodoRepository{}
	repository := adapter.MySqlTodoRepository{ConnectionString: "root:root@tcp(localhost)/todo"}

	server := echo.New()
	initDatabase(&inMemoryRepository)

	web.Endpoints(server, &repository)

	server.Start(":8000")
}

func initDatabase(repository *adapter.InMemoryTodoRepository) {
	fmt.Println("save a todo")

	repository.SaveTodo(&model.Todo{
		Id:      1,
		Content: "a content",
	})
	todoList, _ := repository.GetAllTodo()
	for _, value := range todoList {
		fmt.Println(*value)
	}
}
