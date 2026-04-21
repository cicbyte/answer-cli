package answer

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var answerDeleteYes bool

func getDeleteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "delete <id>",
		Short:   "删除回答",
		Aliases: []string{"rm", "remove"},
		Long: `删除回答。此操作不可撤销！

示例:
  answer-cli answer delete 456
  answer-cli answer delete 456 --yes`,
		Args: cobra.ExactArgs(1),
		Run:  runDelete,
	}

	cmd.Flags().BoolVarP(&answerDeleteYes, "yes", "y", false, "跳过确认")

	return cmd
}

func runDelete(cmd *cobra.Command, args []string) {
	id := args[0]

	if !answerDeleteYes {
		fmt.Printf("Confirm delete answer #%s? [y/N]: ", id)
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

	if err := cli.Answer.Delete(context.Background(), id); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Answer #%s deleted successfully!\n", id)
}
