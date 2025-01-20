package postgres

import (
	"context"
	"log"
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

func ConnectDB(config *map[string]string) (*Postgres, error) {
	if dbURL, ok := (*config)["POSTGRES_URL"]; ok {
		pgOnce.Do(func() {
			db, err := pgxpool.New(context.Background(), dbURL)
			if err != nil {
				log.Fatalf("Failed to connect to a connection pool: %s", err)
				return
			}
			pgInstance = &Postgres{db}
		})
		initDB(pgInstance.Db)
		return pgInstance, nil
	} else {
		userName := (*config)["POSTGRES_USER"]
		passWord := (*config)["POSTGRES_PASSWORD"]
		hostName := (*config)["POSTGRES_HOST"]
		dbName := (*config)["POSTGRES_DATABASE"]
		port := (*config)["POSTGRES_PORT"]
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
}

func initDB(conn *pgxpool.Pool) {
	tableCreate := `CREATE TABLE IF NOT EXISTS gotiny (
		ID BIGSERIAL PRIMARY KEY,
					  LONGURL TEXT NOT NULL,
								  SHORTURL VARCHAR(7) NOT NULL,
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
