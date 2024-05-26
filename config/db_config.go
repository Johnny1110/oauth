package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"oauth/sys"
	"time"
)

var db *sql.DB

func init() {
	properties := GetProperties()
	var err error
	// username:password@tcp(127.0.0.1:3306)/basemall?parseTime=true&charset=utf8mb4,utf8
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		properties.DB.Username,
		properties.DB.Password,
		properties.DB.IP,
		properties.DB.Port,
		properties.DB.ConnectionString)
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		sys.Logger().Errorf("Error opening database: %v\n", err)
	}
	db.SetMaxOpenConns(properties.DB.MaxConnection)     // max connection
	db.SetMaxIdleConns(properties.DB.MaxIdleConnection) // max idle connection
	db.SetConnMaxLifetime(time.Minute * time.Duration(properties.DB.MaxConnectionLifetime))
	if err = db.Ping(); err != nil {
		sys.Logger().Errorf("Error connecting to database: %v\n", err)
	}
}

// GetDB returns the database connection.
func GetDB() *sql.DB {
	returnObj := db
	return returnObj
}
