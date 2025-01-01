package cmd

import (
	"todo-cli/todo"

	"github.com/spf13/cobra"
)

var all bool
var listDesc = `
default behavior too list all incomplete task...
when used with -a or --all can be use to list all task that are both incomplete and complete
`
var listTask = &cobra.Command{
	Use:          "list",
	Short:        "list only incompleted tasks",
	Long:         listDesc,
	SilenceUsage: true,
	Run: func(cmd *cobra.Command, args []string) {

		todo.List(all)
	},
}

func init() {
	listTask.Flags().BoolVarP(&all, "all", "a", false, "list all task including completed task")
}
