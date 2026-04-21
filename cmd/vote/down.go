package vote

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var downCmd = &cobra.Command{
	Use:   "down <object-id>",
	Short: "Downvote a question or answer",
	Long: `Downvote a question or answer.

Examples:
  answer-cli vote down 123`,
	Args: cobra.ExactArgs(1),
	Run:  runDown,
}

func runDown(cmd *cobra.Command, args []string) {
	objectID := args[0]

	cli, err := getClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if err := cli.Vote.VoteDown(context.Background(), objectID); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Downvoted object #%s.\n", objectID)
}
