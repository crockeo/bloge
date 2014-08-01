package controllers

import (
	"github.com/crockeo/bloge/database"
	"github.com/crockeo/bloge/models"
	"github.com/crockeo/bloge/templater"
	"github.com/zenazn/goji/web"
	"net/http"
	"time"
)

func GetNewPost(c web.C, w http.ResponseWriter, r *http.Request) {
	if _, ok := c.Env["auth"].(models.Auth); ok {
		templater.SendPage(c, w, "post/new", struct{}{})
	} else {
		templater.SendPage(c, w, "logreq", struct{}{})
	}
}

func PostNewPost(c web.C, w http.ResponseWriter, r *http.Request) {
	if username, ok := c.Env["username"].(string); ok {
		title := r.FormValue("title")
		body := r.FormValue("body")

		if title != "" && body != "" {
			db, _ := c.Env["db"].(database.DB)

			post := models.Post{
				Title:   title,
				Author:  username,
				Body:    body,
				Written: time.Now(),
			}
			db.Save(&post)

			http.Redirect(w, r, "/", 301)
		} else {
			templater.SendPage(c, w, "post/new", struct{}{})
		}
	}
}
