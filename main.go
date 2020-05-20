package main

import (
	"fmt"
	"githab/mrflick72/go-playground/src/model"
	"githab/mrflick72/go-playground/src/web"
	"net/http"
)

func main() {
	repository := model.InMemoryTodoRepository{}
	endpoint := web.TodoEndPoint{TodoRepository: &repository}

	initDatabase(repository)

	http.Handle("/todos", endpoint)
	http.HandleFunc("/echo", handleEcho)
	http.ListenAndServe(":3000", nil)
}
func handleEcho(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Echo endpoint is called")
	if r.URL.Path == "/echo" {
		switch r.Method {
		case http.MethodGet:
			w.Write([]byte("echo"))
		}
	}
}

func initDatabase(repository model.InMemoryTodoRepository) {
	repository.SaveTodo(model.Todo{
		Id:      "id",
		Content: "a content",
	})
	todoList, _ := repository.GetAllTodo()
	for _, value := range todoList {
		fmt.Println(*value)
	}
}
