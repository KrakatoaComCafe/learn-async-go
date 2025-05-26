package main

import (
	"log"

	"github.com/krakatoa/learn-async-go/internal/adapter/http"
	"github.com/krakatoa/learn-async-go/internal/adapter/kafka"
	"github.com/krakatoa/learn-async-go/internal/app"
	"github.com/krakatoa/learn-async-go/internal/infra"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			kafka.NewKafkaConsumer,
			newKafkaProducer,
			infra.NewMemoryRepository,
			app.NewMessageService,
			http.NewMessageHandler,
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
