package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/todos", todosHandler)
	http.HandleFunc("/todos/", todosIdHandler)
}

func todosHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getAllTodos(w, r)
	case "POST":
		addNewTodo(w, r)
	}
}

func todosIdHandler(w http.ResponseWriter, r *http.Request) {

}

func getAllTodos() {

}

func addNewTodo() {
	
}
