package dtos

type CreateTodoListInput struct {
	Name string `json:"name" validate:"required"`
}

type CreateTodoListOutput struct {
	Id string
}
