package app

import (
	"context"
	"simple-todo-list/internal/adapters/repositories/mongodb"
	"simple-todo-list/internal/adapters/services/auth"
	"simple-todo-list/internal/adapters/services/hashing"
	"simple-todo-list/internal/app/todolist"
	userHandlers "simple-todo-list/internal/app/user"
	"simple-todo-list/internal/domain/todo_list"
	userDomain "simple-todo-list/internal/domain/user"
	"simple-todo-list/internal/dtos"
	"time"
)

type Config struct {
	MongoDBURI string
}

type Application struct {
	create   todolist.CreateTodoListHandler
	getById  todolist.GetByIdTodoListHandler
	register userHandlers.RegisterHandler
	login    userHandlers.LoginHandler
}

func New(
	config Config,
) (Application, error) {
	ctx := context.Background()

	client, err := mongodb.NewClient(ctx, config.MongoDBURI)
	if err != nil {
		return Application{}, err
	}

	todolistRepo, err := mongodb.NewRepo[todo_list.TodoList](client, "todolists")
	if err != nil {
		return Application{}, err
	}
	userRepo, err := mongodb.NewRepo[userDomain.User](client, "users")
	if err != nil {
		return Application{}, err
	}

	hashingService := hashing.New()
	authService := auth.New()
	return Application{
		create:   todolist.CreateTodoListHandler{Repo: todolistRepo},
		getById:  todolist.GetByIdTodoListHandler{Repo: todolistRepo},
		register: userHandlers.RegisterHandler{Repo: userRepo, HashService: hashingService, AuthService: authService},
		login:    userHandlers.LoginHandler{Repo: userRepo, HashService: hashingService, AuthService: authService},
	}, nil
}

func (a Application) CreateTodoList(ctx context.Context, in dtos.CreateTodoListInput) (dtos.CreateTodoListOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return a.create.Handle(ctx, in)
}

func (a Application) GetByIdTodoList(ctx context.Context, in dtos.GetByIdTodoListInput) (dtos.GetByIdTodoListOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return a.getById.Handle(ctx, in)
}

func (a Application) Register(ctx context.Context, in dtos.RegisterInput) (dtos.RegisterOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return a.register.Handle(ctx, in)
}

func (a Application) Login(ctx context.Context, in dtos.LoginInput) (dtos.LoginOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return a.login.Handle(ctx, in)
}
