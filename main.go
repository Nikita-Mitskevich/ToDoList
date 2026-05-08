package main

import (
	"errors"
	"fmt"
	"os"
	"restapi/http"
	"restapi/todo"
)

func main() {
	todoList, err := todo.CreateNewList()
	if err != nil {
		panic(err)
	}
	HTTPHandlers := http.CreateHTTPHandlers(todoList)
	HTTPServer := http.NewHTTPServer(HTTPHandlers)

	if err := HTTPServer.StartServer(); err != nil {
		fmt.Fprintln(os.Stderr, errors.New("Сервер не запустился, произошла ошибка"))
	}
}
