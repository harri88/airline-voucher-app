//go:build wireinject
// +build wireinject

package main

import (
	"airline/backend/internal/application/service"
	"airline/backend/internal/domain/repository"
	"airline/backend/internal/infrastructure/database"
	"airline/backend/internal/infrastructure/persistence"
	"airline/backend/internal/infrastructure/web"
	"airline/backend/internal/infrastructure/web/handler"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// Provider sets for each layer
var dbProvider = wire.NewSet(database.NewDatabaseConnection)
var repoProvider = wire.NewSet(persistence.NewSQLiteVoucherRepository, wire.Bind(new(repository.IVoucherRepository), new(*persistence.SQLiteVoucherRepository)))
var serviceProvider = wire.NewSet(service.NewSeatGenerator, service.NewVoucherService)
var handlerProvider = wire.NewSet(handler.NewVoucherHandler)
var routerProvider = wire.NewSet(web.NewRouter)

// InitializeApp wires the dependencies and returns the Gin engine.
func InitializeApp() (*gin.Engine, error) {
	wire.Build(
		dbProvider,
		repoProvider,
		serviceProvider,
		handlerProvider,
		routerProvider,
	)
	return nil, nil // Wire will replace this
}
