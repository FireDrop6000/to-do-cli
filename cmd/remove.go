/*
Copyright © 2024 PRATIK GAYEN pratik.neel02@gmail.com
*/
package cmd

import (
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Removes a task with a valid ID",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		delTask(id)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
