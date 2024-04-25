package myapp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type User struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type fooHandler struct{}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	u := new(User)
	err := json.NewDecoder(req.Body).Decode(u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	u.CreatedAt = time.Now()

	// struct -> json([]byte)
	data, _ := json.Marshal(u)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusCreated) // 順番大事、w.Header().Add()が後にくるとAddされない
	fmt.Fprint(w, string(data))
}

func barHadnler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name") // query parameter
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello %s", name)
}

func NewHttpHandler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/bar", barHadnler)

	mux.Handle("/foo", &fooHandler{})

	return mux
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
