package models

import (
	"clean-arsitektur/simple-restapi/db"
	"clean-arsitektur/simple-restapi/response"
	"errors"
	"fmt"
	"html"
	"log"
	"net/http"
	"strings"
)

type Todo struct {
	ID          uint64 `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type RepositoryTodo interface {
	Insert(todo *Todo) (*Todo, error)
	GetAllTodo(todo *Todo) ([]Todo, error)
	GetTodo(tid uint64, todo *Todo) ([]Todo, error)
	UpdateTodo(tid uint64, todo *Todo, w http.ResponseWriter) error
	DeleteTodo(tid uint64, w http.ResponseWriter) error
}

type Repo struct{}

func NewRepositoryTodo() RepositoryTodo {
	return &Repo{}
}

func (t *Todo) Prepare() {
	t.ID = 0
	t.Title = html.EscapeString(strings.TrimSpace(t.Title))
	t.Description = html.EscapeString(strings.TrimSpace(t.Description))
}

func (t *Todo) Validate() error {
	if t.Title == "" {
		return errors.New("Required Title")
	}

	if t.Description == "" {
		return errors.New("Required Description")
	}

	return nil
}

func (repo *Repo) Insert(todo *Todo) (*Todo, error) {
	row, err := db.DB.Query(db.InsertTodoQuery, todo.Title, todo.Description)
	if err != nil {
		panic(err)
	}

	defer row.Close()

	return todo, nil
}

func (repo *Repo) GetAllTodo(todo *Todo) ([]Todo, error) {
	row, err := db.DB.Query(db.GetAllTodoQuery)
	if err != nil {
		panic(err)
	}

	defer row.Close()

	res := []Todo{}

	for row.Next() {
		todo := Todo{}
		err = row.Scan(&todo.ID, &todo.Title, &todo.Description)
		if err != nil {
			fmt.Printf(err.Error())
		}

		res = append(res, todo)
	}

	if err = row.Err(); err != nil {
		panic(err)
	}

	return res, err
}

func (repo *Repo) GetTodo(tid uint64, todo *Todo) ([]Todo, error) {
	row, err := db.DB.Query(db.GetTodoQuery, tid)
	if err != nil {
		panic(err)
	}

	defer row.Close()

	res := []Todo{}

	for row.Next() {
		todo := Todo{}
		err = row.Scan(&todo.ID, &todo.Title, &todo.Description)
		if err != nil {
			fmt.Printf(err.Error())
		}

		res = append(res, todo)
	}

	if err = row.Err(); err != nil {
		panic(err)
	}

	return res, err
}

func (repo *Repo) UpdateTodo(tid uint64, todo *Todo, w http.ResponseWriter) error {
	row, err := db.DB.Exec(db.UpdateTodoQuery, todo.Title, todo.Description, tid)
	if err != nil {
		panic(err)
	}

	rowsAffected, err := row.RowsAffected()
	if err != nil {
		log.Fatal(err.Error())
	}

	if rowsAffected == 0 {
		response.ERROR(w, http.StatusNotFound, errors.New("Todo Not Found!"))
	}

	if rowsAffected > 1 {
		response.ERROR(w, http.StatusFound, errors.New("Total Affected"+string(rowsAffected)))
	}

	return nil
}

func (repo *Repo) DeleteTodo(tid uint64, w http.ResponseWriter) error {
	row, err := db.DB.Exec(db.DeleteQuery, tid)
	if err != nil {
		panic(err)
	}

	rowsAffected, err := row.RowsAffected()
	if err != nil {
		log.Fatal(err.Error())
	}

	if rowsAffected == 0 {
		response.ERROR(w, http.StatusNotFound, errors.New("Todo Not Found!"))
	}

	if rowsAffected > 1 {
		response.ERROR(w, http.StatusFound, errors.New("Total Affected"+string(rowsAffected)))
	}

	return nil
}
