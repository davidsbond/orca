package controller

import (
	"github.com/spf13/cobra"

	"github.com/davidsbond/orca/internal/daemon/controller"
)

func run() *cobra.Command {
	var config controller.Config

	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run a controller instance",
		Long: "Starts a controller instance that will handle worker registration and orchestration of workflows and their " +
			"individual tasks",
		RunE: func(cmd *cobra.Command, args []string) error {
			return controller.Run(cmd.Context(), config)
		},
	}

	flags := cmd.PersistentFlags()
	flags.StringVar(&config.DatabaseURL, "database-url", "postgres://postgres:postges@localhost:5432/postgres?sslmode=disable", "URL for connecting to the database")
	flags.IntVar(&config.HTTPPort, "http-port", 8080, "Port to use for HTTP traffic")
	flags.IntVar(&config.GRPCPort, "grpc-port", 8081, "Port to use for gRPC traffic")

	return cmd
}
