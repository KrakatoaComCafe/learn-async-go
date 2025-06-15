package main

import (
	"log"

	"github.com/krakatoa/learn-async-go/internal/adapter/http"
	"github.com/krakatoa/learn-async-go/internal/adapter/http/middleware"
	"github.com/krakatoa/learn-async-go/internal/adapter/kafka"
	"github.com/krakatoa/learn-async-go/internal/app"
	"github.com/krakatoa/learn-async-go/internal/app/auth"
	"github.com/krakatoa/learn-async-go/internal/infra"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		infra.Module(),
		fx.Provide(
			kafka.NewKafkaConsumer,
			newKafkaProducer,
			auth.NewLoginUseCase,
			app.NewMessageService,
			middleware.NewAuthMiddleware,
			http.NewMessageHandler,
			http.NewAuthHandler,
			http.NewRouter,
		),
		fx.Invoke(
			http.StartServer,
			kafka.StartConsumer,
		),
	)
	app.Run()
}

func newKafkaProducer() app.MessagePublisher {
	producer, err := kafka.NewKafkaProducer()
	if err != nil {
		log.Printf("Error creating producer %+v", err)
	}
	return producer
}
