package search

import (
	"context"
	"fmt"
	"os"

	"github.com/cicbyte/answer-cli/internal/client"
	"github.com/cicbyte/answer-cli/internal/common"
	"github.com/cicbyte/answer-cli/internal/models"
	"github.com/cicbyte/answer-cli/internal/output"
	"github.com/spf13/cobra"
)

var (
	searchOrder string
	searchPage  int
	searchSize  int
	searchJSON  bool
)

func getClient() (*client.Client, error) {
	cfg := common.GetAppConfig()
	if cfg.GetServerURL() == "" {
		return nil, fmt.Errorf("server not configured, run: answer-cli config set server.base_url <url>")
	}
	return client.NewClient(&client.Config{
		BaseURL: cfg.GetServerURL(),
		Token:   cfg.GetServerToken(),
	}), nil
}

func GetSearchCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "search <query>",
		Short: "Search questions and answers",
		Long: `Search for questions and answers.

Examples:
  answer-cli search "golang"
  answer-cli search "golang" --order=relevance
  answer-cli search "golang" --order=newest --page=2 --size=10
  answer-cli search "golang" --json`,
		Args: cobra.ExactArgs(1),
		Run:  runSearch,
	}

	cmd.Flags().StringVar(&searchOrder, "order", "relevance", "Sort order (relevance|newest|active|score)")
	cmd.Flags().IntVarP(&searchPage, "page", "p", 1, "Page number")
	cmd.Flags().IntVarP(&searchSize, "size", "s", 20, "Page size")
	cmd.Flags().BoolVar(&searchJSON, "json", false, "Output in JSON format")

	return cmd
}

func runSearch(cmd *cobra.Command, args []string) {
	query := args[0]

	cli, err := getClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	req := &models.SearchReq{
		Query: query,
		Page:  searchPage,
		Size:  searchSize,
		Order: searchOrder,
	}

	resp, err := cli.Search.Search(context.Background(), req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if searchJSON || output.GetOutputFormat("") == "json" {
		output.PrintJSON(resp)
		return
	}

	if len(resp.List) == 0 {
		fmt.Println("No results found.")
		return
	}

	headers := []string{"ID", "Type", "Title", "Answers", "Votes", "Created"}
	var rows [][]string
	for _, item := range resp.List {
		created := models.FormatTimestamp(item.CreatedAt).Format("2006-01-02 15:04")
		rows = append(rows, []string{
			item.ID,
			item.Type,
			output.Truncate(item.Title, 50),
			fmt.Sprintf("%d", item.AnswerCount),
			fmt.Sprintf("%d", item.VoteCount),
			created,
		})
	}

	output.PrintTable(headers, rows)
	fmt.Printf("\nTotal: %d (page %d, %d per page)\n", resp.Count, searchPage, searchSize)
}
