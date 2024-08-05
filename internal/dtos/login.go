package dtos

type LoginInput struct {
	Username string `json:"username" validate:"required,gte=1"`
	Password string `json:"password" validate:"required,gte=1"`
}

type LoginOutput struct {
	JWT          string `json:"jwt"`
	RefreshToken string `json:"refresh_token"`
}

func NewLoginOutput(jwt, refreshToken string) LoginOutput {
	return LoginOutput{JWT: jwt, RefreshToken: refreshToken}
}
