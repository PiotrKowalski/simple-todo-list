package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"simple-todo-list/internal/domain"
)

type RepoAdapter[T domain.AggregateRoot] struct {
	client     *mongo.Client
	database   string
	collection string
}

func (t RepoAdapter[T]) Save(ctx context.Context, agg *T) (*T, error) {
	coll := t.client.Database(t.database).Collection(t.collection)

	_, err := coll.InsertOne(ctx, agg)
	if err != nil {
		return nil, err
	}
	return agg, nil
}

func (t RepoAdapter[T]) GetById(ctx context.Context, id string) (*T, error) {
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

func (t RepoAdapter[T]) Remove(id string) error {
	//TODO implement me
	panic("implement me")
}

func NewRepoAdapter[T domain.AggregateRoot](client *mongo.Client, db, coll string) (*RepoAdapter[T], error) {
	return &RepoAdapter[T]{
		client:     client,
		database:   db,
		collection: coll,
	}, nil
}
