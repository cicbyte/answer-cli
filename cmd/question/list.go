package question

import (
	"context"
	"fmt"
	"os"

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
)

func getListCommand() *cobra.Command {
	return NewListAsSearch()
}

func NewListAsSearch() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list [keyword]",
		Short:   "列出问题，支持关键词搜索",
		Aliases: []string{"ls", "search"},
		Long: `列出问题列表，支持排序、标签过滤和关键词搜索。

示例:
  answer-cli question list
  answer-cli question list "chrome"
  answer-cli question list --format json
  answer-cli question list --format jsonl`,
		Run: runList,
	}

	cmd.Flags().StringVarP(&questionListKeyword, "keyword", "k", "", "搜索关键词")
	cmd.Flags().StringVar(&questionListOrder, "order", "newest", "排序 (newest|active|hot|score|unanswered|relevance)")
	cmd.Flags().StringVarP(&questionListTag, "tag", "t", "", "按标签过滤")
	cmd.Flags().IntVarP(&questionListPage, "page", "p", 1, "页码")
	cmd.Flags().IntVarP(&questionListSize, "size", "s", 20, "每页数量")

	return cmd
}

func runList(cmd *cobra.Command, args []string) {
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
		Page: questionListPage, PageSize: questionListSize,
		Order: questionListOrder, Tag: questionListTag,
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

	if output.IsJSON("") {
		items := buildQuestionItems(resp.List)
		if output.IsJSONL("") {
			output.PrintJSONL(items)
		} else {
			output.PrintJSON(map[string]any{"count": resp.Count, "list": items})
		}
		return
	}

	if len(resp.List) == 0 {
		fmt.Println("暂无问题")
		return
	}

	displayItems := make([]output.Item, len(resp.List))
	for i, q := range resp.List {
		author := ""
		if q.UserInfo != nil {
			author = q.UserInfo.DisplayName
		}
		date := models.FormatTimestamp(q.CreatedAt).Format("01-02")
		tags := make([]string, len(q.Tags))
		for j, t := range q.Tags {
			tags[j] = t.DisplayName
		}
		displayItems[i] = output.Item{
			Title:    q.Title,
			Subtitle: fmt.Sprintf("%s  %s  ↑%d  💬%d  #%s", date, author, q.VoteCount, q.AnswerCount, q.ID),
			Tags:     tags,
		}
	}

	output.PrintItems(displayItems, fmt.Sprintf("共 %d 条（第 %d 页）", resp.Count, questionListPage))
}

func runSearch(cli *client.Client, keyword string) {
	order := questionListOrder
	if order == "newest" {
		order = "relevance"
	}

	resp, err := cli.Search.Search(context.Background(), &models.SearchReq{
		Query: keyword, Page: questionListPage, Size: questionListSize, Order: order,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "错误: %v\n", err)
		os.Exit(1)
	}

	if output.IsJSON("") {
		items := make([]map[string]any, 0, len(resp.List))
		for _, sr := range resp.List {
			if sr.Object == nil {
				continue
			}
			obj := sr.Object
			items = append(items, map[string]any{
				"id": obj.ID, "title": obj.Title,
				"answers": obj.AnswerCount,
				"created_at": models.FormatTimestamp(obj.CreatedAt).Format("2006-01-02"),
			})
		}
		if output.IsJSONL("") {
			output.PrintJSONL(items)
		} else {
			output.PrintJSON(map[string]any{"count": resp.Count, "list": items})
		}
		return
	}

	if len(resp.List) == 0 {
		fmt.Printf("未找到与 \"%s\" 相关的结果\n", keyword)
		return
	}

	displayItems := make([]output.Item, 0, len(resp.List))
	for _, sr := range resp.List {
		if sr.Object == nil {
			continue
		}
		obj := sr.Object
		date := models.FormatTimestamp(obj.CreatedAt).Format("01-02")
		displayItems = append(displayItems, output.Item{
			Title:    obj.Title,
			Subtitle: fmt.Sprintf("%s  ↑%d  💬%d", date, obj.VoteCount, obj.AnswerCount),
		})
	}

	output.PrintItems(displayItems, fmt.Sprintf("共 %d 条", resp.Count))
}

func buildQuestionItems(list []models.QuestionListItem) []map[string]any {
	items := make([]map[string]any, 0, len(list))
	for _, q := range list {
		item := map[string]any{
			"id": q.ID, "title": q.Title,
			"answers": q.AnswerCount,
			"created_at": models.FormatTimestamp(q.CreatedAt).Format("2006-01-02"),
		}
		if q.UserInfo != nil {
			item["author"] = q.UserInfo.DisplayName
		}
		if len(q.Tags) > 0 {
			tags := make([]string, len(q.Tags))
			for j, t := range q.Tags {
				tags[j] = t.DisplayName
			}
			item["tags"] = tags
		}
		items = append(items, item)
	}
	return items
}
