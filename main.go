package main

import (
	"server/server"
)

func main() {
	app := server.NewApp()

	router := server.NewRouter()

	app.Listen(":4000", router)
}
