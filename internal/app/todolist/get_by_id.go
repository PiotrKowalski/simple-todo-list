package todolist

import (
	"context"
	"simple-todo-list/internal/domain"
	"simple-todo-list/internal/domain/todo_list"
	"simple-todo-list/internal/dtos"
)

type GetByIdTodoListHandler struct {
	Repo domain.GenericGetByIdRepo[todo_list.TodoList]
}

func (h *GetByIdTodoListHandler) Handle(ctx context.Context, input dtos.GetByIdTodoListInput) (dtos.GetByIdTodoListOutput, error) {
	list, err := h.Repo.GetById(ctx, input.Id)
	if err != nil {
		return dtos.GetByIdTodoListOutput{}, err
	}

	return dtos.NewDtoGetByIdTodoListOutput(*list), nil
}
