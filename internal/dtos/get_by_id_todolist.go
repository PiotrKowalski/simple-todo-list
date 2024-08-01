package dtos

import "simple-todo-list/internal/domain/todo_list"

type GetByIdTodoListInput struct {
	Id string `param:"id"`
}

type GetByIdTodoListOutput struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Entries []entry
}

type entry struct {
	Id          uint   `json:"id" bson:"id"`
	Description string `json:"description" bson:"description"`
	Status      string `json:"status" bson:"status"`
}

func NewDtoGetByIdTodoListOutput(list todo_list.TodoList) GetByIdTodoListOutput {
	out := &GetByIdTodoListOutput{}
	out.Id = list.Id
	out.Name = list.Name

	for _, e := range list.Entries {
		out.Entries = append(out.Entries, entry{
			Id:          e.Id,
			Description: e.Description,
			Status:      string(e.Status),
		})
	}
	return *out
}
