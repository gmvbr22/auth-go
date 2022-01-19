package password

type PasswordService interface {
	Generate(password []byte) ([]byte, error)
	Compare(hashedPassword, password []byte) error
}
