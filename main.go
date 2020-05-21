package main

import (
	"fmt"
	"githab/mrflick72/go-playground/src/model"
	"githab/mrflick72/go-playground/src/web"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	repository := model.InMemoryTodoRepository{}
	endpoint := web.TodoEndPoint{TodoRepository: &repository}

	initDatabase(repository)

	router := mux.NewRouter()
	router.Handle("/todos", endpoint)
	router.HandleFunc("/echo", web.HandleEcho)
	http.ListenAndServe(":3000", router)
}

func initDatabase(repository model.InMemoryTodoRepository) {
	fmt.Println("save a todo")

	repository.SaveTodo(model.Todo{
		Id:      "id",
		Content: "a content",
	})
	todoList, _ := repository.GetAllTodo()
	for _, value := range todoList {
		fmt.Println(*value)
	}
}
