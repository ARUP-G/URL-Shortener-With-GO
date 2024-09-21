package api

import (
	"net/http"

	"github.com/ARUP-G/URL-Shortener-With-GO/handler"
)

func RedirectURL(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.ServeFile(w, r, "./frontend/static/index.html")
	} else {
		handler.Redirect(urlStore)(w, r)
	}
}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	RedirectURL(w, r)
}
