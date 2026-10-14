package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

func CreateTables(conn *pgx.Conn) error {
	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()
	query := `
	CREATE EXTENSION IF NOT EXISTS "pgcrypto";

	CREATE TABLE IF NOT EXISTS products (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		title TEXT NOT NULL,
		description TEXT NOT NULL,
		price NUMERIC(10,2) NOT NULL,
		created_at TIMESTAMPTZ DEFAULT NOW()
	);`

	_, err := conn.Exec(ctx, query)
	if err != nil {
		return fmt.Errorf("failed to run query: %w", err)
	}

	fmt.Println("db schema initialized")
	return nil
}

func SeedProducts(conn *pgx.Conn) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Optional: Check if products already exist so you don't duplicate them on every restart
	var count int
	err := conn.QueryRow(ctx, "SELECT COUNT(*) FROM product").Scan(&count)
	if err == nil && count > 0 {
		fmt.Println("Products already seeded, skipping...")
		return nil
	}

	var batch pgx.Batch

	// Queue up 1,000 inserts
	for i := 1; i <= 1000; i++ {
		title := fmt.Sprintf("Awesome Product %d", i)
		description := fmt.Sprintf("This is a high-quality description for product number %d.", i)
		price := float64(i)*5.50 + 9.99 // Generates varied prices

		batch.Queue(
			"INSERT INTO product (title, description, price) VALUES ($1, $2, $3)",
			title, description, price,
		)
	}

	// Send the entire batch to PostgreSQL in one go
	br := conn.SendBatch(ctx, &batch)
	defer br.Close()

	// Verify every statement in the batch executed successfully
	for i := 1; i <= 1000; i++ {
		_, err := br.Exec()
		if err != nil {
			return fmt.Errorf("failed to execute batch item %d: %w", i, err)
		}
	}

	fmt.Println("Successfully seeded 1,000 products!")
	return nil
}
