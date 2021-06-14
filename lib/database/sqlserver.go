package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
)

type SqlServer struct{}

// GetDSN construct data source name for sql server database
func (s SqlServer) GetDSN() (string, error) {

	env := func(k string, failOnMissing bool) string {
		if value := os.Getenv(k); value != "" {
			return value
		}
		if failOnMissing {
			log.Fatalf("%v environment variable is not set.", k)
		}
		return ""
	}

	server := env("SQLSERVER_HOST", true)
	port := env("SQLSERVER_PORT", true)
	user := env("SQLSERVER_USER", true)
	password := env("SQLSERVER_PASSWORD", true)
	database := env("SQLSERVER_DATABASE", false)

	// Build connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;",
		server, user, password, port, database)
	return connString, nil
}

// Query sql server database
func (s SqlServer) Query(dsn string, query string) (*sql.Rows, error) {
	db, err := sql.Open("sqlserver", dsn)
	if err != nil {
		log.Fatalf("failed to connect. %v, err: %v", dsn, err)
		return nil, err
	}
	defer db.Close()
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	rows, err := db.QueryContext(ctx, query)
	return rows, err
}
