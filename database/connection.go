package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	dsn := "root:@tcp(127.0.0.1:3306)/e_commerce"
	// os.Getenv("DB_USER"),
	// // os.Getenv("DB_PASS"),
	// os.Getenv("DB_HOST"),
	// os.Getenv("DB_PORT"),
	// os.Getenv("DB_NAME"),

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	return db, nil
}
