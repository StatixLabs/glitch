package database

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/jakealves/glitch/internal/utils"
)

// ConvertDatabaseRowsToJSON convert db.rows to json
func ConvertDatabaseRowsToJSON(rows *sql.Rows) ([]byte, error) {
	defer rows.Close()
	columnTypes, err := rows.ColumnTypes()

	if err != nil {
		fmt.Print(err)
	}

	count := len(columnTypes)
	finalRows := []interface{}{}

	for rows.Next() {
		scanArgs := make([]interface{}, count)
		for i, v := range columnTypes {
			switch v.DatabaseTypeName() {
			case "VARCHAR", "TEXT", "UUID", "TIMESTAMP":
				scanArgs[i] = new(sql.NullString)
			case "BOOL":
				scanArgs[i] = new(sql.NullBool)
			case "INT4":
				scanArgs[i] = new(sql.NullInt64)
			default:
				scanArgs[i] = new(sql.NullString)
			}
		}
		err := rows.Scan(scanArgs...)

		if err != nil {
			return nil, err
		}

		masterData := map[string]interface{}{}
		for i, v := range columnTypes {
			if z, ok := (scanArgs[i]).(*sql.NullBool); ok {
				masterData[v.Name()] = z.Bool
				continue
			}

			if z, ok := (scanArgs[i]).(*sql.NullString); ok {
				masterData[v.Name()] = z.String
				continue
			}

			if z, ok := (scanArgs[i]).(*sql.NullInt64); ok {
				masterData[v.Name()] = z.Int64
				continue
			}

			if z, ok := (scanArgs[i]).(*sql.NullFloat64); ok {
				masterData[v.Name()] = z.Float64
				continue
			}

			if z, ok := (scanArgs[i]).(*sql.NullInt32); ok {
				masterData[v.Name()] = z.Int32
				continue
			}

			masterData[v.Name()] = scanArgs[i]
		}
		finalRows = append(finalRows, utils.ForceKeysToLowercase(masterData))
	}
	jsonByteArray, err := json.MarshalIndent(finalRows, "", " ")
	if err != nil {
		return nil, err
	}
	return jsonByteArray, nil
}

// Database interface to define interactions with databases
type Database interface {
	GetDSN() (string, error)
	Query(dsn string, query string) (*sql.Rows, error)
}
