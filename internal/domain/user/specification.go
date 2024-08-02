package user

type HasUsernameSpecification struct {
	Username string
}

func (spec HasUsernameSpecification) IsSatisfiedBy(candidate User) bool {
	return candidate.Username == spec.Username
}
