package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	version  = ""
	revision = ""
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(verInfo())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func verInfo() string {
	return version + "-" + revision + fmt.Sprintf(" (built with %s)", runtime.Version())
}
