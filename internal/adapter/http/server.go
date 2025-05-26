package http

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func StartServer(lc fx.Lifecycle, router *gin.Engine) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				log.Println("[Fx] Server HTTP starting on port: 8080")
				if err := router.Run(":8080"); err != nil {
					log.Fatal(err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("[Fx] Stopping server HTTP...")
			return nil
		},
	})
}
