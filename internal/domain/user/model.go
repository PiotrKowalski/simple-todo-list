package user

type Role string

type User struct {
	Id       string `bson:"_id"`
	Username string `bson:"username"`
	Password string `bson:"password"`
	Email    string `bson:"email"`
	Role     Role   `bson:"role"`
}

func (u User) IsAggregateRoot() {}
