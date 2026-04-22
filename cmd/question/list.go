package question

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/cicbyte/answer-cli/internal/client"
	"github.com/cicbyte/answer-cli/internal/models"
	"github.com/cicbyte/answer-cli/internal/output"
	"github.com/spf13/cobra"
)

var (
	questionListKeyword string
	questionListOrder   string
	questionListTag     string
	questionListPage    int
	questionListSize    int
	questionListJSON    bool
)

func getListCommand() *cobra.Command {
	return NewListAsSearch()
}

// NewListAsSearch 导出 list 命令，供 search 子命令复用
func NewListAsSearch() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list [keyword]",
		Short:   "列出问题，支持关键词搜索",
		Aliases: []string{"ls", "search"},
		Long: `列出问题列表，支持排序、标签过滤和关键词搜索。
指定关键词时走全文搜索，否则走列表过滤。

示例:
  answer-cli question list
  answer-cli question list "chrome"
  answer-cli question list --order=hot
  answer-cli question list --tag=go
  answer-cli question list "golang" --order=newest --page=2
  answer-cli question list --json`,
		Run: runList,
	}

	cmd.Flags().StringVarP(&questionListKeyword, "keyword", "k", "", "搜索关键词")
	cmd.Flags().StringVar(&questionListOrder, "order", "newest", "排序方式 (newest|active|hot|score|unanswered|relevance)")
	cmd.Flags().StringVarP(&questionListTag, "tag", "t", "", "按标签过滤")
	cmd.Flags().IntVarP(&questionListPage, "page", "p", 1, "页码")
	cmd.Flags().IntVarP(&questionListSize, "size", "s", 20, "每页数量")
	cmd.Flags().BoolVar(&questionListJSON, "json", false, "以 JSON 格式输出")

	return cmd
}

func runList(cmd *cobra.Command, args []string) {
	// 位置参数也可以作为关键词
	keyword := questionListKeyword
	if len(args) > 0 && keyword == "" {
		keyword = args[0]
	}

	cli, err := getClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "错误: %v\n", err)
		os.Exit(1)
	}

	if keyword != "" {
		runSearch(cli, keyword)
	} else {
		runPage(cli)
	}
}

func runPage(cli *client.Client) {
	req := &models.QuestionListReq{
		Page:     questionListPage,
		PageSize: questionListSize,
		Order:    questionListOrder,
		Tag:      questionListTag,
	}

	resp, err := cli.Question.Page(context.Background(), req)
	if err != nil {
		if client.IsNotFoundError(err) && questionListPage > 1 {
			fmt.Printf("暂无数据（第 %d 页超出范围）\n", questionListPage)
		} else {
			fmt.Fprintf(os.Stderr, "错误: %v\n", err)
		}
		os.Exit(1)
	}

	if questionListJSON || output.GetOutputFormat("") == "json" {
		output.PrintJSON(resp)
		return
	}

	if len(resp.List) == 0 {
		fmt.Println("暂无问题")
		return
	}

	headers := []string{"ID", "标题", "回答数", "投票数", "创建时间", "标签"}
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
	fmt.Printf("\n共 %d 条（第 %d 页，每页 %d 条）\n", resp.Count, questionListPage, questionListSize)
}

func runSearch(cli *client.Client, keyword string) {
	order := questionListOrder
	if order == "newest" {
		order = "relevance"
	}

	resp, err := cli.Search.Search(context.Background(), &models.SearchReq{
		Query: keyword,
		Page:  questionListPage,
		Size:  questionListSize,
		Order: order,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "错误: %v\n", err)
		os.Exit(1)
	}

	if questionListJSON || output.GetOutputFormat("") == "json" {
		output.PrintJSON(resp)
		return
	}

	if len(resp.List) == 0 {
		fmt.Printf("未找到与 \"%s\" 相关的结果\n", keyword)
		return
	}

	headers := []string{"ID", "类型", "标题", "回答数", "投票数", "创建时间"}
	var rows [][]string
	for _, item := range resp.List {
		if item.Object == nil {
			continue
		}
		obj := item.Object
		created := models.FormatTimestamp(obj.CreatedAt).Format("2006-01-02 15:04")
		rows = append(rows, []string{
			obj.ID,
			item.ObjectType,
			output.Truncate(obj.Title, 50),
			fmt.Sprintf("%d", obj.AnswerCount),
			fmt.Sprintf("%d", obj.VoteCount),
			created,
		})
	}

	output.PrintTable(headers, rows)
	fmt.Printf("\n共 %d 条（第 %d 页，每页 %d 条）\n", resp.Count, questionListPage, questionListSize)
}
