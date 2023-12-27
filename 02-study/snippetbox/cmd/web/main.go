/**
 * file: cmd/web/main.go
 * description: file responsible for running the application.
 * data: 12/27/2023
 * author: Glaucia Lemos <Twitter: @glaucia_lemos86>
 */

package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	mux := http.NewServeMux()

	// Configuração do FileServer para servir arquivos estáticos.
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Demais rotas da aplicação
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	logger.Info("Starting the server on port", "addr", *addr)

	err := http.ListenAndServe(*addr, mux)
	logger.Error(err.Error())
	os.Exit(1)
}
