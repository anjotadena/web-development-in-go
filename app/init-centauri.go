package main

import (
	"log"
	"os"

	"github.com/anjotadena/centauri"
)

func initApplication() *application {
	path, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	cen := &centauri.Centauri{}

	err = cen.New(path)

	if err != nil {
		log.Fatal(err)
	}

	cen.AppName = "Centauri App"

	cen.InfoLog.Println("Debug is set to", cen.Debug)

	app := &application{
		App: cen,
	}

	return app
}
