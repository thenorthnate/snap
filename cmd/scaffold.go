package cmd

import (
	"github.com/spf13/cobra"
)

var scaffoldCmd = &cobra.Command{
	Use:   "scaffold",
	Short: "Generate folders for new date",
	Long:  `TBD`,
	Run:   makeScaffold,
}

func makeScaffold(cmd *cobra.Command, args []string) {

}
