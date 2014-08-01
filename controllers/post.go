package controllers

import (
	"github.com/crockeo/bloge/database"
	"github.com/crockeo/bloge/models"
	"github.com/crockeo/bloge/templater"
	"github.com/zenazn/goji/web"
	"net/http"
	"strconv"
)

func GetPost(c web.C, w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(c.URLParams["id"], 10, 64)

	db, _ := c.Env["db"].(database.DB)

	var post models.Post
	db.Where("id = $1", id).First(&post)

	templater.SendPage(c, w, "post/post", post)
}
