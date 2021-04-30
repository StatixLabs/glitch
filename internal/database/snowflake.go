package database

import (
	"database/sql"
	"log"
	"os"
	"strconv"

	"github.com/snowflakedb/gosnowflake"
)

type Snowflake struct{}

// GetDSN construct data source name for Snowflake database
func (s Snowflake) GetDSN() (string, error) {
	env := func(k string, failOnMissing bool) string {
		if value := os.Getenv(k); value != "" {
			return value
		}
		if failOnMissing {
			log.Fatalf("%v environment variable is not set.", k)
		}
		return ""
	}

	account := env("SNOWFLAKE_ACCOUNT", true)
	user := env("SNOWFLAKE_USER", true)
	password := env("SNOWFLAKE_PASSWORD", true)
	host := env("SNOWFLAKE_HOST", false)
	port := env("SNOWFLAKE_PORT", false)
	protocol := env("SNOWFLAKE_PROTOCOL", false)
	role := env("SNOWFLAKE_ROLE", false)
	warehouse := env("SNOWFLAKE_WAREHOUSE", false)

	portStr, _ := strconv.Atoi(port)
	cfg := &gosnowflake.Config{
		Account:   account,
		User:      user,
		Password:  password,
		Host:      host,
		Port:      portStr,
		Protocol:  protocol,
		Role:      role,
		Warehouse: warehouse,
	}

	dsn, err := gosnowflake.DSN(cfg)
	return dsn, err
}

// Query Snowflake database
func (s Snowflake) Query(dsn string, query string) (*sql.Rows, error) {
	db, err := sql.Open("snowflake", dsn)
	if err != nil {
		log.Fatalf("failed to connect. %v, err: %v", dsn, err)
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query(query) // no cancel is allowed
	if err != nil {
		log.Fatalf("failed to run a query. %v, err: %v", query, err)
		return nil, err
	}
	return rows, nil
}
