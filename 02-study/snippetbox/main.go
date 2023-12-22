/**
 * file: main.go
 * description: file responsible for running the application.
 * data: 22/12/2023
 * author: Glaucia Lemos <Twitter: @glaucia_lemos86>
 */

package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Print("Starting the server on port...: 4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
