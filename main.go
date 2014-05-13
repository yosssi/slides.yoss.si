package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/staticbin"
)

func main() {
	var m *martini.ClassicMartini
	if martini.Env == martini.Prod {
		m = staticbin.Classic(Asset)
		m.Use(staticbin.Static("slides", Asset))
	} else {
		m = martini.Classic()
		m.Use(martini.Static("slides"))
	}
	m.Run()
}
