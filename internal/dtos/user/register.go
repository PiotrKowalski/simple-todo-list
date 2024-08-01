package user

type RegisterInput struct {
	Username string `json:"username" validate:"required,gte=1"`
	Password string `json:"password" validate:"required,gte=1"`
	Email    string `json:"email" validate:"required,email"`
}

type RegisterOutput struct {
	Id  string `json:"id"`
	JWT string `json:"jwt"`
}

func NewRegisterOutput(id, jwt string) RegisterOutput {
	return RegisterOutput{Id: id, JWT: jwt}
}
