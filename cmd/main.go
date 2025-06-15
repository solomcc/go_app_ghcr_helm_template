// main.go
package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from Go in MicroK8s!")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8083", nil)
}