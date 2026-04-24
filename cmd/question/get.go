package question

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/cicbyte/answer-cli/internal/client"
	"github.com/cicbyte/answer-cli/internal/models"
	"github.com/cicbyte/answer-cli/internal/output"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
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

	printQuestionHeader(q)
	fmt.Println()
	fmt.Println(output.RenderMarkdown(q.Content))

	if q.AnswerCount > 0 {
		answers, err := cli.Answer.Page(context.Background(), &models.AnswerListReq{
			QuestionID: id, Page: 1, PageSize: 20,
		})
		if err == nil && len(answers.List) > 0 {
			printAnswers(answers.List, q.AcceptedAnswerID)
		}
	}
}

func printQuestionHeader(q *models.QuestionInfoResp) {
	author := "匿名"
	if q.UserInfo != nil {
		author = q.UserInfo.DisplayName
		if author == "" {
			author = q.UserInfo.Username
		}
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleRounded)
	t.Style().Color.Row = text.Colors{}
	t.Style().Color.RowAlternate = text.Colors{}
	t.Style().Options.SeparateRows = false
	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, WidthMax: 10},
		{Number: 2, WidthMax: 80},
	})

	t.AppendHeader(table.Row{"属性", "详情"})
	t.AppendRow(table.Row{"标题", text.Bold.Sprint(q.Title)})
	t.AppendRow(table.Row{"作者", author})
	t.AppendRow(table.Row{"状态", statusText(q.Status)})

	if q.AcceptedAnswerID != "" {
		t.AppendRow(table.Row{"采纳", text.FgGreen.Sprint("✓ " + q.AcceptedAnswerID)})
	}

	date := models.FormatTimestamp(q.CreatedAt).Format("2006-01-02 15:04:05")
	t.AppendRow(table.Row{"创建", date})

	stats := fmt.Sprintf("↑%d  💬%d  👁%d", q.VoteCount, q.AnswerCount, q.ViewCount)
	t.AppendRow(table.Row{"统计", stats})

	if len(q.Tags) > 0 {
		tagNames := make([]string, len(q.Tags))
		for i, tag := range q.Tags {
			tagNames[i] = tag.DisplayName
		}
		t.AppendRow(table.Row{"标签", text.FgCyan.Sprint(strings.Join(tagNames, "  "))})
	}

	t.Render()
}

func printAnswers(answers []models.AnswerInfo, acceptedID string) {
	fmt.Println()
	for i, a := range answers {
		author := "匿名"
		if a.UserInfo != nil {
			author = a.UserInfo.DisplayName
			if author == "" {
				author = a.UserInfo.Username
			}
		}
		aDate := models.FormatTimestamp(a.CreatedAt).Format("2006-01-02 15:04")

		title := fmt.Sprintf("回答 %d", i+1)
		if a.ID == acceptedID {
			title += " " + text.FgGreen.Sprint("✓ 已采纳")
		}
		fmt.Println(text.Bold.Sprint(title))

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.SetStyle(table.StyleRounded)
		t.Style().Color.Row = text.Colors{}
		t.Style().Color.RowAlternate = text.Colors{}
		t.Style().Options.SeparateRows = false
		t.SetColumnConfigs([]table.ColumnConfig{
			{Number: 1, WidthMax: 10},
			{Number: 2, WidthMax: 80},
		})

		t.AppendHeader(table.Row{"属性", "详情"})
		t.AppendRow(table.Row{"作者", author})
		t.AppendRow(table.Row{"时间", aDate})
		t.AppendRow(table.Row{"投票", fmt.Sprintf("↑%d", a.VoteCount)})
		t.Render()
		fmt.Println()

		fmt.Println(output.RenderMarkdown(a.Content))
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

var _ *client.Client
