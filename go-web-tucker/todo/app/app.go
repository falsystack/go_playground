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
var rd *render.Render

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

func MakeHandler() http.Handler {
	//todoMap = make(map[int]*model.Todo)
	//addTestTodos()

	rd = render.New()
	r := mux.NewRouter()

	r.HandleFunc("/complete-todo/{id:[0-9]+}", completeTodoHanlder).Methods("GET")
	r.HandleFunc("/todos/{id:[0-9]+}", removeTodoHandler).Methods("DELETE")
	r.HandleFunc("/todos", addTodoHandler).Methods("POST")
	r.HandleFunc("/todos", getTodoListHandler).Methods("GET")
	r.HandleFunc("/", indexHandler)
	return r
}

func completeTodoHanlder(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])
	complete := req.FormValue("completed") == "true"
	ok := model.CompleteTodo(id, complete)

	if ok {
		rd.JSON(w, http.StatusOK, Success{Success: true})
		return
	}
	rd.JSON(w, http.StatusBadRequest, Success{Success: false})
}

func removeTodoHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])
	ok := model.RemoveTodo(id)
	if ok {
		rd.JSON(w, http.StatusOK, Success{Success: true})
		return
	}
	rd.JSON(w, http.StatusBadRequest, Success{Success: false})
}

func addTodoHandler(w http.ResponseWriter, req *http.Request) {
	name := req.FormValue("name")
	todo := model.AddTodo(name)

	rd.JSON(w, http.StatusCreated, todo)
}

func getTodoListHandler(w http.ResponseWriter, req *http.Request) {
	list := model.GetTodos()
	rd.JSON(w, http.StatusOK, list)
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "/todo.html", http.StatusTemporaryRedirect)
}
