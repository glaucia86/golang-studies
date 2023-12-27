/**
 * file: cmd/web/routes.go
 * description: file responsible for handling the routes of the application.
 * data: 12/27/2023
 * author: Glaucia Lemos <Twitter: @glaucia_lemos86>
 */

package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	// Configuração do FileServer para servir arquivos estáticos.
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Demais rotas da aplicação
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	return mux
}
