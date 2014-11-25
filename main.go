package main

import (
	"controller"
	"github.com/go-martini/martini"
	"repository"
)

func main() {
	app := martini.Classic()
	repo := repository.NewTodo()
	todo := controller.NewTodo(repo)

	app.Group("/todo", func(route martini.Router) {
		route.Post("", todo.Post)
		route.Get("", todo.Get)
	})

	app.Run()
}
