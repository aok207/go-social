package main

import (
	"log"
)

func main() {
	cfg := config{
		addr: ":6969",
	}

	app := &application{
		config: cfg,
	}

	log.Fatal(app.serve(app.registerRoutes()))
}
