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

var (
	commentAddText    string
	commentAddReplyTo string
)

var addCmd = &cobra.Command{
	Use:   "add <object-id>",
	Short: "Add a comment",
	Long: `Add a comment to a question or answer. Supports pipe input for text.

Examples:
  answer-cli comment add 123 --text="Nice answer!"
  answer-cli comment add 123 --text="Reply" --reply-to=456
  echo "Great question!" | answer-cli comment add 123`,
	Args: cobra.ExactArgs(1),
	Run:  runAdd,
}

func init() {
	addCmd.Flags().StringVarP(&commentAddText, "text", "t", "", "Comment text")
	addCmd.Flags().StringVar(&commentAddReplyTo, "reply-to", "", "Reply to a specific comment ID")
}

func runAdd(cmd *cobra.Command, args []string) {
	objectID := args[0]
	text := commentAddText

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

	req := &models.CommentAddReq{
		ObjectID:       objectID,
		OriginalText:   text,
		ReplyCommentID: commentAddReplyTo,
	}

	result, err := cli.Comment.Add(context.Background(), req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Comment added successfully!\n")
	fmt.Printf("  ID:      %s\n", result.ID)
	fmt.Printf("  Object:  %s\n", result.ObjectID)
}
