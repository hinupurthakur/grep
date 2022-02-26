package cli

import (
	"github.com/hinupurthakur/grep/cli/constants"
	"github.com/hinupurthakur/grep/helpers"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Version of democtl",
	Long:  `Version of decmoctl installed on your machine`,
	Run: func(cmd *cobra.Command, args []string) {
		version := [][]string{{"client", constants.Version()}}
		header := []string{"grep version"}
		helpers.PrintTable(header, version)
	},
}
