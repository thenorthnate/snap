package cmd

import (
	"github.com/spf13/cobra"
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Perform all sync operations for all photos",
	Long:  `TBD`,
	Run:   sync,
}

func init() {
	// selectCmd.Flags().BoolP("all", "a", false, "Check all available folders")
	// syncCmd.Flags().StringP("date", "d", "", "Date folder to look in")
	rootCmd.AddCommand(syncCmd)
}

func sync(cmd *cobra.Command, args []string) {
	// First check if there is an SSD attached...
	// Check the photos on the SSD to see if there are any new ones on there
	// Create new folders in the target location for any new photos and copy them off the SSD

	// Next check if there are any photos in any of the "selected" directories that don't have
	// their raw photos copied into the "edits" directory

	// Then check for any .jpg/.png photos in the edits directory and move them to the
	// final product directory
}
