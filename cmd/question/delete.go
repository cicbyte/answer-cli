package question

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var questionDeleteYes bool

func getDeleteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "delete <id>",
		Short:   "删除问题",
		Aliases: []string{"rm", "remove"},
		Long: `删除问题。此操作不可撤销！

示例:
  answer-cli question delete 123
  answer-cli question delete 123 --yes`,
		Args: cobra.ExactArgs(1),
		Run:  runDelete,
	}

	cmd.Flags().BoolVarP(&questionDeleteYes, "yes", "y", false, "跳过确认")

	return cmd
}

func runDelete(cmd *cobra.Command, args []string) {
	id := args[0]

	if !questionDeleteYes {
		fmt.Printf("Confirm delete question #%s? [y/N]: ", id)
		var input string
		fmt.Scanln(&input)
		if input != "y" && input != "Y" && input != "yes" {
			fmt.Println("Cancelled.")
			return
		}
	}

	cli, err := getClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if err := cli.Question.Delete(context.Background(), id); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Question #%s deleted successfully!\n", id)
}
