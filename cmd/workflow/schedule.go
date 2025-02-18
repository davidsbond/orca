package workflow

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/davidsbond/orca/internal/workflow"
	"github.com/davidsbond/orca/pkg/client"
)

func schedule() *cobra.Command {
	var wait bool

	cmd := &cobra.Command{
		Use:   "schedule name [input]",
		Short: "Schedule a workflow run",
		Long: "Schedules a workflow run for a specified workflow using the given input. This will return the workflow \n" +
			"run identifier that can be used to query the current state of the workflow run.\n\n" +
			"The workflow input is optional but must be JSON data when provided.\n\n" +
			"You can use the --wait flag to force the CLI to wait until the workflow run is complete and print its output.\n" +
			"(erroneous or otherwise)",
		Example: `orca workflow schedule ReticulateSplines '{ "Foo": "Bar" }'`,
		Args:    cobra.RangeArgs(1, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			cl, ok := client.FromContext(ctx)
			if !ok {
				return errors.New("no client available")
			}

			var input json.RawMessage
			if len(args) > 1 {
				input = json.RawMessage(args[1])
			}

			runID, err := cl.ScheduleWorkflow(ctx, args[0], input)
			if err != nil {
				return err
			}

			if !wait {
				fmt.Println(runID)
				return nil
			}

			ticker := time.NewTicker(time.Second)
			defer ticker.Stop()

			for {
				select {
				case <-ctx.Done():
					return ctx.Err()
				case <-ticker.C:
					run, err := cl.GetWorkflowRun(ctx, runID)
					switch {
					case errors.Is(err, client.ErrNotFound):
						return fmt.Errorf("workflow run %s does not exist", runID)
					case err != nil:
						return err
					}

					if run.Status == workflow.StatusComplete || run.Status == workflow.StatusSkipped {
						fmt.Println(string(run.Output))
						return nil
					}

					if run.Status == workflow.StatusFailed {
						return errors.New(string(run.Output))
					}
				}
			}
		},
	}

	flags := cmd.PersistentFlags()
	flags.BoolVarP(&wait, "wait", "w", false, "Wait until the workflow run is complete and print its output")

	return cmd
}
