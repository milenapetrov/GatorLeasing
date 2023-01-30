package app

import (
	"GatorLeasing/gator-leasing-server/config"
	"GatorLeasing/gator-leasing-server/database"
	"GatorLeasing/gator-leasing-server/repository"
	"GatorLeasing/gator-leasing-server/server"
	"GatorLeasing/gator-leasing-server/service"
)

type App struct {
	config *config.Config
	server *server.Server
}

func NewApp(config *config.Config) *App {
	return &App{
		config: config,
	}
}

func (a *App) Initialize() {
	var db database.Database
	err := db.GetConnection(a.config.DB)
	if err != nil {
		panic(err)
	}
	db.AutoMigrate()

	leaseRepository := repository.NewLeaseRepository(db.DB)
	leaseService := service.NewLeaseService(leaseRepository)

	a.server = server.NewServer(a.config.Server, leaseService)
}

func (a *App) Run() {
	a.server.Run()
}
