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
	connStr := "dbname=todo user=todoapi password=ob1uh8j sslmode=disable"
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

func GetTodos() ([]models.Todo, error) {
	query := `SELECT * FROM todos`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		var todo models.Todo
		err := rows.Scan(&todo.ID, &todo.Name, &todo.Completed)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func GetTodo(id int) (models.Todo, error) {
	query := `SELECT * FROM todos WHERE id = $1`
	var todo models.Todo
	err := db.QueryRow(query, id).Scan(&todo.ID, &todo.Name, &todo.Completed)
	if err != nil {
		return models.Todo{}, err
	}
	return todo, nil
}

func UpdateTodo(updatedTodo models.Todo) (models.Todo, error) {
	query := `UPDATE todos SET name = $1, completed = $2 WHERE id = $3`
	_, err := db.Exec(query, updatedTodo.Name, updatedTodo.Completed, updatedTodo.ID)
	if err != nil {
		return models.Todo{}, err
	}
	return updatedTodo, nil
}

func DeleteTodo(id int) error {
	query := `DELETE FROM todos WHERE id = $1`
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
