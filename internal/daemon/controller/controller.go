package controller

import "context"

type (
	Config struct {
		GRPCPort    int
		HTTPPort    int
		DatabaseURL string
	}
)

func Run(ctx context.Context, config Config) error {
	return nil
}
