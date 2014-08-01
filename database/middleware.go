package database

import (
	"github.com/zenazn/goji/web"
	"net/http"
)

func Middleware() func(*web.C, http.Handler) http.Handler {
	db := OpenDB()

	return func(c *web.C, h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			c.Env["db"] = db
			h.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}
}
