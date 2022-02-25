package db

import (
	"database/sql"
	"fmt"

	H "../helper"
)


const (
    DB_USER     = "postgres"
    DB_PASSWORD = "1234"
    DB_NAME     = "movies"
)


func SetupDB() *sql.DB {
    dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
    db, err := sql.Open("postgres", dbinfo)
	
    H.CheckErr(err)
    return db
}
