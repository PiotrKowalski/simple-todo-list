package domain

import (
	"context"
	"simple-todo-list/pkg/specification"
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

type GenericGetForMatchingRepo[T AggregateRoot] interface {
	FindForMatching(ctx context.Context, specification specification.Specification[T]) ([]T, error)
}
