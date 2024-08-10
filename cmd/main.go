package main

import (
	"MOA/config"
	handlers "MOA/infrastructure/Handlers"
	"MOA/service"
	"MOA/services"

	"net/http"

	"github.com/gorilla/mux"
)

type Main struct {
	userService        *services.UserService
	catergoryService   *services.CategoryService
	transactionService *services.TransactionService

	userHandler *handlers.UserHandler
}

func main() {
	cfg := config.MustLoad()

	logger := service.SetupLogger(cfg.Environment)

	userService, err := services.NewUserService(cfg, logger)
	if err != nil {
		panic(err)
	}

	catergoryService, err := services.NewCategoryService(cfg, logger)
	if err != nil {
		panic(err)
	}

	transactionService, err := services.NewTransactionService(cfg, logger)
	if err != nil {
		panic(err)
	}

	userHandler := handlers.NewUserHandler(userService, logger)

	r := mux.NewRouter()

	userHandler.RegisterRoutes(r)

	http.ListenAndServe(cfg.HttpServer.Host+":"+cfg.HttpServer.Port, r)
}

func (this *Main) addServices(cfg *config.Config) {
	userService, err := services.NewUserService(cfg)
	if err != nil {
		panic(err)
	}
	this.userService = userService

}
