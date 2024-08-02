package specification

type Specification[T any] interface {
	IsSatisfiedBy(candidate T) bool
}

type CompositeSpecification[T any] struct{}

func (spec CompositeSpecification[T]) And(other Specification[T]) Specification[T] {
	return AndSpecification[T]{CompositeSpecification: spec, Left: other}
}

func (spec CompositeSpecification[T]) Or(other Specification[T]) Specification[T] {
	return OrSpecification[T]{CompositeSpecification: spec, Left: other}
}

func (spec CompositeSpecification[T]) Not() Specification[T] {
	return NotSpecification[T]{CompositeSpecification: spec}
}

type AndSpecification[T any] struct {
	CompositeSpecification[T]
	Left, Right Specification[T]
}

func (spec AndSpecification[T]) IsSatisfiedBy(candidate T) bool {
	return spec.Left.IsSatisfiedBy(candidate) && spec.Right.IsSatisfiedBy(candidate)
}

type OrSpecification[T any] struct {
	CompositeSpecification[T]
	Left, Right Specification[T]
}

func (spec OrSpecification[T]) IsSatisfiedBy(candidate T) bool {
	return spec.Left.IsSatisfiedBy(candidate) || spec.Right.IsSatisfiedBy(candidate)
}

type NotSpecification[T any] struct {
	CompositeSpecification[T]
	Spec Specification[T]
}

func (spec NotSpecification[T]) IsSatisfiedBy(candidate T) bool {
	return !spec.Spec.IsSatisfiedBy(candidate)
}
