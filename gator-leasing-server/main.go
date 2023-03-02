package main

import (
	"GatorLeasing/gator-leasing-server/app"
	"GatorLeasing/gator-leasing-server/config"
)

//	@title			ReLease API
//	@version		1.0
//	@description	Server for subleasing website ReLease

//	@host	localhost:8080

//	@securitydefinitions.oauth2.application	OAuth2Application
//	@tokenUrl								https://dev-nkzmwy1mucvvl5xb.us.auth0.com/oauth/token
//	@description							Auth0 protects our endpoints
func main() {
	app := app.NewApp(config.GetConfig())
	app.Initialize()
	app.Run()
}
