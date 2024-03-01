package services

import (
	"unicode/utf8"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	PasswordHash(password string) (string, error)
	TokenGenerator()
	TokenExtractor()
}

type AuthServiceImpl struct{}

func (a *AuthServiceImpl) PasswordHash(password string) (string, error) {
	length := utf8.RuneCountInString(password)
	if length == 0 {
		return "", nil
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password+"salt"), 10)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (a *AuthServiceImpl) TokenGenerator() {

}

func (a *AuthServiceImpl) TokenExtractor() {

}
