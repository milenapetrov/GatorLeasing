package main

import (
	"GatorLeasing/gator-leasing-server/app"
	"GatorLeasing/gator-leasing-server/config"
)

func main() {
	app := app.NewApp(config.GetConfig())
	app.Initialize()
	app.Run()
}
