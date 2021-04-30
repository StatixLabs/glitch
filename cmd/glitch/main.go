package main

import (
	"fmt"
	"os"

	"github.com/jakealves/glitch/internal/database"
	"github.com/jakealves/glitch/internal/file"
)

// Run SQL query against SQL Server database
func RunSqlQueryAgainstSqlServer(query string) {
	sqlServer := database.SqlServer{}
	dsn, err := sqlServer.GetDSN()
	if err != nil {
		fmt.Print(err, dsn)
	}

	rows, err := sqlServer.Query(dsn, query)
	if err != nil {
		fmt.Print(err, dsn)
	}

	jsonResults, err := database.ConvertDatabaseRowsToJSON(rows)
	if err != nil {
		fmt.Print(err)
	}

	err = file.WriteContentsToFile("sql_server.json", jsonResults)
	if err != nil {
		fmt.Print(err)
	}
}

// Run SQL query against Snowflake database
func RunSqlQueryAgainstSnowflake(query string) {
	snowflake := database.Snowflake{}
	dsn, err := snowflake.GetDSN()
	if err != nil {
		fmt.Print(err, dsn)
	}

	rows, err := snowflake.Query(dsn, query)
	if err != nil {
		fmt.Print(err)
	}

	jsonResults, err := database.ConvertDatabaseRowsToJSON(rows)
	if err != nil {
		fmt.Print(err)
	}

	err = file.WriteContentsToFile("snowflake.json", jsonResults)
	if err != nil {
		fmt.Print(err)
	}
}

func main() {
	tsql, err := file.ReadContentsFromFile(os.Args[1])
	if err != nil {
		fmt.Print(err)
	}
	RunSqlQueryAgainstSqlServer(tsql)
	RunSqlQueryAgainstSnowflake(tsql)
}
