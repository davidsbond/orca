package workflow

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/davidsbond/orca/pkg/orca"
)

func describe() *cobra.Command {
	var (
		jsonOut bool
	)

	cmd := &cobra.Command{
		Use:   "describe run-id",
		Short: "Describe a workflow run",
		Long: "Outputs a tree describing a workflow run. This includes child workflows and tasks, recursing all the way\n" +
			"to the end of the root workflow.",
		Example: "orca workflow describe 74cd0814-13eb-4e98-b814-ddf684fc2496",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			cl, ok := orca.FromContext(ctx)
			if !ok {
				return errors.New("no client available")
			}

			description, err := cl.DescribeWorkflowRun(ctx, args[0])
			if err != nil {
				return err
			}

			if jsonOut {
				return json.NewEncoder(os.Stdout).Encode(description)
			}

			fmt.Println(description)
			return nil
		},
	}

	flags := cmd.PersistentFlags()
	flags.BoolVar(&jsonOut, "json", false, "Output workflow description as JSON")

	return cmd
}
