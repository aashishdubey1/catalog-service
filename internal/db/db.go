package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

func Connect() (*pgx.Conn, error) {
	dsn := "postgres://postgres:12345678@localhost:5432/catalog_service?sslmode=disable"

	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()

	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db : %w", err)
	}

	err = conn.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to ping to db : %w", err)
	}

	return conn, nil

}
