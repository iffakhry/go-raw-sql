package config

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

func ConnectToDB() *sql.DB {
	// <username>:<password>@tcp(<hostname>:<port>)/<db_name>
	var connectionString = "root:qwerty123@tcp(127.0.0.1:3306)/db_be15"
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal("error open connection", err.Error())
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	errPing := db.Ping()
	if errPing != nil {
		log.Fatal("error connect to db", errPing.Error())
	} else {
		fmt.Println("berhasil")
	}

	return db

}
