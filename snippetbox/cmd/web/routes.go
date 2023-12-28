package main 

import (
	"net/http"
)

func (app *application) routes() *http.ServeMux{
	mux := http.NewServeMux()
    mux.HandleFunc("/", app.home)
    mux.HandleFunc("/order", app.order)
    mux.HandleFunc("/snippet", app.showOrder)
    mux.HandleFunc("/snippet/create", app.createOrder)

	fs := http.FileServer(neuteredFileSystem{http.Dir("./ui/static/")})
    mux.Handle("/static", http.NotFoundHandler())
    mux.Handle("/static/", http.StripPrefix("/static", fs))

	return mux
}