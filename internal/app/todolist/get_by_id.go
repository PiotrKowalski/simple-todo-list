package todolist

import (
	"context"
	"simple-todo-list/internal/domain"
	"simple-todo-list/internal/domain/todo_list"
	"simple-todo-list/internal/dtos"
	"time"
)

type GetByIdTodoListHandler struct {
	Repo domain.GenericGetByIdRepo[todo_list.TodoList]
}

func (h *GetByIdTodoListHandler) Handle(ctx context.Context, input dtos.GetByIdTodoListInput) (dtos.GetByIdTodoListOutput, error) {
	deadline, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	_, err := h.Repo.GetById(deadline, input.Id)
	if err != nil {
		return dtos.GetByIdTodoListOutput{}, err
	}

	return dtos.GetByIdTodoListOutput{}, nil
}
