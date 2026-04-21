package question

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func getReopenCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "reopen <id>",
		Short: "重新打开问题",
		Long: `重新打开已关闭的问题。

示例:
  answer-cli question reopen 123`,
		Args: cobra.ExactArgs(1),
		Run:  runReopen,
	}

	return cmd
}

func runReopen(cmd *cobra.Command, args []string) {
	id := args[0]

	cli, err := getClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if err := cli.Question.Reopen(context.Background(), id); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Question #%s reopened successfully!\n", id)
}
