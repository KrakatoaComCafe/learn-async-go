package main

import (
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
			kafka.NewKafkaProducer,
			infra.NewMemoryRepository,
			newMessagePublisher,
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

func newMessagePublisher(p *kafka.KafkaProducer) app.MessagePublisher {
	return p
}
