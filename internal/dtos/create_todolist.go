package dtos

import (
	"simple-todo-list/internal/domain/todo_list"
)

type CreateTodoListInput struct {
	Name string `json:"name" validate:"required"`
}

type CreateTodoListOutput struct {
	Id string
}

func NewDtoCreateTodoListOutput(list todo_list.TodoList) CreateTodoListOutput {
	return CreateTodoListOutput{Id: list.Id}
}
