package main

import (
	"app/handlers"
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

	handlers := &handlers.Handlers{
		App: cen,
	}

	cen.InfoLog.Println("Debug is set to", cen.Debug)

	app := &application{
		App:      cen,
		Handlers: handlers,
	}

	app.App.Routes = app.routes()

	return app
}
