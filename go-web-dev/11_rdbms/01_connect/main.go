package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"net/http"
)

func main() {
	db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/udemy?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}

	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func index(w http.ResponseWriter, req *http.Request) {
	_, err := io.WriteString(w, "Successfully completed")
	if err != nil {
		fmt.Println(err)
	}
}
