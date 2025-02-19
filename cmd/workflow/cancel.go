package workflow

import (
	"errors"

	"github.com/spf13/cobra"

	"github.com/davidsbond/orca/pkg/orca"
)

func cancel() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "cancel run-id",
		Short:   "Cancel a workflow run",
		Long:    "Stops the execution of a scheduled/running workflow run once its current task is completed.",
		Example: "orca workflow cancel 74cd0814-13eb-4e98-b814-ddf684fc2496",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			cl, ok := orca.FromContext(ctx)
			if !ok {
				return errors.New("no client available")
			}

			if err := cl.CancelWorkflowRun(ctx, args[0]); err != nil {
				return err
			}

			return nil
		},
	}

	return cmd
}
