package repository

import "entity"

type Todo struct {
	last int
	data map[int]entity.Todo
}

func NewTodo() *Todo {
	t := new(Todo)
	t.data = make(map[int]entity.Todo)
	return t
}

func (t *Todo) Add(item entity.Todo) {
	item.Id = t.last + 1
	t.data[item.Id] = item
	t.last = item.Id
}

func (t *Todo) List() map[int]entity.Todo {
	return t.data
}
