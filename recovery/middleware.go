package recovery

import (
	"github.com/crockeo/bloge/templater"
	"github.com/zenazn/goji/web"
	"net/http"
)

type stringable interface {
	String() string
}

func Middleware(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				templater.SendPage(*c, w, "error", r)
			}
		}()

		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
