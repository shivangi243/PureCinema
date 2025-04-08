package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Conn *sql.DB

func Connect() {
	var err error

	// Replace with your actual credentials
	dsn := "root:shivangi@12*@tcp(127.0.0.1:3306)/cinema"
	Conn, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	if err = Conn.Ping(); err != nil {
		log.Fatal("Database unreachable:", err)
	}

	fmt.Println("✅ Connected to MySQL database")
}
