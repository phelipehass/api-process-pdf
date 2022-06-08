package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

func InitPostgres() *sql.DB {
	psqlInfo := postgresURIFromEnv()
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

func postgresURIFromEnv() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("HOST_POSTGRES"), os.Getenv("PORT_POSTGRES"), os.Getenv("USER_POSTGRES"),
		os.Getenv("PASSWORD_POSTGRES"), os.Getenv("DATABASE_NAME"))

}
