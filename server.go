package main

import (
	"net/http"

	"github.com/codegangsta/negroni"
)

func main() {

	n := negroni.Classic()
	n.Use(negroni.NewStatic(http.Dir(".")))

	// listen on port
	n.Run(":8080")
}
