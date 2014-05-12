package main

import "github.com/martini-contrib/staticbin"

func main() {
	m := staticbin.Classic(Asset)
	m.Use(staticbin.Static("slides", Asset))
	m.Run()
}
