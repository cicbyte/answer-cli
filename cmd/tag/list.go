package tag

import (
	"context"
	"fmt"
	"os"

	"github.com/cicbyte/answer-cli/internal/models"
	"github.com/cicbyte/answer-cli/internal/output"
	"github.com/spf13/cobra"
)

var (
	tagListOrder string
	tagListPage  int
	tagListSize  int
	tagListJSON  bool
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List tags",
	Long: `List all tags with optional sorting.

Examples:
  answer-cli tag list
  answer-cli tag list --order=popular
  answer-cli tag list --order=name --page=2 --size=10
  answer-cli tag list --json`,
	Run: runList,
}

func init() {
	listCmd.Flags().StringVar(&tagListOrder, "order", "popular", "Sort order (popular|name|newest)")
	listCmd.Flags().IntVarP(&tagListPage, "page", "p", 1, "Page number")
	listCmd.Flags().IntVarP(&tagListSize, "size", "s", 20, "Page size")
	listCmd.Flags().BoolVar(&tagListJSON, "json", false, "Output in JSON format")
}

func runList(cmd *cobra.Command, args []string) {
	cli, err := getClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	req := &models.TagListReq{
		Page:      tagListPage,
		PageSize:  tagListSize,
		QueryCond: tagListOrder,
	}

	resp, err := cli.Tag.Page(context.Background(), req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if tagListJSON || output.GetOutputFormat("") == "json" {
		output.PrintJSON(resp)
		return
	}

	if len(resp.List) == 0 {
		fmt.Println("No tags found.")
		return
	}

	headers := []string{"Slug", "Name", "Questions", "Followers"}
	var rows [][]string
	for _, t := range resp.List {
		rows = append(rows, []string{
			t.SlugName,
			t.DisplayName,
			fmt.Sprintf("%d", t.QuestionCount),
			fmt.Sprintf("%d", t.FollowCount),
		})
	}

	output.PrintTable(headers, rows)
	fmt.Printf("\nTotal: %d (page %d, %d per page)\n", resp.Count, tagListPage, tagListSize)
}
