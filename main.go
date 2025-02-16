//go:generate go tool buf generate
package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"

	"github.com/spf13/cobra"

	"github.com/davidsbond/orca/cmd/controller"
	"github.com/davidsbond/orca/cmd/workflow"
	"github.com/davidsbond/orca/internal/log"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT)
	defer cancel()

	var version string
	info, ok := debug.ReadBuildInfo()
	if ok {
		version = info.Main.Version
	}

	var (
		logLevel string
	)

	cmd := &cobra.Command{
		Use:     "orca",
		Short:   "A simple distributed workflow orchestrator",
		Version: version,
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
		SilenceErrors: true,
		SilenceUsage:  true,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				Level: parseLogLevel(logLevel),
			}))

			ctx = log.ToContext(ctx, logger)
			cmd.SetContext(ctx)
		},
	}

	cmd.AddCommand(
		controller.Command(),
		workflow.Command(),
	)

	flags := cmd.PersistentFlags()
	flags.StringVar(&logLevel, "log-level", "error", "Log level (debug, info, warn, error)")

	if err := cmd.ExecuteContext(ctx); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func parseLogLevel(input string) slog.Level {
	switch input {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelError
	}
}
