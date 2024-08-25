/*
Copyright © 2024 PRATIK GAYEN pratik.neel02@gmail.com
*/
package cmd

import (
	"strings"

	"github.com/spf13/cobra"
)

var due string

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Adds a new task",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		task := ""
		for _, v := range args {
			task = task + " " + v
		}

		if due != "" {
			addTask(strings.TrimSpace(task), strings.TrimSpace(due))
		} else {
			addTask(strings.TrimSpace(task), "Not Set")
		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
	newCmd.PersistentFlags().
		StringVarP(&due, "due", "d", "", "Enter a due date to complete the task")
}
