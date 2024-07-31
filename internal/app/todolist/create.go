package todolist

import (
	"context"
	"simple-todo-list/internal/domain"
	"simple-todo-list/internal/domain/todo_list"
	"simple-todo-list/internal/dtos"
	"time"
)

type CreateTodoListHandler struct {
	Repo domain.GenericSaveRepo[todo_list.TodoList]
}

func (h *CreateTodoListHandler) Handle(ctx context.Context, input dtos.CreateTodoListInput) (dtos.CreateTodoListOutput, error) {
	deadline, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	todoList, err := h.Repo.Save(deadline, todo_list.NewFromCreateTodoListRequest(input))
	if err != nil {
		return dtos.CreateTodoListOutput{}, err
	}

	return dtos.CreateTodoListOutput{todoList.GetId()}, nil
}
