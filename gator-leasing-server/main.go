// Package main ReLease API.
//
// # Server for subleasing website ReLease
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//	Schemes: http, https
//	Host: localhost:8080
//	Version: 1.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	SecurityDefinitions:
//	oauth2:
//	    type: oauth2
//	    authorizationUrl: https://dev-nkzmwy1mucvvl5xb.us.auth0.com/oauth/token
//	    tokenUrl: https://dev-nkzmwy1mucvvl5xb.us.auth0.com/oauth/token
//	    in: header
//
// swagger:meta
package main

import (
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/app"
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/config"
)

//go:generate swagger generate spec -o ./docs/swagger.yaml
//go:generate swagger generate spec -o ./docs/swagger.json
func main() {
	app := app.NewApp(config.GetConfig())
	app.Initialize()
	app.Run()
}
