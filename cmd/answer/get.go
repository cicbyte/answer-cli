package answer

import (
	"context"
	"fmt"
	"os"

	"github.com/cicbyte/answer-cli/internal/models"
	"github.com/cicbyte/answer-cli/internal/output"
	"github.com/spf13/cobra"
)

var answerGetJSON bool

func getGetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get <id>",
		Short: "查看回答详情",
		Long: `查看回答的详细信息。

示例:
  answer-cli answer get 456
  answer-cli answer get 456 --json`,
		Args: cobra.ExactArgs(1),
		Run:  runGet,
	}

	cmd.Flags().BoolVar(&answerGetJSON, "json", false, "以 JSON 格式输出")

	return cmd
}

func runGet(cmd *cobra.Command, args []string) {
	id := args[0]

	cli, err := getClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	a, err := cli.Answer.Get(context.Background(), id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if answerGetJSON || output.GetOutputFormat("") == "json" {
		output.PrintJSON(a)
		return
	}

	fmt.Printf("ID:          %s\n", a.ID)
	fmt.Printf("Question ID: %s\n", a.QuestionID)
	fmt.Printf("Status:      %d\n", a.Status)
	fmt.Printf("Votes:       %d\n", a.VoteCount)
	fmt.Printf("Comments:    %d\n", a.CommentCount)
	fmt.Printf("Accepted:    %s\n", map[bool]string{true: "Yes", false: "No"}[a.Accepted == 1])
	fmt.Printf("Created:     %s\n", models.FormatTimestamp(a.CreatedAt).Format("2006-01-02 15:04:05"))
	fmt.Printf("Updated:     %s\n", models.FormatTimestamp(a.UpdatedAt).Format("2006-01-02 15:04:05"))

	if a.UserInfo != nil {
		fmt.Printf("Author:      %s\n", a.UserInfo.DisplayName)
	}

	fmt.Println()
	fmt.Println("Content:")
	fmt.Println(a.Content)
}
