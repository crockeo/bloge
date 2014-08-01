package controllers

import (
	"github.com/crockeo/bloge/database"
	"github.com/crockeo/bloge/models"
	"github.com/crockeo/bloge/templater"
	"github.com/zenazn/goji/web"
	"net/http"
)

func GetRegister(c web.C, w http.ResponseWriter, r *http.Request) {
	templater.SendPage(c, w, "auth/register", struct{}{})
}

func PostRegister(c web.C, w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	cpassword := r.FormValue("cpassword")

	if username != "" && password != "" && cpassword != "" {
		if password == cpassword {
			db, _ := c.Env["db"].(database.DB)

			auth := models.Auth{
				Username: username,
				Password: password,
			}

			auths := make([]models.Auth, 0)
			db.Where(map[string]interface{}{
				"username": auth.Username,
			}).Find(&auths)

			if len(auths) == 0 {
				db.Save(&auth)
				http.Redirect(w, r, "/login", 301)
			} else {
				templater.SendPageMsg(c, w, "auth/register", "Error: User already exists.", struct{}{})
			}
		} else {
			templater.SendPageMsg(c, w, "auth/register", "Error: Passwords do not match.", struct{}{})
		}
	}
}
