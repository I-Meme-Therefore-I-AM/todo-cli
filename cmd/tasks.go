package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var description = `
Task command for managing your todo task performs
 - add
 - list
 - delete
for your tasks... very easy to use you'll love it's
`
var taskCommand = &cobra.Command{
	Use:   "task",
	Short: "A simple todo task command",
	Long:  description,

	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	taskCommand.AddCommand(create, listTask)
	if err := taskCommand.Execute(); err != nil {
		fmt.Println(err)
		return
	}

}
