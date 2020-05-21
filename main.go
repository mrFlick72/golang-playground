package main

import (
	"fmt"
	"githab/mrflick72/go-playground/src/adapter"
	"githab/mrflick72/go-playground/src/model"
	"githab/mrflick72/go-playground/src/web"
	"github.com/labstack/echo"
)

func main() {
	repository := adapter.InMemoryTodoRepository{}

	server := echo.New()
	initDatabase(&repository)

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
