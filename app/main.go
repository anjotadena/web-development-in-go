package main

import (
	"app/handlers"

	"github.com/anjotadena/centauri"
)

type application struct {
	App      *centauri.Centauri
	Handlers *handlers.Handlers
}

func main() {
	c := initApplication()

	c.App.ListenAndServe()
}
