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
		Short: "查看问题详情（含回答）",
		Long: `查看问题的详细信息，默认同时展示回答列表。

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
		fmt.Fprintf(os.Stderr, "错误: %v\n", err)
		os.Exit(1)
	}

	q, err := cli.Question.Get(context.Background(), id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "错误: %v\n", err)
		os.Exit(1)
	}

	if questionGetJSON || output.GetOutputFormat("") == "json" {
		output.PrintJSON(q)
		return
	}

	fmt.Printf("ID:          %s\n", q.ID)
	fmt.Printf("标题:        %s\n", q.Title)
	fmt.Printf("状态:        %s\n", statusText(q.Status))
	fmt.Printf("浏览:        %d（独立: %d）\n", q.ViewCount, q.UniqueViewCount)
	fmt.Printf("投票:        %d\n", q.VoteCount)
	fmt.Printf("回答:        %d\n", q.AnswerCount)
	fmt.Printf("评论:        %d\n", q.CommentCount)
	fmt.Printf("收藏:        %d\n", q.CollectionCount)
	fmt.Printf("关注:        %d\n", q.FollowCount)
	fmt.Printf("热度:        %d\n", q.HotScore)
	fmt.Printf("创建时间:    %s\n", models.FormatTimestamp(q.CreatedAt).Format("2006-01-02 15:04:05"))
	fmt.Printf("更新时间:    %s\n", models.FormatTimestamp(q.UpdatedAt).Format("2006-01-02 15:04:05"))

	if q.AcceptedAnswerID != "" {
		fmt.Printf("已采纳:      %s\n", q.AcceptedAnswerID)
	}

	if q.UserInfo != nil {
		name := q.UserInfo.DisplayName
		if name == "" {
			name = q.UserInfo.Username
		}
		fmt.Printf("作者:        %s\n", name)
	}

	if len(q.Tags) > 0 {
		tagNames := make([]string, len(q.Tags))
		for i, tag := range q.Tags {
			tagNames[i] = tag.DisplayName
		}
		fmt.Printf("标签:        %s\n", strings.Join(tagNames, ", "))
	}

	fmt.Println()
	fmt.Println("--- 内容 ---")
	fmt.Println(q.Content)

	// 展示回答
	if q.AnswerCount > 0 {
		answers, err := cli.Answer.Page(context.Background(), &models.AnswerListReq{
			QuestionID: id, Page: 1, PageSize: 20,
		})
		if err == nil && len(answers.List) > 0 {
			fmt.Println()
			fmt.Printf("--- 回答（%d 条）---\n", len(answers.List))
			for i, a := range answers.List {
				author := ""
				if a.UserInfo != nil {
					author = a.UserInfo.DisplayName
					if author == "" {
						author = a.UserInfo.Username
					}
				}
				accepted := ""
				if a.Accepted == 1 {
					accepted = " [已采纳]"
				}
				fmt.Printf("\n## 回答 %d%s (ID: %s)\n", i+1, accepted, a.ID)
				fmt.Printf("作者: %s | 投票: %d | %s\n", author, a.VoteCount,
					models.FormatTimestamp(a.CreatedAt).Format("2006-01-02 15:04"))
				fmt.Println()
				fmt.Println(a.Content)
			}
		}
	}
}

func statusText(s int) string {
	switch s {
	case 1:
		return "正常"
	case 2:
		return "已关闭"
	case 10:
		return "待审核"
	case 11:
		return "已删除"
	default:
		return fmt.Sprintf("未知(%d)", s)
	}
}
