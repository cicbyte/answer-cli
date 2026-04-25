package comment

import (
	"context"
	"fmt"
	"os"

	"github.com/cicbyte/answer-cli/internal/models"
	"github.com/cicbyte/answer-cli/internal/output"
	"github.com/spf13/cobra"
)

var commentGetJSON bool

var getCmd = &cobra.Command{
	Use:   "get <id>",
	Short: "Get comment detail",
	Long: `Get detailed information about a comment.

Examples:
  answer-cli comment get 123
  answer-cli comment get 123 --json`,
	Args: cobra.ExactArgs(1),
	Run:  runGet,
}

func init() {
	getCmd.Flags().BoolVar(&commentGetJSON, "json", false, "Output in JSON format")
}

func runGet(cmd *cobra.Command, args []string) {
	id := args[0]

	cli, err := getClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	c, err := cli.Comment.Get(context.Background(), id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if commentGetJSON || output.GetOutputFormat("") == "json" {
		output.PrintJSON(c)
		return
	}

	fmt.Printf("ID:        %s\n", c.CommentID)
	fmt.Printf("ObjectID:  %s\n", c.ObjectID)
	if c.QuestionID != "" {
		fmt.Printf("Question:  %s\n", c.QuestionID)
	}
	fmt.Printf("User:      %s\n", c.DisplayAuthor())
	fmt.Printf("Votes:     %d\n", c.VoteCount)
	fmt.Printf("Created:   %s\n", models.FormatTimestamp(c.CreatedAt).Format("2006-01-02 15:04:05"))
	if c.UpdatedAt > 0 {
		fmt.Printf("Updated:   %s\n", models.FormatTimestamp(c.UpdatedAt).Format("2006-01-02 15:04:05"))
	}
	fmt.Println()
	fmt.Println("Content:")
	fmt.Println(c.OriginalText)
}
