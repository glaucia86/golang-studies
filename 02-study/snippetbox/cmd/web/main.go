/**
 * file: cmd/web/main.go
 * description: file responsible for running the application.
 * data: 12/27/2023
 * author: Glaucia Lemos <Twitter: @glaucia_lemos86>
 */

package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")

	flag.Parse()

	mux := http.NewServeMux()

	// Configuração do FileServer para servir arquivos estáticos.
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Demais rotas da aplicação
	//mux.HandleFunc("/", home)
	mux.Handle("/", http.HandlerFunc(home))
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Printf("Starting the server on %s", *addr)

	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
