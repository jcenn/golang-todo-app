package todos

type Todo struct {
	Id         uint   `json:"id"`
	Name       string `json:"name"`
	IsFinished bool   `json:"isFinished"`
}

type AddTodoRequest struct {
	Name string `json:"name"`
}


type EditTodoRequest struct {
	Name string `json:"name"`
	IsFinished bool `json:"isFinished"`
}