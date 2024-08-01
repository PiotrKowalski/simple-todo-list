package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"simple-todo-list/internal/domain"
)

type Repo[T domain.AggregateRoot] struct {
	client     *mongo.Client
	database   string
	collection string
}

func (t Repo[T]) Save(ctx context.Context, agg *T) (*T, error) {
	coll := t.client.Database(t.database).Collection(t.collection)

	_, err := coll.InsertOne(ctx, agg)
	if err != nil {
		return nil, err
	}
	return agg, nil
}

func (t Repo[T]) GetById(ctx context.Context, id string) (*T, error) {
	coll := t.client.Database(t.database).Collection(t.collection)

	filter := bson.D{{"_id", id}}
	res := coll.FindOne(ctx, filter)
	if err := res.Err(); err != nil {
		return nil, err
	}
	var obj T
	err := res.Decode(&obj)
	if err != nil {
		return nil, err
	}

	return &obj, nil
}

func (t Repo[T]) Remove(id string) error {
	//TODO implement me
	panic("implement me")
}

func NewRepo[T domain.AggregateRoot](client *mongo.Client, coll string) (*Repo[T], error) {
	return &Repo[T]{
		client:     client,
		database:   "todolist",
		collection: coll,
	}, nil
}
