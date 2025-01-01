package cmd

import "github.com/spf13/cobra"

var descCompleteTask = `
Allows task to be marked as complete

Usage:
	task complete <task-id>
`
var completeTask = &cobra.Command{
	Use:   "complete",
	Short: "User mark task has complete",
	Long:  descCompleteTask,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

	},
}
