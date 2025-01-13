package postgres

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDB(config *map[string]string) *pgxpool.Pool {
	if dbURL, ok := (*config)["POSTGRES_URL"]; ok {
		conn, err := pgxpool.New(context.Background(), dbURL)
		if err != nil {
			log.Fatalf("Failed to connect DB: %v", err)
		}
		initDB(conn)
		return conn
	}
	userName := (*config)["POSTGRES_USER"]
	passWord := (*config)["POSTGRES_PASSWORD"]
	hostName := (*config)["POSTGRES_HOST"]
	dbName := (*config)["POSTGRES_DATABASE"]
	port := (*config)["POSTGRES_PORT"]
	connString := "postgres://" + userName + ":" + passWord + "@" + hostName + ":" + port + "/" + dbName
	conn, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %s", err)
	}
	initDB(conn)
	return conn
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
