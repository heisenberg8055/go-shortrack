package postgres

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InsertData(conn *pgxpool.Pool, longURL, shortURL string) {
	_, err := conn.Exec(context.Background(), "INSERT INTO gotiny (longurl, shorturl) VALUES($1, $2)", longURL, shortURL)
	if err != nil {
		log.Printf("Failed to insert DB: %v", err)
	}
}

func FetchShortUrl(conn *pgxpool.Pool, longURL string) string {
	var shorturl string
	row := conn.QueryRow(context.Background(), "SELECT SHORTURL FROM GOTINY WHERE LONGURL = $1", longURL)
	row.Scan(&shorturl)
	return shorturl
}

func ValidateHash(conn *pgxpool.Pool, shortURL string) bool {
	if shortURL == "" {
		return true
	}
	var ok bool
	row := conn.QueryRow(context.Background(), "SELECT COUNT(shorturl) > 0 FROM gotiny WHERE shorturl = $1", shortURL)
	_ = row.Scan(&ok)
	return ok
}
