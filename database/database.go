package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func DBConnection() (connection *sql.DB) {

	DRIVER := "mysql"
	USER := "u3ld0sazbrzmmcpa"
	PASS := "2pCUIaAlXSm2F7PQKFa0"
	DB_NAME := "bkabj8wc6zhqrkc6juvw"
	HOST := "bkabj8wc6zhqrkc6juvw-mysql.services.clever-cloud.com"

	conn, err := sql.Open(DRIVER, USER+":"+PASS+"@tcp("+HOST+")/"+DB_NAME)
	if err != nil {
		panic(err.Error())
	}

	return conn
}
