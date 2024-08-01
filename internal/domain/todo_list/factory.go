package todo_list

import "github.com/google/uuid"

type Option func(*TodoList)

func NewTodoList(name string, opts ...Option) *TodoList {
	t := &TodoList{
		Id:   uuid.New().String(),
		Name: name,
	}
	for _, opt := range opts {
		opt(t)
	}
	return t
}

func WithId(id string) Option {
	return func(l *TodoList) {
		l.Id = id
	}
}
