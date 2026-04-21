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
	answerListJSON bool
)

func getListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list <question-id>",
		Short:   "列出问题的回答",
		Aliases: []string{"ls"},
		Long: `列出指定问题的所有回答。

示例:
  answer-cli answer list 123
  answer-cli answer list 123 --page=2 --size=10
  answer-cli answer list 123 --json`,
		Args: cobra.ExactArgs(1),
		Run:  runList,
	}

	cmd.Flags().IntVarP(&answerListPage, "page", "p", 1, "页码")
	cmd.Flags().IntVarP(&answerListSize, "size", "s", 20, "每页数量")
	cmd.Flags().BoolVar(&answerListJSON, "json", false, "以 JSON 格式输出")

	return cmd
}

func runList(cmd *cobra.Command, args []string) {
	questionID := args[0]

	cli, err := getClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	req := &models.AnswerListReq{
		QuestionID: questionID,
		Page:       answerListPage,
		PageSize:   answerListSize,
	}

	resp, err := cli.Answer.Page(context.Background(), req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if answerListJSON || output.GetOutputFormat("") == "json" {
		output.PrintJSON(resp)
		return
	}

	if len(resp.List) == 0 {
		fmt.Println("No answers found.")
		return
	}

	headers := []string{"ID", "Author", "Votes", "Comments", "Accepted", "Created"}
	var rows [][]string
	for _, a := range resp.List {
		accepted := ""
		if a.Accepted == 1 {
			accepted = "Yes"
		}

		author := ""
		if a.UserInfo != nil {
			author = a.UserInfo.DisplayName
		}

		created := models.FormatTimestamp(a.CreatedAt).Format("2006-01-02 15:04")
		rows = append(rows, []string{
			a.ID,
			author,
			fmt.Sprintf("%d", a.VoteCount),
			fmt.Sprintf("%d", a.CommentCount),
			accepted,
			created,
		})
	}

	output.PrintTable(headers, rows)
	fmt.Printf("\nTotal: %d (page %d, %d per page)\n", resp.Count, answerListPage, answerListSize)
}
