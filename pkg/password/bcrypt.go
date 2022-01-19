package password

import "golang.org/x/crypto/bcrypt"

type BcryptAdapter struct {
	cost int
}

func NewBcryptAdapter(cost int) PasswordService {
	return &BcryptAdapter{cost: cost}
}

func (adapter *BcryptAdapter) Generate(password []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, adapter.cost)
}

func (adapter *BcryptAdapter) Compare(hashedPassword, password []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, password)
}
