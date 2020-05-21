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
			endpoint.getAll(w, repository)
		case http.MethodPost:
			endpoint.saveATodo(w, r, repository)

		case http.MethodDelete:
			endpoint.deleteATodo(w, r, repository)
		}
	}
}

func (endpoint *TodoEndPoint) getAll(w http.ResponseWriter, repository model.TodoRepository) {
	todos, _ := repository.GetAllTodo()
	encodeResponseAsJSON(&todos, w)
}

func (endpoint TodoEndPoint) saveATodo(w http.ResponseWriter, r *http.Request, repository model.TodoRepository) {
	u, _ := decodeRequestAsJSON(r)
	repository.SaveTodo(&u)
	w.WriteHeader(http.StatusCreated)
}

func (endpoint TodoEndPoint) deleteATodo(w http.ResponseWriter, r *http.Request, repository model.TodoRepository) {
	//todo	w.WriteHeader(http.StatusCreated)
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
