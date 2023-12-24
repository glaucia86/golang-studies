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

	// Configuração do FileServer para servir arquivos estáticos.
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Demais rotas da aplicação
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("Starting the server on port...: 4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
