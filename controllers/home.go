package controllers

import (
	"github.com/crockeo/bloge/database"
	"github.com/crockeo/bloge/models"
	"github.com/crockeo/bloge/templater"
	"github.com/zenazn/goji/web"
	"net/http"
)

func GetHome(c web.C, w http.ResponseWriter, r *http.Request) {
	db, _ := c.Env["db"].(database.DB)

	posts := make([]models.Post, 0)
	db.Order("id desc").Find(&posts)

	templater.SendPage(c, w, "home", struct {
		Posts []models.Post
	}{
		Posts: posts,
	})
}
