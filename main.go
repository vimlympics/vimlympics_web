package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/vimlympics/vimlympics_web/db"
	// "htmxx/service"

	"github.com/alexedwards/scs/v2"
	"github.com/joho/godotenv"
	_ "github.com/tursodatabase/go-libsql"
	"golang.org/x/oauth2"
	githubOAuth2 "golang.org/x/oauth2/github"
)

type application struct {
	config  config
	db      *sql.DB
	query   *db.Queries
	oauth   *oauth2.Config
	session *scs.SessionManager
}

type config struct {
	httpPort string
	db       struct {
		dsn string
	}
}

func main() {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			panic("Failed to load .env file")
		}
	}
	var config config

	config.httpPort = GetStringEnv("PORT", "8081")
	config.db.dsn = GetStringEnv("SQLITE_PATH", "")

	dbConn, err := NewDB(config.db.dsn)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}
	defer dbConn.Close()

	sessionManager := scs.New()
	sessionManager.Lifetime = 24 * time.Hour

	app := &application{
		config: config,
		db:     dbConn,
		query:  db.New(dbConn),
		oauth: &oauth2.Config{
			// TODO: Move this
			ClientID:    GetStringEnv("GH_CLIENT_ID", ""),
			ClientSecret: GetStringEnv("GH_CLIENT_SECRET", ""),
			RedirectURL:  GetStringEnv("GGH_CALLBACK", ""),
			Endpoint:     githubOAuth2.Endpoint,
		},
		session: sessionManager,
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
