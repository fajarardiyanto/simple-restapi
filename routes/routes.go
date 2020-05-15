package routes

import (
	"clean-arsitektur/simple-restapi/controller"
	"clean-arsitektur/simple-restapi/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

func Routes() {
	r := mux.NewRouter()

	r.HandleFunc("/insert", middlewares.SetMiddlewareJSON(controller.InsertTodo)).Methods("POST")
	r.HandleFunc("/todos", middlewares.SetMiddlewareJSON(controller.GetAllTodos)).Methods("GET")
	r.HandleFunc("/todos/{id}", middlewares.SetMiddlewareJSON(controller.GetTodos)).Methods("GET")
	r.HandleFunc("/update/todos/{id}", middlewares.SetMiddlewareJSON(controller.UpdateTodos)).Methods("PUT")
	r.HandleFunc("/delete/todos/{id}", middlewares.SetMiddlewareJSON(controller.DeleteTodos)).Methods("DELETE")

	http.ListenAndServe(":8000", r)
}
