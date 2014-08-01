package controllers

import (
	"github.com/zenazn/goji/web"
	"net/http"
)

func GetLogout(c web.C, w http.ResponseWriter, r *http.Request) {
	cook, err := r.Cookie("auth")

	if err == nil {
		cook.MaxAge = -1
		http.SetCookie(w, cook)
	}

	http.Redirect(w, r, "/", 301)
}
