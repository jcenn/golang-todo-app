package db_test

import (
	"fmt"
	"os"
	"testing"
	"todo-list/internal/db"
)

func TestMain(m *testing.M) {
    code, err := run(m)
    if err != nil {
        fmt.Println(err)
    }
    os.Exit(code)
}

func run(m *testing.M) (code int, err error) {
    // pseudo-code, some implementation excluded:
    //
    // 1. create test.db if it does not exist
    // 2. run our DDL statements to create the required tables if they do not exist
    // 3. run our tests
    // 4. truncate the test db tables

    db :=  db.ConnectToDB()
    db.Exec(`CREATE TABLE "todos_test" AS (SELECT * FROM "todos") WITH NO DATA`)

    // truncates all test data after the tests are run
    defer func() {
        db.Query(`DROP TABLE "todos_test"`)
        db.Close()
    }()

    return m.Run(), nil
}