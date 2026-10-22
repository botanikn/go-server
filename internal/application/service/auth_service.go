package service

import (
	"time"

	"github.com/botanikn/go-server/internal/entities"
)

type AuthRepositoryInterface interface {
	SaveUser(username string, passwordHash string, email string) error
	GetUser(username string) (entities.User, error)
	SaveToken(userId int, token string, exp int) error
}

type CryptoServiceInterface interface {
	Hash(randomString string) (string, error)
	Compare(password string, passwordHash string) error
	CreateToken() (string, error)
}

type AuthService struct {
	rep    AuthRepositoryInterface
	crypto CryptoServiceInterface
}

func NewAuthService(rep AuthRepositoryInterface, crypto CryptoServiceInterface) AuthService {
	return AuthService{
		rep:    rep,
		crypto: crypto,
	}
}

func (a *AuthService) Register(username string, password string, email string) error {
	passwordHash, err := a.crypto.Hash(password)

	if err != nil {
		return err
	}

	result := a.rep.SaveUser(username, passwordHash, email)

	return result
}

func (a *AuthService) Login(username string, password string) (string, error) {
	user, err := a.rep.GetUser(username)

	if err != nil {
		return "", err
	}

	compareResult := a.crypto.Compare(password, user.PassHash)

	if compareResult != nil {
		return "", err
	}

	token, err := a.crypto.CreateToken()

	if err != nil {
		return "", err
	}

	err = a.rep.SaveToken(user.Id, token, int(time.Now().UTC().Add(10*time.Minute).Unix()))

	return token, nil
}
