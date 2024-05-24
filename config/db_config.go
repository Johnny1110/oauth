package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:Jarvan1110@tcp(127.0.0.1:3306)/basemall?parseTime=true&charset=utf8mb4,utf8")
	if err != nil {
		log.Fatalf("Error opening database: %v\n", err)
	}
	db.SetMaxOpenConns(10) // max connection
	db.SetMaxIdleConns(10) // max idle connection
	db.SetConnMaxLifetime(time.Minute * 3)
	if err = db.Ping(); err != nil {
		log.Fatalf("Error connecting to database: %v\n", err)
	}
}

// GetDB returns the database connection.
func GetDB() *sql.DB {
	returnObj := db
	//log.Println("[GetDB] connection address: ", &returnObj)
	return returnObj
}
