package main

import (
	"errors"
	"fmt"
	"net/http"
)

func (c *Config) Authenticate(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := c.readJSON(w, r, &requestPayload)
	if err != nil {
		c.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// validate the user against the database
	foundUser, err := c.Models.User.GetByEmail(requestPayload.Email)
	if err != nil {
		c.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	valid, err := foundUser.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		c.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Logged in user %s", foundUser.Email),
		Data:    foundUser,
	}
	c.writeJSON(w, http.StatusAccepted, payload)
}
