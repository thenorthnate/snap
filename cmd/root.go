package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	editDir = "edits"
	rawDir  = "raw"
	rootCmd = &cobra.Command{
		Use:   "snap",
		Short: "Photo manager",
		Long:  `TBD`,
	}
	selectedDir = "selected"
	thumbDir    = "thumbnails"
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	setViperDefaults()

	// 	cobra.OnInitialize(initConfig)

	// 	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	// 	rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "author name for copyright attribution")
	// 	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "name of license for the project")
	// 	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	// 	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	// 	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	// 	viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
	// 	viper.SetDefault("license", "apache")
}

func setViperDefaults() {
	home, err := os.UserHomeDir()
	if err == nil {
		defaultRootPath := filepath.Join(home, "Documents/Photos")
		viper.SetDefault("RootDirectory", defaultRootPath)
	}
}
