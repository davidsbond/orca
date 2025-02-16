package workflow

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/davidsbond/orca/pkg/client"
)

func schedule() *cobra.Command {
	var (
		wait bool
		addr string
	)

	cmd := &cobra.Command{
		Use:   "schedule name [input]",
		Short: "Schedule a workflow run",
		Long: "Schedules a workflow run for a specified workflow using the given input. This will return the workflow \n" +
			"run identifier that can be used to query the current state of the workflow run.\n\n" +
			"The workflow input is optional but must be JSON data when provided.\n\n" +
			"You can use the --wait flag to force the CLI to pol the state of the workflow run and return once the run\n" +
			"is in a finished state (erroneous or otherwise)",
		Args: cobra.RangeArgs(1, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cl, err := client.Dial(addr)
			if err != nil {
				return err
			}

			var input json.RawMessage
			if len(args) > 1 {
				input = json.RawMessage(args[1])
			}

			runID, err := cl.ScheduleWorkflow(cmd.Context(), args[0], input)
			if err != nil {
				return err
			}

			if !wait {
				fmt.Println(runID)
				return nil
			}

			return nil
		},
	}

	flags := cmd.PersistentFlags()
	flags.BoolVarP(&wait, "wait", "w", false, "Wait until the workflow run is complete")
	flags.StringVar(&addr, "addr", "localhost:8081", "Address of the orca controller")

	return cmd
}
