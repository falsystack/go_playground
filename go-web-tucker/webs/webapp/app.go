package webapp

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func NewHandler() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/users", userHandler).Methods("GET")
	r.HandleFunc("/users", createUserHandler).Methods("POST")
	r.HandleFunc("/users/{id:[0-9]+}", getUserHandler)

	return r
}

func createUserHandler(w http.ResponseWriter, req *http.Request) {
	u := new(User)
	err := json.NewDecoder(req.Body).Decode(u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	// create user
	u.ID = 2
	u.CreatedAt = time.Now()
	w.WriteHeader(http.StatusCreated)
	data, _ := json.Marshal(u)
	fmt.Fprint(w, string(data))
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func userHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Get UserInfo by /users")
}

func getUserHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	fmt.Fprintf(w, "User Id:%s", vars["id"])
}

func getUserInfoHandler(w http.ResponseWriter, req *http.Request) {

}
