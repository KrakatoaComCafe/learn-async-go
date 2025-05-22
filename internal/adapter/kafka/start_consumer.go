package kafka

import (
	"context"
	"log"

	"go.uber.org/fx"
)

func StartConsumer(lc fx.Lifecycle, consumer *KafkaConsumer) {
	ctx, cancel := context.WithCancel(context.Background())
	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			consumer.Start(ctx)
			return nil
		},
		OnStop: func(_ context.Context) error {
			cancel()
			log.Println("[Kafka] Shutting down consumer...")
			return consumer.Close()
		},
	})
}
