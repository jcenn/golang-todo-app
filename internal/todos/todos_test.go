package todos

import (
	"database/sql"
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

var database *sql.DB

func run(m *testing.M) (code int, err error) {
	// pseudo-code, some implementation excluded:
	//
	// 1. create test.db if it does not exist
	// 2. run our DDL statements to create the required tables if they do not exist
	// 3. run our tests
	// 4. truncate the test db tables

	database = db.ConnectToDB()
	database.Exec(`CREATE TABLE "todos_test" AS (SELECT * FROM "todos") WITH NO DATA`)
	// truncates all test data after the tests are run
	defer func() {
		database.Query(`DROP TABLE "todos_test"`)
		database.Close()
	}()

	return m.Run(), nil
}

func TestAddTodo(t *testing.T) {
	var initTodoCount uint

	countRow := database.QueryRow(`SELECT COUNT(*) FROM "todos_test"`)
	countRow.Scan(&initTodoCount)

	query := fmt.Sprintf(`INSERT INTO todos_test(name, is_finished) VALUES('%s', false) RETURNING id, name, is_finished`, "test_name")
	_, err := database.Exec(query)
	if err != nil {
		t.Fatal(err)
	}
	var newTodoCount uint
	countRow = database.QueryRow(`SELECT COUNT(*) FROM "todos_test"`)
	countRow.Scan(&newTodoCount)
	if newTodoCount != initTodoCount+1 {
		t.Logf("records before operation: %d, after: %d, expected: %d", initTodoCount, newTodoCount, initTodoCount+1)
		t.Fatal()
	}
}
