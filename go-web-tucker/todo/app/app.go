package app

import (
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"go-web-tucker/todo/model"
	"net/http"
	"strconv"
)

type Success struct {
	Success bool `json:"success"`
}

// for marshal
var rd = render.New()

//func addTestTodos() {
//	todoMap[1] = &Todo{
//		ID:        1,
//		Name:      "Buy a milk",
//		Completed: false,
//		CreatedAt: time.Now(),
//	}
//	todoMap[2] = &Todo{
//		ID:        2,
//		Name:      "Exercise",
//		Completed: true,
//		CreatedAt: time.Now(),
//	}
//	todoMap[3] = &Todo{
//		ID:        3,
//		Name:      "Home work",
//		Completed: false,
//		CreatedAt: time.Now(),
//	}
//}

type AppHandler struct {
	http.Handler
	db model.DBHandler
}

func MakeHandler() *AppHandler {
	r := mux.NewRouter()
	a := &AppHandler{
		Handler: r,
		db:      model.NewDBHandler(),
	}

	r.HandleFunc("/complete-todo/{id:[0-9]+}", a.completeTodoHanlder).Methods("GET")
	r.HandleFunc("/todos/{id:[0-9]+}", a.removeTodoHandler).Methods("DELETE")
	r.HandleFunc("/todos", a.addTodoHandler).Methods("POST")
	r.HandleFunc("/todos", a.getTodoListHandler).Methods("GET")
	r.HandleFunc("/", a.indexHandler)
	return a
}

func (a *AppHandler) Close() {
	a.db.Close()
}

func (a *AppHandler) completeTodoHanlder(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])
	complete := req.FormValue("completed") == "true"
	ok := a.db.CompleteTodo(id, complete)

	if ok {
		rd.JSON(w, http.StatusOK, Success{Success: true})
		return
	}
	rd.JSON(w, http.StatusBadRequest, Success{Success: false})
}

func (a *AppHandler) removeTodoHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])
	ok := a.db.RemoveTodo(id)
	if ok {
		rd.JSON(w, http.StatusOK, Success{Success: true})
		return
	}
	rd.JSON(w, http.StatusBadRequest, Success{Success: false})
}

func (a *AppHandler) addTodoHandler(w http.ResponseWriter, req *http.Request) {
	name := req.FormValue("name")
	todo := a.db.AddTodo(name)

	rd.JSON(w, http.StatusCreated, todo)
}

func (a *AppHandler) getTodoListHandler(w http.ResponseWriter, req *http.Request) {
	list := a.db.GetTodos()
	rd.JSON(w, http.StatusOK, list)
}

func (a *AppHandler) indexHandler(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "/todo.html", http.StatusTemporaryRedirect)
}
