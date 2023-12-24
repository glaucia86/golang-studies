/**
 * file: cmd/web/main.go
 * description: file responsible for running the application.
 * data: 12/24/2023
 * author: Glaucia Lemos <Twitter: @glaucia_lemos86>
 */

package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("Starting the server on port...: 4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
