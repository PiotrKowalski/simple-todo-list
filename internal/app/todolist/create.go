package todolist

import (
	"context"
	"simple-todo-list/internal/domain"
	"simple-todo-list/internal/domain/todo_list"
	"simple-todo-list/internal/dtos"
)

type CreateTodoListHandler struct {
	Repo domain.GenericSaveRepo[todo_list.TodoList]
}

func (h *CreateTodoListHandler) Handle(ctx context.Context, input dtos.CreateTodoListInput) (dtos.CreateTodoListOutput, error) {
	todoList, err := h.Repo.Save(ctx, todo_list.NewTodoList(input.Name))
	if err != nil {
		return dtos.CreateTodoListOutput{}, err
	}

	return dtos.NewDtoCreateTodoListOutput(*todoList), nil
}
