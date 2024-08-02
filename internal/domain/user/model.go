package user

import "simple-todo-list/internal/domain"

type Role string

type User struct {
	domain.AggregateRoot `bson:",omitempty"`
	Id                   string `bson:"_id"`
	Username             string `bson:"username"`
	Password             string `bson:"password"`
	Email                string `bson:"email"`
	Role                 Role   `bson:"role"`
}
