package logic 

import(
	"context"
	"golang.org/x/crypto/bcrypt"
	"errors"
)

type Hash interface{
	Hash(ctx context.Context, password string) (string, error)
	IsHashEqual(ctx context.Context, password string, hashed string) (bool, error)
}

type hash struct{

}

func NewHash() Hash{
	return &hash{

	}
}

func (h *hash) Hash(ctx context.Context, password string) (string, error){
	hashed, err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("failed to create hash data")
	}
	return string(hashed), nil 
}

func (h *hash) IsHashEqual(ctx context.Context, password string, hashed string) (bool, error){
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password)); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return false, nil
		}

		return false,errors.New("failed to check if data equal hash")
	}

	return true, nil
}