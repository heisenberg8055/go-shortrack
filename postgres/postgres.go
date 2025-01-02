package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func ConnectDB(config *map[string]string) {
	userName := (*config)["POSTGRES_USERNAME"]
	passWord := (*config)["POSTGRES_PASSWORD"]
	hostName := (*config)["POSTGRES_HOSTNAME"]
	dbName := (*config)["POSTGRES_DBNAME"]
	port := (*config)["POSTGRES_PORT"]
	connstring := "postgres://" + userName + ":" + passWord + "@" + hostName + ":" + port + "/" + dbName
	conn, err := pgx.Connect(context.Background(), connstring)
	if err != nil {
		panic(err)
	}
	defer conn.Close(context.Background())
	rows, err := conn.Query(context.Background(), "select longurl, shorturl from gotiny;")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var longurl, shorturl string
		rows.Scan(&longurl, &shorturl)
		fmt.Println(longurl, shorturl)
	}
}
