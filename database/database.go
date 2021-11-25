package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func GoDotEnvVar(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func DBConnection() (conecction *sql.DB) {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	Driver := GoDotEnvVar("DRIVER")
	User := GoDotEnvVar("USER")
	Password := GoDotEnvVar("PASSWORD")
	DB_Name := GoDotEnvVar("DB_NAME")
	Host := GoDotEnvVar("HOST")

	conn, err := sql.Open(Driver, User+":"+Password+"@tcp("+Host+")/"+DB_Name)
	if err != nil {
		panic(err.Error())
	}

	return conn
}
