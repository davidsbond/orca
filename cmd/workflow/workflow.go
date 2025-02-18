package workflow

import (
	"github.com/spf13/cobra"

	"github.com/davidsbond/orca/pkg/orca"
)

func Command() *cobra.Command {
	var (
		addr string
	)

	cmd := &cobra.Command{
		Use:     "workflow",
		Short:   "Workflow management commands",
		Aliases: []string{"workflows", "wf"},
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			cl, err := orca.Dial(addr)
			if err != nil {
				return err
			}

			cmd.SetContext(orca.ToContext(cmd.Context(), cl))
			return nil
		},
	}

	cmd.AddCommand(
		schedule(),
		describe(),
	)

	flags := cmd.PersistentFlags()
	flags.StringVar(&addr, "addr", "localhost:4001", "Address of the orca controller")

	return cmd
}
