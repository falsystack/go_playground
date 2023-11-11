package main

import (
	"io"
	"net/http"
)

/*
ServeFile は指定されたファイルまたはディレクトリの内容でリクエストに応答します。
指定されたファイル名またはディレクトリ名が相対パスの場合は、
現在のディレクトリを基準として解釈され、上に上がる可能性があります。
*/

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/toby.jpg", dogPic)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="toby.jpg />"`)
}

func dogPic(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "toby.jpg")
}
