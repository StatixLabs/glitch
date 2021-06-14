package cmd

import (
	"fmt"

	"github.com/jakealves/glitch/lib/database"
	"github.com/jakealves/glitch/lib/file"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(sqlcompareCmd)
}

var sqlcompareCmd = &cobra.Command{
	Use:   "sqlcompare",
	Short: "one query against two types of databases.",
	Long:  `Will take in a file with a sql query and run it against sql server and snowflake.`,
	Run: func(cmd *cobra.Command, args []string) {
		runSqlCompareCmd(cmd, args)
	},
}

// Command to Run SQL Compare
func runSqlCompareCmd(cmd *cobra.Command, args []string) {
	tsql, err := file.ReadContentsFromFile(args[1])
	if err != nil {
		fmt.Print(err)
	}
	RunSqlQueryAgainstSqlServer(tsql)
	RunSqlQueryAgainstSnowflake(tsql)
}

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
