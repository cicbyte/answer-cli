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

var (
	questionListOrder string
	questionListTag   string
	questionListPage  int
	questionListSize  int
	questionListJSON  bool
)

func getListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Short:   "列出问题",
		Aliases: []string{"ls"},
		Long: `列出问题列表，支持排序和标签过滤。

示例:
  answer-cli question list
  answer-cli question list --order=hot
  answer-cli question list --tag=go
  answer-cli question list --page=2 --size=10
  answer-cli question list --json`,
		Run: runList,
	}

	cmd.Flags().StringVar(&questionListOrder, "order", "newest", "排序方式 (newest|active|hot|score|unanswered)")
	cmd.Flags().StringVarP(&questionListTag, "tag", "t", "", "按标签过滤")
	cmd.Flags().IntVarP(&questionListPage, "page", "p", 1, "页码")
	cmd.Flags().IntVarP(&questionListSize, "size", "s", 20, "每页数量")
	cmd.Flags().BoolVar(&questionListJSON, "json", false, "以 JSON 格式输出")

	return cmd
}

func runList(cmd *cobra.Command, args []string) {
	cli, err := getClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	req := &models.QuestionListReq{
		Page:     questionListPage,
		PageSize: questionListSize,
		Order:    questionListOrder,
		Tag:      questionListTag,
	}

	resp, err := cli.Question.Page(context.Background(), req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if questionListJSON || output.GetOutputFormat("") == "json" {
		output.PrintJSON(resp)
		return
	}

	if len(resp.List) == 0 {
		fmt.Println("No questions found.")
		return
	}

	headers := []string{"ID", "Title", "Answers", "Votes", "Created", "Tags"}
	var rows [][]string
	for _, q := range resp.List {
		tagNames := make([]string, len(q.Tags))
		for i, tag := range q.Tags {
			tagNames[i] = tag.DisplayName
		}
		created := models.FormatTimestamp(q.CreatedAt).Format("2006-01-02 15:04")
		rows = append(rows, []string{
			q.ID,
			output.Truncate(q.Title, 50),
			fmt.Sprintf("%d", q.AnswerCount),
			fmt.Sprintf("%d", q.VoteCount),
			created,
			strings.Join(tagNames, ","),
		})
	}

	output.PrintTable(headers, rows)
	fmt.Printf("\nTotal: %d (page %d, %d per page)\n", resp.Count, questionListPage, questionListSize)
}
