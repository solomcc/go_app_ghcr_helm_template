// main.go
package main

import (
	"net/http"
	"embed"
)
//go:embed static/index.html
var content embed.FS



func handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFileFS(w, r, content, "/static/index.html")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8083", nil)
}