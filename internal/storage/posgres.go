package storage

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Storage struct {
	Connection *pgx.Conn
}

func NewStorage(connString string) (*Storage, error) {
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		return nil, err
	}

	return &Storage{
		Connection: conn,
	}, nil
}
