package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/vimlympics/vimlympics_web/db"
	// "htmxx/service"

	_ "github.com/tursodatabase/go-libsql"
)

type application struct {
	config config
	db     *sql.DB
	query  *db.Queries
}

type config struct {
	httpPort string
	db       struct {
		dsn string
	}
}

func main() {
	var config config

	config.httpPort = GetStringEnv("PORT", "8081")
	config.db.dsn = GetStringEnv("HTMXX_DB_URL", "file:./db.sqlite3")

	dbConn, err := NewDB(config.db.dsn)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}
	defer dbConn.Close()

	app := &application{
		config: config,
		db:     dbConn,
		query:  db.New(dbConn),
	}
	//
	err = app.serveHTTP()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func GetStringEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func NewDB(dsn string) (*sql.DB, error) {
	dbConn, err := sql.Open("libsql", dsn)
	if err != nil {
		return nil, err
	}
	return dbConn, err
}
