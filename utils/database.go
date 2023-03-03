// database connection
package utils

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

const (
	host     = "postgres-server"
	port     = 5432
	user     = "postgres"
	password = "supersecretpassword"
	dbname   = "postgres"
)

func ConnectDB() *sqlx.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Successfully connected to database")
	return db
}
