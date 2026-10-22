package repository

import (
	"context"

	"github.com/botanikn/go-server/internal/entities"
	"github.com/botanikn/go-server/internal/storage"
)

type AuthRepository struct {
	storage *storage.Storage
}

func NewAuthRepository(storage *storage.Storage) *AuthRepository {
	return &AuthRepository{
		storage: storage,
	}
}

func (a *AuthRepository) SaveUser(username string, passwordHash string, email string) error {
	query := "INSERT INTO users (username, pass_hash, email) VALUES ($1, $2, $3)"

	ctx := context.Background()

	_, err := a.storage.Connection.Exec(ctx, query, username, passwordHash, email)

	if err != nil {
		return err
	}

	return nil
}

func (a *AuthRepository) GetUser(username string) (entities.User, error) {
	query := "SELECT id, username, pass_hash, email FROM users WHERE username = $1"

	ctx := context.Background()

	row := a.storage.Connection.QueryRow(ctx, query, username)

	var user entities.User
	err := row.Scan(&user.Id, &user.Username, &user.PassHash, &user.Email)
	if err != nil {
		return entities.User{}, err
	}

	return user, nil
}

func (a *AuthRepository) SaveToken(userId int, token string, exp int) error {
	query := "INSERT INTO tokens (user_id, token, exp) VALUES ($1, $2, $3)"

	ctx := context.Background()

	_, err := a.storage.Connection.Exec(ctx, query, userId, token, exp)

	if err != nil {
		return err
	}

	return nil
}
