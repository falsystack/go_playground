package main

import (
	"github.com/google/uuid"
	"net/http"
)

func getUser(w http.ResponseWriter, req *http.Request) user {
	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		id := uuid.New()
		c = &http.Cookie{
			Name:  "session",
			Value: id.String(),
		}
	}
	http.SetCookie(w, c)

	// if the user exists
	var u user
	if un, ok := dbSessions[c.Value]; ok {
		u = dbUsers[un]
	}
	return u
}

func alreadyLoggedIn(req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	un := dbSessions[c.Value]
	_, ok := dbUsers[un]
	return ok
}
