package todos

import "database/sql"

type TodoService struct {
	repo *TodoRepo
}

func NewTodoService(db *sql.DB) *TodoService {
	return &TodoService{
		repo: NewTodoRepo(db),
	}
}

func (s TodoService) AddTodo(req AddTodoRequest) *Todo {
	todo, err := s.repo.AddTodo(req)
	if err != nil {
		panic(err)
	}
	return todo
}

func (s TodoService) GetTodos() []Todo {
	todos, err := s.repo.GetTodos()
	if err != nil {
		panic(err)
	}
	return todos
}

func (s TodoService) GetTodoById(id uint) *Todo {
	todo, err := s.repo.GetTodoById(id)
	if err != nil {
		panic(err)
	}
	return todo
}

func (s TodoService) EditTodo(id uint, req EditTodoRequest) *Todo {
	todo, err := s.repo.EditTodo(id, req)
	if err != nil {
		panic(err)
	}
	return todo
}

func (s TodoService) DeleteTodo(id uint) {
	err := s.repo.RemoveTodo(id)
	if err != nil {
		panic(err)
	}
}

func (s TodoService) DeleteAllTodos() {
	err := s.repo.RemoveAllTodos()
	if err != nil {
		panic(err)
	}
}
