package comment

import (
	"context"
	"fmt"
	"os"

	"github.com/cicbyte/answer-cli/internal/models"
	"github.com/cicbyte/answer-cli/internal/output"
	"github.com/spf13/cobra"
)

var (
	commentListPage int
	commentListSize int
	commentListJSON bool
)

var listCmd = &cobra.Command{
	Use:   "list <object-id>",
	Short: "List comments for an object",
	Long: `List comments for a question or answer.

Examples:
  answer-cli comment list 123
  answer-cli comment list 123 --page=2 --size=10
  answer-cli comment list 123 --json`,
	Args: cobra.ExactArgs(1),
	Run:  runList,
}

func init() {
	listCmd.Flags().IntVarP(&commentListPage, "page", "p", 1, "Page number")
	listCmd.Flags().IntVarP(&commentListSize, "size", "s", 20, "Page size")
	listCmd.Flags().BoolVar(&commentListJSON, "json", false, "Output in JSON format")
}

func runList(cmd *cobra.Command, args []string) {
	objectID := args[0]

	cli, err := getClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	req := &models.CommentListReq{
		ObjectID: objectID,
		Page:     commentListPage,
		PageSize: commentListSize,
	}

	resp, err := cli.Comment.Page(context.Background(), req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if commentListJSON || output.GetOutputFormat("") == "json" {
		output.PrintJSON(resp)
		return
	}

	if len(resp.List) == 0 {
		fmt.Println("No comments found.")
		return
	}

	headers := []string{"ID", "User", "Content", "Created"}
	var rows [][]string
	for _, c := range resp.List {
		user := ""
		if c.UserInfo != nil {
			user = c.UserInfo.DisplayName
			if user == "" {
				user = c.UserInfo.Username
			}
		}
		created := models.FormatTimestamp(c.CreatedAt).Format("2006-01-02 15:04")
		rows = append(rows, []string{
			c.ID,
			user,
			output.Truncate(c.OriginalText, 50),
			created,
		})
	}

	output.PrintTable(headers, rows)
	fmt.Printf("\nTotal: %d (page %d, %d per page)\n", resp.Count, commentListPage, commentListSize)
}
