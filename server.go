package main

import (
	"github.com/go-martini/martini"
	"myproject/controller"
	"myproject/entity"
)

func main() {
	m := martini.Classic()

	todoList := make(map[string]entity.Todo)
	todo := controller.NewTodo(todoList)

	m.Group("/todo", func(r martini.Router) {
		r.Post("", todo.Post)
		r.Get("", todo.Get)
	})

	m.Run()
}
