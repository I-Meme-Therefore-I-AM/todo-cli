package cmd

import (
	"strings"
	"todo-cli/todo"

	"github.com/spf13/cobra"
)

var create = &cobra.Command{
	Use:   "create",
	Short: "For creating user task",
	Run: func(cmd *cobra.Command, args []string) {
		t := todo.Tasks{Description: strings.Join(args, " ")}
		t.Add()
	},
}

func init() {

}
