/*
Copyright © 2024 PRATIK GAYEN pratik.neel02@gmail.com
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var showExtended bool

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if showExtended {
			showTasksExtended()
		} else {
			showTasksCompact()
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.PersistentFlags().
		BoolVarP(&showExtended, "all", "a", false, "List expanded view of all tasks")
}
