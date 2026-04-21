package question

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/cicbyte/answer-cli/internal/models"
	"github.com/cicbyte/answer-cli/internal/output"
	"github.com/spf13/cobra"
)

var questionGetJSON bool

func getGetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get <id>",
		Short: "查看问题详情",
		Long: `查看问题的详细信息。

示例:
  answer-cli question get 123
  answer-cli question get 123 --json`,
		Args: cobra.ExactArgs(1),
		Run:  runGet,
	}

	cmd.Flags().BoolVar(&questionGetJSON, "json", false, "以 JSON 格式输出")

	return cmd
}

func runGet(cmd *cobra.Command, args []string) {
	id := args[0]

	cli, err := getClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	q, err := cli.Question.Get(context.Background(), id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if questionGetJSON || output.GetOutputFormat("") == "json" {
		output.PrintJSON(q)
		return
	}

	fmt.Printf("ID:          %s\n", q.ID)
	fmt.Printf("Title:       %s\n", q.Title)
	fmt.Printf("Status:      %d\n", q.Status)
	fmt.Printf("Views:       %d (unique: %d)\n", q.ViewCount, q.UniqueViewCount)
	fmt.Printf("Votes:       %d\n", q.VoteCount)
	fmt.Printf("Answers:     %d\n", q.AnswerCount)
	fmt.Printf("Comments:    %d\n", q.CommentCount)
	fmt.Printf("Collections: %d\n", q.CollectionCount)
	fmt.Printf("Follows:     %d\n", q.FollowCount)
	fmt.Printf("Hot Score:   %d\n", q.HotScore)
	fmt.Printf("Created:     %s\n", models.FormatTimestamp(q.CreatedAt).Format("2006-01-02 15:04:05"))
	fmt.Printf("Updated:     %s\n", models.FormatTimestamp(q.UpdatedAt).Format("2006-01-02 15:04:05"))

	if q.AcceptedAnswerID != "" {
		fmt.Printf("Accepted:    %s\n", q.AcceptedAnswerID)
	}

	if q.UserInfo != nil {
		fmt.Printf("Author:      %s\n", q.UserInfo.DisplayName)
	}

	if len(q.Tags) > 0 {
		tagNames := make([]string, len(q.Tags))
		for i, tag := range q.Tags {
			tagNames[i] = tag.DisplayName
		}
		fmt.Printf("Tags:        %s\n", strings.Join(tagNames, ", "))
	}

	fmt.Println()
	fmt.Println("Content:")
	fmt.Println(q.Content)
}
