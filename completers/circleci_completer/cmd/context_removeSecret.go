package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var context_removeSecretCmd = &cobra.Command{
	Use:   "remove-secret",
	Short: "Remove an environment variable from the named context",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(context_removeSecretCmd).Standalone()
	contextCmd.AddCommand(context_removeSecretCmd)
}
