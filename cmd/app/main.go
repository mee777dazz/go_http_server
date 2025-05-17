package main

import (
	rabbitMQ "http_server/repository/rabbit_mq"
	"http_server/repository/ram_storage"
	"http_server/usecases/service"
	"log"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"

	"http_server/api/http"
	"http_server/cmd/app/config"
	_ "http_server/docs"
	pkgHttp "http_server/pkg/http"
)

// @title My API
// @version 1.0
// @description This is a sample server.

// @host localhost:8080
// @BasePath /
func main() {
	appFlags := config.ParseFlags()
	var cfg config.AppConfig
	config.MustLoad(appFlags.ConfigPath, &cfg)

	objectRepo := ram_storage.NewObject()
	objectSender, err := rabbitMQ.NewRabbitMQSender(cfg.RabbitMQ)
	if err != nil {
		log.Fatalf("failed creating rabbitMQ: %s", err.Error())
	}
	objectService := service.NewObject(objectRepo, objectSender)
	objectHandlers := http.NewHandler(objectService)

	r := chi.NewRouter()
	r.Get("/swagger/*", httpSwagger.WrapHandler)
	objectHandlers.WithObjectHandlers(r)

	log.Printf("Starting server on %s", cfg.Address)
	if err := pkgHttp.CreateAndRunServer(r, cfg.HTTPConfig); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
