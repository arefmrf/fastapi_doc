package main

import (
	"database/sql"
	"log"
	"os"
	"time"
)

const webPort = 80

func main() {
	initDB()
}

func initDB() *sql.DB {
	conn := connectToDB()
	if conn == nil {
		log.Panic("Can't connect to database")
	}
}

func connectToDB() *sql.DB {
	counts := 0
	dsn := os.Getenv("DSN")
	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("postgres not ready")
		} else {
			return connection
		}
		if counts > 10 {
			return nil
		}
		log.Println("waiting for database to become available 1 second")
		time.Sleep(1 * time.Second)
		counts++
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
