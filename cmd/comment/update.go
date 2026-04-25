package comment

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/cicbyte/answer-cli/internal/models"
	"github.com/cicbyte/answer-cli/internal/output"
	"github.com/spf13/cobra"
)

var commentUpdateText string

var updateCmd = &cobra.Command{
	Use:   "update <id>",
	Short: "Update a comment",
	Long: `Update an existing comment. Supports pipe input for text.

Examples:
  answer-cli comment update 123 --text="Updated text"
  echo "New content" | answer-cli comment update 123`,
	Args: cobra.ExactArgs(1),
	Run:  runUpdate,
}

func init() {
	updateCmd.Flags().StringVarP(&commentUpdateText, "text", "t", "", "New comment text")
}

func runUpdate(cmd *cobra.Command, args []string) {
	id := args[0]
	text := commentUpdateText

	if text == "" {
		pipeContent, err := output.ReadPipeOrFile("")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading pipe input: %v\n", err)
			os.Exit(1)
		}
		if pipeContent != "" {
			text = strings.TrimSpace(pipeContent)
		}
	}

	if strings.TrimSpace(text) == "" {
		fmt.Fprintln(os.Stderr, "Error: comment text is required (use --text or pipe input)")
		os.Exit(1)
	}

	cli, err := getClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	req := &models.CommentUpdateReq{
		CommentID:    id,
		OriginalText: text,
	}

	if err := cli.Comment.Update(context.Background(), req); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Comment #%s updated successfully.\n", id)
}
