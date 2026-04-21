package notification

import (
	"context"
	"fmt"
	"os"

	"github.com/cicbyte/answer-cli/internal/models"
	"github.com/cicbyte/answer-cli/internal/output"
	"github.com/spf13/cobra"
)

var (
	notificationListType string
	notificationListPage int
	notificationListSize int
	notificationListJSON bool
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List notifications",
	Long: `List notifications with optional filtering.

Examples:
  answer-cli notification list
  answer-cli notification list --type=inbox
  answer-cli notification list --type=achievement --page=2 --size=10`,
	Run: runList,
}

func init() {
	listCmd.Flags().StringVar(&notificationListType, "type", "inbox", "Notification type (inbox|achievement)")
	listCmd.Flags().IntVarP(&notificationListPage, "page", "p", 1, "Page number")
	listCmd.Flags().IntVarP(&notificationListSize, "size", "s", 20, "Page size")
	listCmd.Flags().BoolVar(&notificationListJSON, "json", false, "Output in JSON format")
}

func runList(cmd *cobra.Command, args []string) {
	cli, err := getClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	req := &models.NotificationListReq{
		Page:     notificationListPage,
		Size:     notificationListSize,
		InboxType: notificationListType,
	}

	resp, err := cli.Notification.Page(context.Background(), req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if notificationListJSON || output.GetOutputFormat("") == "json" {
		output.PrintJSON(resp)
		return
	}

	if len(resp.List) == 0 {
		fmt.Println("No notifications found.")
		return
	}

	headers := []string{"ID", "Type", "Read", "Title", "Created"}
	var rows [][]string
	for _, n := range resp.List {
		readStatus := "N"
		if n.IsRead {
			readStatus = "Y"
		}
		created := models.FormatTimestamp(n.CreatedAt).Format("2006-01-02 15:04")
		rows = append(rows, []string{
			n.ID,
			n.Type,
			readStatus,
			output.Truncate(n.Title, 50),
			created,
		})
	}

	output.PrintTable(headers, rows)
	fmt.Printf("\nTotal: %d (page %d, %d per page)\n", resp.Count, notificationListPage, notificationListSize)
}
