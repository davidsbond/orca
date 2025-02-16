package workflow

import "github.com/spf13/cobra"

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "workflow",
		Short:   "Workflow management commands",
		Aliases: []string{"workflows", "wf"},
	}

	cmd.AddCommand(
		schedule(),
	)

	return cmd
}
