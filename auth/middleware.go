package auth

import (
	"github.com/crockeo/bloge/database"
	"github.com/crockeo/bloge/models"
	"github.com/zenazn/goji/web"
	"net/http"
)

func Middleware() func(*web.C, http.Handler) http.Handler {
	return func(c *web.C, h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			if db, ok := c.Env["db"].(database.DB); ok {
				if cookie, err := r.Cookie("auth"); err == nil {
					if cauth, err := models.AuthFromString(cookie.Value); err == nil {
						auths := make([]models.Auth, 0)
						db.Where(map[string]interface{}{
							"username": cauth.Username,
							"password": cauth.Password,
						}).Find(&auths)

						if len(auths) == 1 && auths[0].Equals(cauth) {
							c.Env["auth"] = cauth
							c.Env["username"] = cauth.Username
							c.Env["password"] = cauth.Password
						}
					}
				}
			}

			h.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}
}
