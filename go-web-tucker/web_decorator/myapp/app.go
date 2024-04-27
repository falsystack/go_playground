package myapp

import (
	"fmt"
	"go-web-tucker/web_decorator/decoHandler"
	"log"
	"net/http"
	"time"
)

func NewHandler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	h := decoHandler.NewDecoHandler(mux, logger)

	return h
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func logger(w http.ResponseWriter, req *http.Request, h http.Handler) {
	startTime := time.Now()
	log.Print("[LOGGER1] Started")
	h.ServeHTTP(w, req)
	log.Print("[LOGGER1] Completed time", time.Since(startTime).Milliseconds())
}
