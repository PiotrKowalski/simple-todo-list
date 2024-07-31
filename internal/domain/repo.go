package domain

import (
	"context"
)

type AggregateRoot interface {
	IsAggregateRoot()
}

type GenericSaveRepo[T AggregateRoot] interface {
	Save(ctx context.Context, obj *T) (*T, error)
}

type GenericGetByIdRepo[T AggregateRoot] interface {
	GetById(ctx context.Context, id string) (*T, error)
}
