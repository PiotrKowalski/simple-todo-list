package app

import (
	"context"
	"simple-todo-list/internal/adapters/repo/mongodb"
	"simple-todo-list/internal/app/todolist"
	"simple-todo-list/internal/domain/todo_list"
	"simple-todo-list/internal/dtos"
)

type Config struct {
	MongoDBURI string
}

type Application struct {
	create  todolist.CreateTodoListHandler
	getById todolist.GetByIdTodoListHandler
}

func New(
	config Config,
) (Application, error) {
	ctx := context.Background()

	client, err := mongodb.NewClient(ctx, config.MongoDBURI)
	if err != nil {
		return Application{}, err
	}

	todolistRepo, err := mongodb.NewRepoAdapter[todo_list.TodoList](client, "todolist", "todolists")
	if err != nil {
		return Application{}, err
	}

	return Application{
		create:  todolist.CreateTodoListHandler{Repo: todolistRepo},
		getById: todolist.GetByIdTodoListHandler{Repo: todolistRepo},
	}, nil
}

func (a Application) CreateTodoList(ctx context.Context, in dtos.CreateTodoListInput) (dtos.CreateTodoListOutput, error) {
	return a.create.Handle(ctx, in)
}

func (a Application) GetByIdTodoList(ctx context.Context, in dtos.GetByIdTodoListInput) (dtos.GetByIdTodoListOutput, error) {
	return a.getById.Handle(ctx, in)
}
