package config

import "fmt"

type EnvVarNotSetError struct {
	key string
}

func (e EnvVarNotSetError) Error() string {
	return fmt.Sprintf("key=%s is not set", e.key)
}

type EnvVarNotValid struct {
	key string
}

func (e EnvVarNotValid) Error() string {
	return fmt.Sprintf("key=%s is not valid", e.key)
}
