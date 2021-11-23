package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func DBConnection() (conecction *sql.DB) {
	Driver := "mysql"
	User := "u3ld0sazbrzmmcpa"
	Password := "2pCUIaAlXSm2F7PQKFa0"
	DBName := "bkabj8wc6zhqrkc6juvw"
	Host := "bkabj8wc6zhqrkc6juvw-mysql.services.clever-cloud.com"

	conecction, err := sql.Open(Driver, User+":"+Password+"@tcp("+Host+")/"+DBName)
	if err != nil {
		panic(err.Error())
	}

	return conecction
}
