package todo_list

import (
	"github.com/google/uuid"
	"simple-todo-list/internal/dtos"
)

func NewFromCreateTodoListRequest(request dtos.CreateTodoListInput) *TodoList {
	id := uuid.New().String()
	return &TodoList{
		Id:   id,
		Name: request.Name,
	}

}
