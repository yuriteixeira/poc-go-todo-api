package controller

import (
	"entity"
	"fmt"
	"net/http"
	"repository"
)

type Todo struct {
	repo *repository.Todo
}

func NewTodo(repo *repository.Todo) *Todo {
	t := new(Todo)
	t.repo = repo

	return t
}

func (t *Todo) Post(res http.ResponseWriter, req *http.Request) {
	item := entity.Todo{}
	item.Text = req.FormValue("text")

	t.repo.Add(item)

	res.WriteHeader(200)
}

func (t *Todo) Get(res http.ResponseWriter) string {
	response := ""
	data := t.repo.List()

	if len(data) < 1 {
		res.WriteHeader(204)
	}

	for k := range data {
		response += fmt.Sprintf("%d: %s\n", data[k].Id, data[k].Text)
	}

	return response
}
