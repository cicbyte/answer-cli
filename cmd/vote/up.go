package vote

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var upCmd = &cobra.Command{
	Use:   "up <object-id>",
	Short: "Upvote a question or answer",
	Long: `Upvote a question or answer.

Examples:
  answer-cli vote up 123`,
	Args: cobra.ExactArgs(1),
	Run:  runUp,
}

func runUp(cmd *cobra.Command, args []string) {
	objectID := args[0]

	cli, err := getClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if err := cli.Vote.VoteUp(context.Background(), objectID); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Upvoted object #%s.\n", objectID)
}
