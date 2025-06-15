package infra

import (
	domainAuth "github.com/krakatoa/learn-async-go/internal/domain/auth"
	infraAuth "github.com/krakatoa/learn-async-go/internal/infra/auth"
	"go.uber.org/fx"
)

type Config struct {
	SecretKey string
}

func newConfig() Config {
	// em produção, ler de envs
	return Config{
		SecretKey: "super-secret-key",
	}
}

func newJwtService(config Config) domainAuth.TokenService {
	return infraAuth.NewJwtService(config.SecretKey)
}

func Module() fx.Option {
	return fx.Module("infra",
		fx.Provide(
			newConfig,
			newJwtService,
			NewMemoryRepository,
		),
	)
}
