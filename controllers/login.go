package controllers

import (
	"github.com/crockeo/bloge/database"
	"github.com/crockeo/bloge/models"
	"github.com/crockeo/bloge/templater"
	"github.com/zenazn/goji/web"
	"net/http"
)

func GetLogin(c web.C, w http.ResponseWriter, r *http.Request) {
	templater.SendPage(c, w, "auth/login", struct{}{})
}

func PostLogin(c web.C, w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	if username != "" && password != "" {
		db, _ := c.Env["db"].(database.DB)

		auth := models.Auth{
			Username: username,
			Password: password,
		}

		auths := make([]models.Auth, 0)
		db.Where(&auth).Find(&auths)

		if len(auths) == 1 {
			http.SetCookie(w, auth.MakeCookie())
			http.Redirect(w, r, "/", 301)
		} else {
			templater.SendPageMsg(c, w, "auth/login", "Error: No such username/password pair.", struct{}{})
		}
	}
}
