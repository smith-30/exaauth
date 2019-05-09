package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// genApiCmd represents the genApi command
var genApiCmd = &cobra.Command{
	Use:   "generate api boiler",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("generate api boiler called")
	},
}

func init() {
	rootCmd.AddCommand(genApiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genApiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genApiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
