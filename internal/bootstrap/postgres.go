package bootstrap

import (
	sql "database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func CreatePostgres() (*sql.DB, error) {
	host := "localhost"
	port := 5432
	user := "social-network-user"
	password := "social-network-password"
	dbname := "social_network_otus"
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
