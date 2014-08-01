package templater

import (
	"github.com/zenazn/goji/web"
	"html/template"
	"io"
)

func LoadPage(name string) (*template.Template, error) {
	return template.ParseFiles("views/" + name + ".html")
}

func ULoadPage(name string) *template.Template {
	t, err := LoadPage(name)
	if err != nil {
		panic(err)
	}
	return t
}

func SendPageMsg(c web.C, w io.Writer, name string, msg string, data interface{}) {
	var username string
	if u, ok := c.Env["username"].(string); ok {
		username = u
	} else {
		username = ""
	}

	pd := pageData{
		Name:     name,
		Message:  msg,
		Username: username,
		Data:     data,
	}

	ULoadPage("header").Execute(w, pd)
	ULoadPage(name).Execute(w, pd)
	ULoadPage("footer").Execute(w, pd)
}

func SendPage(c web.C, w io.Writer, name string, data interface{}) {
	SendPageMsg(c, w, name, "", data)
}
