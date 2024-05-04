package app

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
	"go-web-tucker/todo/model"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Success struct {
	Success bool `json:"success"`
}

// for marshal
var rd = render.New()

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

type AppHandler struct {
	http.Handler
	db model.DBHandler
}

func CheckSignin(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if strings.Contains(r.URL.Path, "/signin.html") ||
		strings.Contains(r.URL.Path, "/auth") {
		next(w, r)
		return
	}

	sessionID := getSessionID(r)
	if sessionID != "" {
		next(w, r)
		return
	}
	http.Redirect(w, r, "/signin.html", http.StatusTemporaryRedirect)
}

func MakeHandler() *AppHandler {
	r := mux.NewRouter()
	n := negroni.New(negroni.NewRecovery(), negroni.NewLogger(), negroni.HandlerFunc(CheckSignin), negroni.NewStatic(http.Dir("public")))
	n.UseHandler(r)
	a := &AppHandler{
		Handler: n,
		db:      model.NewDBHandler(),
	}

	r.HandleFunc("/complete-todo/{id:[0-9]+}", a.completeTodoHanlder).Methods("GET")
	r.HandleFunc("/todos/{id:[0-9]+}", a.removeTodoHandler).Methods("DELETE")
	r.HandleFunc("/todos", a.addTodoHandler).Methods("POST")
	r.HandleFunc("/todos", a.getTodoListHandler).Methods("GET")
	r.HandleFunc("/", a.indexHandler)
	r.HandleFunc("/auth/google/login", googleLoginHandler)
	r.HandleFunc("/auth/google/callback", googleAtuCallback)
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

func getSessionID(req *http.Request) string {
	session, err := store.Get(req, "session")
	if err != nil {
		return ""
	}

	val := session.Values["id"]
	if val == nil {
		return ""
	}
	return val.(string)
}
