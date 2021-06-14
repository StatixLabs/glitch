package database_test

import (
	"database/sql"
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jakealves/glitch/lib/database"
)

var _ = Describe("When ConvertDatabaseRowsToJSON function is run", func() {
	Context("given *db.rows returned from a database", func() {
		var rows *sql.Rows
		BeforeEach(func() {
			db, mock, err := sqlmock.New()
			if err != nil {
				fmt.Printf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()
			mockedRows := sqlmock.NewRows([]string{"id", "title", "body"}).
				AddRow(1, true, "hello").
				AddRow(2, false, nil)
			mock.ExpectQuery("^SELECT (.+) FROM posts$").WillReturnRows(mockedRows)
			rows, err = db.Query("SELECT * FROM posts")
			if err != nil {
				fmt.Print(err)
			}
		})

		It("should convert the rows to json", func() {
			convertedRows, err := database.ConvertDatabaseRowsToJSON(rows)
			if err != nil {
				fmt.Print(err)
			}
			expectedJSON := `[
 {
  "body": "hello",
  "id": "1",
  "title": "true"
 },
 {
  "body": "",
  "id": "2",
  "title": "false"
 }
]`
			Expect(string(convertedRows)).To(BeEquivalentTo(expectedJSON))
		})
	})
})
