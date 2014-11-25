package controller

import (
	"entity"
	"fmt"
	"net/http"
)

type Todo struct {
	data map[string]entity.Todo
}

func NewTodo(data map[string]entity.Todo) *Todo {
	t := new(Todo)
	t.data = data
	return t
}

func (t *Todo) Post(res http.ResponseWriter, req *http.Request) {
	item := entity.Todo{}
	item.Id = req.FormValue("id")
	item.Text = req.FormValue("text")
	t.data[item.Id] = item

	res.WriteHeader(200)
}

func (t *Todo) Get(res http.ResponseWriter) string {
	var result string

	if len(t.data) < 1 {
		res.WriteHeader(418)
	}

	for k := range t.data {
		result += fmt.Sprintf("%s: %s\n", t.data[k].Id, t.data[k].Text)
	}

	return result
}
