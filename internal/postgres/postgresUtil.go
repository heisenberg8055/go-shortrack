package postgres

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func InsertData(conn *pgx.Conn, longURL, shortURL string) int64 {
	var lastInsertID int64
	row, err := conn.Query(context.Background(), "INSERT INTO gotiny (longurl, shorturl) VALUES($1, $2) RETURNING ID", longURL, shortURL)
	if err != nil {
		log.Printf("Failed to insert DB: %v", err)
		return -1
	}
	row.Scan(&lastInsertID)
	return lastInsertID
}
