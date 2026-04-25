package answer

import (
	"context"
	"fmt"
	"os"

	"github.com/cicbyte/answer-cli/internal/models"
	"github.com/cicbyte/answer-cli/internal/output"
	"github.com/spf13/cobra"
)

var (
	answerListPage int
	answerListSize int
)

func getListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list <question-id>",
		Short:   "列出问题的回答",
		Aliases: []string{"ls"},
		Long: `列出指定问题的所有回答。

示例:
  answer-cli answer list 123
  answer-cli answer list 123 --page=2
  answer-cli answer list 123 --format json
  answer-cli answer list 123 --format jsonl`,
		Args: cobra.ExactArgs(1),
		Run:  runList,
	}

	cmd.Flags().IntVarP(&answerListPage, "page", "p", 1, "页码")
	cmd.Flags().IntVarP(&answerListSize, "size", "s", 20, "每页数量")

	return cmd
}

func runList(cmd *cobra.Command, args []string) {
	questionID := args[0]

	cli, err := getClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "错误: %v\n", err)
		os.Exit(1)
	}

	req := &models.AnswerListReq{
		QuestionID: questionID, Page: answerListPage, PageSize: answerListSize,
	}

	resp, err := cli.Answer.Page(context.Background(), req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "错误: %v\n", err)
		os.Exit(1)
	}

	if output.IsJSON("") {
		items := make([]map[string]any, 0, len(resp.List))
		for _, a := range resp.List {
			author := ""
			if a.UserInfo != nil {
				author = a.UserInfo.DisplayName
			}
			item := map[string]any{
				"id": a.ID, "author": author,
				"accepted":   a.Accepted == 1,
				"created_at": models.FormatTimestamp(a.CreatedAt).Format("2006-01-02"),
			}
			items = append(items, item)
		}
		if output.IsJSONL("") {
			output.PrintJSONL(items)
		} else {
			output.PrintJSON(map[string]any{"count": resp.Count, "list": items})
		}
		return
	}

	if len(resp.List) == 0 {
		fmt.Println("暂无回答")
		return
	}

	headers := []string{"#", "采纳", "日期", "作者", "投票", "评论"}
	var rows [][]string
	for _, a := range resp.List {
		author := "匿名"
		if a.UserInfo != nil && a.UserInfo.DisplayName != "" {
			author = a.UserInfo.DisplayName
		}
		date := models.FormatTimestamp(a.CreatedAt).Format("01-02")
		accepted := ""
		if a.Accepted == 1 {
			accepted = "✓"
		}
		rows = append(rows, []string{
			a.ID, accepted, date, author,
			fmt.Sprintf("%d", a.VoteCount), fmt.Sprintf("%d", a.CommentCount),
		})
	}
	output.PrintTableRight(headers, rows, 5, 6)
	fmt.Printf("\n共 %d 条（第 %d 页）\n", resp.Count, answerListPage)
}
