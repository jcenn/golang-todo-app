package todos

import (
	"database/sql"
	"fmt"
)

type TodoRepo struct {
	db *sql.DB
}

func NewTodoRepo(database *sql.DB) *TodoRepo {
	return &TodoRepo{
		db: database,
	}
}

func (r TodoRepo) AddTodo(req AddTodoRequest) (*Todo, error) {
	query := fmt.Sprintf(`INSERT INTO todos(name, is_finished) VALUES('%s', false) RETURNING id, name, is_finished`, req.Name)
	row := r.db.QueryRow(query)
	var (
		id         uint
		name       string
		isFinished bool
	)
	err := row.Scan(&id, &name, &isFinished)
	if err != nil {
		return nil, err
	}
	return &Todo{
		Id:         id,
		Name:       name,
		IsFinished: isFinished,
	}, nil
}

func (r TodoRepo) GetTodos() ([]Todo, error) {
	rows, err := r.db.Query(`SELECT "id", "name", "is_finished" FROM "todos" ORDER BY "id"`)
	if err != nil {
		panic(err)
	}
	todos := []Todo{}
	defer rows.Close()

	for rows.Next() {
		var (
			id          uint
			name        string
			is_finished bool
		)
		err = rows.Scan(&id, &name, &is_finished)
		if err != nil {
			panic(err)
		}
		todos = append(todos, Todo{Id: id, Name: name, IsFinished: is_finished})

	}
	return todos, nil
}

func (r TodoRepo) GetTodoById(query_id uint) (*Todo, error) {
	query := fmt.Sprintf(`SELECT * FROM todos WHERE id = %d`, query_id)
	row := r.db.QueryRow(query)
	var (
		id          uint
		name        string
		is_finished bool
	)
	if err := row.Scan(&id, &name, &is_finished); err != nil {
		return nil, err
	}

	return &Todo{
		Id:         id,
		Name:       name,
		IsFinished: is_finished,
	}, nil
}

func (r TodoRepo) EditTodo(queryId uint, request EditTodoRequest) (*Todo, error) {
	// fmt.Printf("id: %d \n", queryId)
	query := fmt.Sprintf(`UPDATE todos SET name = '%s', is_finished = %t WHERE id = %d RETURNING id, name, is_finished`, request.Name, request.IsFinished, queryId)
	row := r.db.QueryRow(query)
	var (
		newId       uint
		newName     string
		newFinished bool
	)

	err := row.Scan(&newId, &newName, &newFinished)

	if err != nil {
		return nil, err
	}
	return &Todo{
		Id:         newId,
		Name:       newName,
		IsFinished: newFinished,
	}, nil
}

func (r TodoRepo) RemoveTodo(query_id uint) error {
	query := fmt.Sprintf(`DELETE FROM todos WHERE id = %d`, query_id)
	_, err := r.db.Exec(query)
	return err
}

func (r TodoRepo) RemoveAllTodos() error {
	query := `TRUNCATE todos`
	_, err := r.db.Exec(query)
	return err
}
