/*
Copyright © 2024 PRATIK GAYEN pratik.neel02@gmail.com
*/
package cmd

import (
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "MArks a task with the valid ID as COMPLETED",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.ParseInt(args[0], 10, 32)
		if err != nil {
			log.Fatal(err)
		}

		markComplete(id)
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
