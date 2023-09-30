# Golang todo app
Simple todo app made using go for the backend and vanilla js for the frontend

## How to run
Set up the postgresql database server. The database connection secrets are hard coded in `db.go` as this is just an offline demo project, not meant for hosting on a public server. 
Clone the repository and run `go run cmd/todo_app/api_server.go` in project root
The app should be accessible via web browser on `localhost:8080`

## Features:
- Adding todos one at a time
- Deleting selected todo
- Deleting all todos at once
- Editing todo state by switching the checkbox next to it
- Editing todo title by double clicking it, renaming and confirming with `enter`
