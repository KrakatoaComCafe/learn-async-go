package kafka

import (
	"context"

	"go.uber.org/fx"
)

func StartConsumer(lc fx.Lifecycle, consumer *KafkaConsumer) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			consumer.Start(ctx)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}
