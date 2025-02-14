package controller

import "github.com/spf13/cobra"

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "controller",
		Short:   "Controller management commands",
		Aliases: []string{"controllers"},
	}

	cmd.AddCommand(
		run(),
	)

	return cmd
}
