package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InsertData(conn *pgxpool.Pool, longURL, shortURL string) error {
	insertQuery := `INSERT INTO GOTINY (longurl, shorturl) VALUES (@longurl, @shorturl)`
	args := pgx.NamedArgs{
		"longurl":  longURL,
		"shorturl": shortURL,
	}
	_, err := conn.Exec(context.Background(), insertQuery, args)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}
	return nil
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

func FetchLongUrl(conn *pgxpool.Pool, shortURL string) string {
	var longurl string
	row := conn.QueryRow(context.Background(), "SELECT LONGURL FROM GOTINY WHERE SHORTURL = $1", shortURL)
	row.Scan(&longurl)
	return longurl
}
