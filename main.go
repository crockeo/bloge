package main

import (
	"github.com/crockeo/bloge/auth"
	"github.com/crockeo/bloge/controllers"
	"github.com/crockeo/bloge/database"
	"github.com/zenazn/goji"
	"regexp"
)

func main() {
	goji.Use(database.Middleware())
	goji.Use(auth.Middleware())

	// Static files
	goji.Get("/public/*", controllers.GetStatic)

	// Other handlers
	goji.Get("/", controllers.GetHome)

	goji.Get(regexp.MustCompile("^/post/(?P<id>[0-9]+)$"), controllers.GetPost)

	goji.Get("/post/new", controllers.GetNewPost)
	goji.Post("/post/new", controllers.PostNewPost)

	goji.Get("/login", controllers.GetLogin)
	goji.Post("/login", controllers.PostLogin)

	goji.Get("/register", controllers.GetRegister)
	goji.Post("/register", controllers.PostRegister)

	goji.Get("/logout", controllers.GetLogout)

	// 404 handler
	goji.NotFound(controllers.NotFoundHandler)

	goji.Serve()
}
