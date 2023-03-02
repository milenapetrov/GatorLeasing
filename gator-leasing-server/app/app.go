package app

import (
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/config"
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/database"
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/entity"
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/handler"
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/repository"
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/server"
	"github.com/milenapetrov/GatorLeasing/gator-leasing-server/service"
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
	if err := db.GetConnection(a.config.DB); err != nil {
		panic(err)
	}
	if a.config.DB.Migrate {
		db.AutoMigrate()
	}
	if a.config.DB.Populate {

	}

	userContext := entity.NewUserContext()
	tenantUserRepository := repository.NewTenantUserRepository(db.DB)
	tenantUserService := service.NewTenantUserService(userContext, tenantUserRepository)
	leaseRepository := repository.NewLeaseRepository(db.DB)
	leaseService := service.NewLeaseService(userContext, leaseRepository)
	leaseHandler := handler.NewLeaseHandler(leaseService)

	a.server = server.NewServer(a.config.Server, leaseHandler, tenantUserService, userContext)
}

func (a *App) Run() {
	a.server.Run()
}
