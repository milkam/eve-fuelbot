package cmd

import (
	"fmt"

	"github.com/lunemec/eve-fuelbot/pkg/version"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print FuelBot version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version.VersionString)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
