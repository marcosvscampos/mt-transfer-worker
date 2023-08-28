package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	Instance *sql.DB
}

func (db *DB) Connect() {

	dsn := buildDSN()
	fmt.Println("Connecting into", dsn)
	db.Instance = openDBConnection(dsn)
	testConnection(db.Instance)
}

func (db *DB) CloseConnection() {
	fmt.Println("Closing connection...")
	db.Instance.Close()
}

func buildDSN() string {
	var (
		username string
		password string
		host     string
		port     string
		dbName   string
	)
	username = "roguesoft"
	password = "develop@123!"
	host = "localhost"
	port = "3306"
	dbName = "db-mt-transfers-service"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbName)
	return dsn
}

func openDBConnection(dsn string) *sql.DB {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	return db
}

func testConnection(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection successful")
}
