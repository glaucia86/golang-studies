/**
 * file: main.go
 * description: file responsible for running the application.
 * data: 12/29/2023
 * author: Glaucia Lemos <Twitter: @glaucia_lemos86>
 */

package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":3000", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	f, _ := os.Open("./menu.txt")
	io.Copy(w, f)
}
