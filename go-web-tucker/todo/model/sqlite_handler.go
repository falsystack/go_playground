package model

import (
	"database/sql"
	"time"
)
import _ "github.com/mattn/go-sqlite3"

type sqliteHandler struct {
	db *sql.DB
}

func newSqliteHandler() DBHandler {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare(
		`CREATE TABLE IF NOT EXISTS todos (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			completed BOOLEAN,
			createdAt DATETIME)
    `)
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec()
	if err != nil {
		panic(err)
	}
	return &sqliteHandler{db: db}
}

func (s *sqliteHandler) Close() {
	s.db.Close()
}

func (s *sqliteHandler) GetTodos() []*Todo {
	var todos []*Todo
	rows, err := s.db.Query("SELECT * FROM todos")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.ID, &todo.Name, &todo.Completed, &todo.CreatedAt)
		if err != nil {
			panic(err)
		}
		defer rows.Close()
		todos = append(todos, &todo)
	}
	return todos
}

func (s *sqliteHandler) AddTodo(name string) *Todo {
	stmt, err := s.db.Prepare("INSERT INTO todos (name, completed, createdAt) VALUES (?,?,datetime('now'))")
	if err != nil {
		panic(err)
	}
	result, err := stmt.Exec(name, false)
	if err != nil {
		panic(err)
	}
	id, _ := result.LastInsertId()
	var todo Todo
	todo.ID = int(id)
	todo.Name = name
	todo.Completed = false
	todo.CreatedAt = time.Now()

	return &todo
}

func (s *sqliteHandler) RemoveTodo(id int) bool {
	stmt, err := s.db.Prepare("DELETE FROM todos WHERE id=?")
	if err != nil {
		panic(err)
	}

	result, err := stmt.Exec(id)
	if err != nil {
		panic(err)
	}

	if cnt, _ := result.RowsAffected(); cnt > 0 {
		return true
	}
	return false
}

func (s *sqliteHandler) CompleteTodo(id int, complete bool) bool {
	stmt, err := s.db.Prepare("UPDATE todos SET completed=? WHERE id=?")
	if err != nil {
		panic(err)
	}

	result, err := stmt.Exec(complete, id)
	if err != nil {
		panic(err)
	}

	if cnt, _ := result.RowsAffected(); cnt > 0 {
		return true
	}
	return false
}
