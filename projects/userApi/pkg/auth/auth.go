package auth

import "golang.org/x/crypto/bcrypt"

func Encrypt(source string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(source), bcrypt.DefaultCost)
	return string(hashedBytes), err
}

func Compare(hashedSource, source string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedSource), []byte(source))
}
