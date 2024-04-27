package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/hot-oranges-86/Todo/models"
)

var db *sql.DB

func ConnectToPostgreSQL() error {
	var err error
	connStr := "dbname=todo password=root sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	if err = db.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	return nil
}

func CreateTodo(todo models.Todo) (models.Todo, error) {
	query := `INSERT INTO todos (name, completed) VALUES ($1, $2) RETURNING id`
	err := db.QueryRow(query, todo.Name, todo.Completed).Scan(&todo.ID)
	if err != nil {
		return models.Todo{}, err
	}
	return todo, nil
}
