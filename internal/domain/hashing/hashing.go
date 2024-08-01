package hashing

type HashService interface {
	Hash(string) (string, error)
}

type HashCheckService interface {
	CompareHashAndClear(string, string) error
}
