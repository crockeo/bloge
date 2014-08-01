package controllers

import (
	"github.com/zenazn/goji/web"
	"net/http"
)

func GetStatic(c web.C, w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}
