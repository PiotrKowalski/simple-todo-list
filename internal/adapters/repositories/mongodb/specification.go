package mongodb

import (
	"go.mongodb.org/mongo-driver/bson"
	"simple-todo-list/internal/domain/user"
	"simple-todo-list/pkg/specification"
)

func ToBsonD[T any](specifications any) bson.D {
	switch spec := specifications.(type) {
	case *specification.AndSpecification[T]:
		return bson.D{{Key: "$and", Value: bson.A{ToBsonD[T](spec.Left), ToBsonD[T](spec.Right)}}}
	case *specification.OrSpecification[T]:
		return bson.D{{"$or", bson.A{ToBsonD[T](spec.Left), ToBsonD[T](spec.Right)}}}
	case *specification.NotSpecification[T]:
		return bson.D{{Key: "$not", Value: ToBsonD[T](spec.Spec)}}
	case user.HasUsernameSpecification:
		username := specifications.(user.HasUsernameSpecification).Username
		return bson.D{{Key: "username", Value: username}}
	default:
		return bson.D{}
	}
}
