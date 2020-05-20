package web

import (
	"encoding/json"
	"fmt"
	"githab/mrflick72/go-playground/src/model"
	"io"
	"net/http"
)

type TodoEndPoint struct {
	TodoRepository model.TodoRepository
}

func (endpoint TodoEndPoint) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("todo is invoched")
	repository := endpoint.TodoRepository
	if r.URL.Path == "/todos" {
		switch r.Method {
		case http.MethodGet:
			todos, _ := repository.GetAllTodo()
			encodeResponseAsJSON(&todos, w)
		case http.MethodPost:
			u, _ := decodeRequestAsJSON(r)
			repository.SaveTodo(u)
			w.WriteHeader(http.StatusCreated)
		}
	}
}

func HandleEcho(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Echo endpoint is called")
	if r.URL.Path == "/echo" {
		switch r.Method {
		case http.MethodGet:
			w.Write([]byte("echo"))
		}
	}
}
func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}

func decodeRequestAsJSON(r *http.Request) (model.Todo, error) {
	dec := json.NewDecoder(r.Body)
	var todo model.Todo
	err := dec.Decode(&todo)
	r.Body.Close()
	if err != nil {
		return model.Todo{}, err
	}
	return todo, nil
}
