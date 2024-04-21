package main

import (
	"fmt"
	"net/http"
)

type Beef int

func (b Beef) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("関数の名前がServeHTTPでそのパラメーターがServeHTTPのパラメーターと一致すればHandlerと考える")
}

func main() {}
