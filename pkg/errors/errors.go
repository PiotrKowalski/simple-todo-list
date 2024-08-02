package errors

import "fmt"

type CustomError interface {
}

type FactoryInvariantNotMetError struct {
	Invariant string
	Factory   string
}

func (e FactoryInvariantNotMetError) Error() string {
	return fmt.Sprintf("Factory Invariant %s not met: %s", e.Invariant, e.Factory)
}
