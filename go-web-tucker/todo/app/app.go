package app

import (
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"net/http"
	"strconv"
	"time"
)

type Todo struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

type Success struct {
	Success bool `json:"success"`
}

var todoMap map[int]*Todo

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
	todoMap = make(map[int]*Todo)
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
	if todo, ok := todoMap[id]; ok {
		todo.Completed = complete
		rd.JSON(w, http.StatusOK, Success{true})
		return
	}
	rd.JSON(w, http.StatusBadRequest, Success{false})
}

func removeTodoHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])
	if _, ok := todoMap[id]; ok {
		delete(todoMap, id)
		rd.JSON(w, http.StatusOK, Success{Success: true})
		return
	}
	rd.JSON(w, http.StatusBadRequest, Success{Success: false})
}

func addTodoHandler(w http.ResponseWriter, req *http.Request) {
	id := len(todoMap) + 1
	newTodo := &Todo{
		ID:        id,
		Name:      req.FormValue("name"),
		Completed: false,
		CreatedAt: time.Now(),
	}
	todoMap[id] = newTodo

	rd.JSON(w, http.StatusCreated, newTodo)
}

func getTodoListHandler(w http.ResponseWriter, req *http.Request) {
	list := []*Todo{}
	for _, v := range todoMap {
		list = append(list, v)
	}
	rd.JSON(w, http.StatusOK, list)
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "/todo.html", http.StatusTemporaryRedirect)
}
