package controller

import (
	"clean-arsitektur/simple-restapi/models"
	"clean-arsitektur/simple-restapi/response"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var repo models.RepositoryTodo = models.NewRepositoryTodo()

type H map[string]interface{}

func InsertTodo(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	todo := models.Todo{}
	err = json.Unmarshal(body, &todo)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	todo.Prepare()
	err = todo.Validate()
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	_, err = repo.Insert(&todo)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	response.JSON(w, http.StatusCreated, H{"status": true, "msg": "created", "data": todo})

}

func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	todo := models.Todo{}
	allPosts, err := repo.GetAllTodo(&todo)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, H{"status": true, "data": allPosts})
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
	}

	todo := models.Todo{}
	getPost, err := repo.GetTodo(tid, &todo)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, H{"status": true, "data": getPost})
}

func UpdateTodos(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	todo := models.Todo{}
	err = json.Unmarshal(body, &todo)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = repo.UpdateTodo(tid, &todo, w)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, H{"status": true, "msg": "updated", "data": todo})
}

func DeleteTodos(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = repo.DeleteTodo(tid, w)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, H{"status": true, "msg": "deleted"})
}
