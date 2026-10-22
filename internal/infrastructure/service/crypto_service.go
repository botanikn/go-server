package service

import (
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

type CryptoService struct{}

func NewCryptoService() *CryptoService {
	return &CryptoService{}
}

func (c *CryptoService) Hash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (c *CryptoService) Compare(password string, passwordHash string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
}

func (c *CryptoService) CreateToken() (string, error) {
	randomBytes := make([]byte, 32)

	if _, err := rand.Read(randomBytes); err != nil {
		return "", err
	}

	token := base64.RawURLEncoding.EncodeToString(randomBytes)
	return token, nil
}
