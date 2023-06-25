package main

import "github.com/anjotadena/centauri"

type application struct {
	App *centauri.Centauri
}

func main() {
	c := initApplication()

	c.App.ListenAndServe()
}
