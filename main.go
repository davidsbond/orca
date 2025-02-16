//go:generate go tool buf generate
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"

	"github.com/spf13/cobra"

	"github.com/davidsbond/orca/cmd/controller"
	"github.com/davidsbond/orca/cmd/workflow"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT)
	defer cancel()

	var version string
	info, ok := debug.ReadBuildInfo()
	if ok {
		version = info.Main.Version
	}

	cmd := &cobra.Command{
		Use:     "orca",
		Short:   "A simple distributed workflow orchestrator",
		Version: version,
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
		TraverseChildren: true,
		SilenceErrors:    true,
		SilenceUsage:     true,
	}

	cmd.AddCommand(
		controller.Command(),
		workflow.Command(),
	)

	if err := cmd.ExecuteContext(ctx); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
