package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var orb_sourceCmd = &cobra.Command{
	Use:   "source",
	Short: "Show the source of an orb",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(orb_sourceCmd).Standalone()
	orbCmd.AddCommand(orb_sourceCmd)
}
