package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Maestro",
	Long:  `All software has versions. This is Maestro's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Maestro v0.1.0")
	},
}
