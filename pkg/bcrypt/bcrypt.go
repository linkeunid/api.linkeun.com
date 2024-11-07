package bcrypt

import (
	"golang.org/x/crypto/bcrypt"
)

type Bcrypt struct {
	salt string
}

func NewBcrypt(salt string) *Bcrypt {
	return &Bcrypt{salt: salt}
}

// HashPassword hashes the given password with a custom salt
func (b *Bcrypt) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password+b.salt), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

// ComparePassword compares the given password with the hashed password
func (b *Bcrypt) ComparePassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password+b.salt))
}
