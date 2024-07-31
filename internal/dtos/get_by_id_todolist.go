package dtos

type GetByIdTodoListInput struct {
	Id string `param:"id"`
}

type GetByIdTodoListOutput struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Entries struct {
		Id          uint   `json:"id" bson:"id"`
		Description string `json:"description" bson:"description"`
		Status      string `json:"status" bson:"status"`
	}
}
