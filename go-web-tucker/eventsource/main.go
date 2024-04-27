package main

import (
	"encoding/json"
	"fmt"
	"github.com/antage/eventsource"
	"github.com/gorilla/pat"
	"github.com/urfave/negroni"
	"net/http"
	"strconv"
	"time"
)

func main() {
	msgCh = make(chan Message)

	es := eventsource.New(nil, nil)
	defer es.Close()

	go processMsgCh(es)

	mux := pat.New()
	mux.Post("/messages", postMessageHandler)
	mux.Handle("/stream", es)
	mux.Post("/users", addUserHandler)

	n := negroni.Classic()
	n.UseHandler(mux)

	http.ListenAndServe(":8080", n)
}

func addUserHandler(w http.ResponseWriter, req *http.Request) {
	username := req.FormValue("name")
	sendMessage("", fmt.Sprintf("add user: %s", username))
}

func postMessageHandler(w http.ResponseWriter, req *http.Request) {
	msg := req.FormValue("msg")
	name := req.FormValue("name")
	sendMessage(name, msg)
}

type Message struct {
	Name string `json:"name"`
	Msg  string `json:"msg"`
}

var msgCh chan Message

func sendMessage(name string, msg string) {
	msgCh <- Message{name, msg}
}

func processMsgCh(es eventsource.EventSource) {
	for msg := range msgCh {
		data, _ := json.Marshal(msg)
		es.SendEventMessage(string(data), "", strconv.Itoa(time.Now().Nanosecond()))
	}
}
