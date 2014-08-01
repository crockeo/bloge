package controllers

import (
	"github.com/crockeo/bloge/templater"
	"github.com/zenazn/goji/web"
	"net/http"
)

func NotFoundHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	templater.SendPage(c, w, "error/404", struct{}{})
}
