package hash

import "golang.org/x/crypto/bcrypt"

type Hash interface {
	Hash(string) (string, error)
	Verify(hashed, data string) (bool, error)
}

type hash struct{}

func New() Hash {
	return &hash{}
}

func (h *hash) Hash(data string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(data), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bytes), err
}

func (h *hash) Verify(hashed, data string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(data))
	if err != nil {
		return false, err
	}
	return true, nil
}
