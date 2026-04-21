package notification

import (
	"context"
	"fmt"
	"os"

	"github.com/cicbyte/answer-cli/internal/output"
	"github.com/spf13/cobra"
)

var notiStatusJSON bool

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show notification unread status",
	Long: `Show unread notification status for inbox and achievement.

Examples:
  answer-cli notification status
  answer-cli notification status --json`,
	Run: runStatus,
}

func init() {
	statusCmd.Flags().BoolVar(&notiStatusJSON, "json", false, "Output in JSON format")
}

func runStatus(cmd *cobra.Command, args []string) {
	cli, err := getClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	status, err := cli.Notification.Status(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if notiStatusJSON || output.GetOutputFormat("") == "json" {
		output.PrintJSON(status)
		return
	}

	inboxStatus := "read"
	if status.Inbox {
		inboxStatus = "unread"
	}
	achievementStatus := "read"
	if status.Achievement {
		achievementStatus = "unread"
	}

	fmt.Printf("Inbox:       %s\n", inboxStatus)
	fmt.Printf("Achievement: %s\n", achievementStatus)
}
