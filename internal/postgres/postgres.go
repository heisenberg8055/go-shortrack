package postgres

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	Db *pgxpool.Pool
}

var (
	pgInstance *Postgres
	pgOnce     sync.Once
)

func ConnectDB() (*Postgres, error) {
	userName := os.Getenv("POSTGRES_USER")
	passWord := os.Getenv("POSTGRES_PASSWORD")
	hostName := os.Getenv("POSTGRES_HOST")
	dbName := os.Getenv("POSTGRES_DATABASE")
	port := os.Getenv("POSTGRES_PORT")
	connString := "postgres://" + userName + ":" + passWord + "@" + hostName + ":" + port + "/" + dbName
	pgOnce.Do(func() {
		db, err := pgxpool.New(context.Background(), connString)
		if err != nil {
			log.Fatalf("Failed to connect to a connection pool: %s", err)
			return
		}
		pgInstance = &Postgres{Db: db}
	})
	initDB(pgInstance.Db)
	return pgInstance, nil
}

func initDB(conn *pgxpool.Pool) {
	tableCreate := `CREATE TABLE IF NOT EXISTS gotiny (
		ID BIGSERIAL PRIMARY KEY,
					  LONGURL TEXT NOT NULL,
								  SHORTURL TEXT NOT NULL,
											  INSERT_AT TIMESTAMP NOT NULL DEFAULT NOW(),
														  COUNT INTEGER NOT NULL DEFAULT 0
																	);`
	res, err := conn.Exec(context.Background(), tableCreate)
	if err != nil {
		log.Fatalf("Relation unable to create: %s", err)
	}
	log.Println(res.String())
}

func (pg *Postgres) Ping() error {
	return pg.Db.Ping(context.Background())
}

func (pg *Postgres) Close() {
	pg.Db.Close()
}
