package database

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func ConnectDataBase() (*sqlx.DB, error) {
	var err error
	DB, err = sqlx.Connect("postgres", "user=postgres dbname=tny sslmode=disable password=root host=localhost")
	if err != nil {
		log.Panicln("Database connection error", err)
		return nil, err
	}
	if err := DB.Ping(); err != nil {
		log.Println("Database Ping failed", err)
		return nil, err
	}
	log.Println("Connected")
	return DB, nil

}
