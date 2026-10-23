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

func (a *AuthRepository) SaveUser(username string, passwordHash string, email string, role string) error {
	query := "INSERT INTO users (username, passwordhash, email, role) VALUES ($1, $2, $3, $4)"

	ctx := context.Background()

	_, err := a.storage.Connection.Exec(ctx, query, username, passwordHash, email, role)

	if err != nil {
		return err
	}

	return nil
}

func (a *AuthRepository) GetUserByName(username string) (entities.User, error) {
	query := "SELECT id, username, passwordhash, email FROM users WHERE username = $1"

	ctx := context.Background()

	row := a.storage.Connection.QueryRow(ctx, query, username)

	var user entities.User
	err := row.Scan(&user.Id, &user.Username, &user.PassHash, &user.Email)
	if err != nil {
		return entities.User{}, err
	}

	return user, nil
}

func (a *AuthRepository) GetUserRoleByUserId(userId int) (string, error) {
	query := "SELECT role FROM users WHERE id = $1"

	ctx := context.Background()

	row := a.storage.Connection.QueryRow(ctx, query, userId)

	var role string
	err := row.Scan(&role)
	if err != nil {
		return "", err
	}

	return role, nil
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

func (a *AuthRepository) GetTokenRecord(token string) (entities.Token, error) {
	query := "SELECT id, user_id, exp, token FROM tokens WHERE token = $1"

	ctx := context.Background()

	row := a.storage.Connection.QueryRow(ctx, query, token)

	var tokenRecord entities.Token
	err := row.Scan(&tokenRecord.Id, &tokenRecord.UserId, &tokenRecord.Exp, &tokenRecord.Token)
	if err != nil {
		return entities.Token{}, err
	}

	return tokenRecord, nil
}
