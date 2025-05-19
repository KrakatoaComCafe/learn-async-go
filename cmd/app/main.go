package main

import (
	"github.com/krakatoa/learn-async-go/internal/adapter/http"
	"github.com/krakatoa/learn-async-go/internal/app"
	"github.com/krakatoa/learn-async-go/internal/infra"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			infra.NewMemoryRepository,
			app.NewMessageService,
			http.NewMessageHandler,
			http.NewRouter,
		),
		fx.Invoke(
			http.StartServer,
		),
	)
	app.Run()
}
