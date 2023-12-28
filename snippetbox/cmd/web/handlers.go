package main

import (
	"encoding/json"
	// "fmt"
	"net/http"
	"strconv"
)

type customer struct {
    Lastname string `json:"surname"`
    Name string `json:"name"`
    Amount uint `json:"sum"`
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        app.notFound(w)
        return
    }
    app.render(w, r, "index.html", nil)
}

func (app *application) order(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "order.html", nil)
}

func (app *application) showOrder(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil || id < 1 {
        app.notFound(w)
        return
    }
	
	s, err := app.orders.Get(id)
    if err != nil {
        app.notFound(w)
        return
    }
    app.render(w, r, "show.page.tmpl", &templateData{Cheque: s})
}

func (app *application) createOrder(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        w.Header().Set("Allow", http.MethodPost)
        app.clientError(w, http.StatusMethodNotAllowed)
        return
    }

    var custom customer
    err := json.NewDecoder(r.Body).Decode(&custom)
    if err != nil {
		app.serverError(w, err)
		return
	}
    
	id, err := app.orders.Insert(custom.Lastname, custom.Name, custom.Amount)

	if err != nil {
		app.serverError(w, err)
		return
	}
    writeJson(w, http.StatusOK, map[string]any {"ok": true, "id": id})
	// http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
}

func writeJson(w http.ResponseWriter, status int, v any) error {
    w.Header().Set("Content-Type", "application/json")
    return json.NewEncoder(w).Encode(v)
}  