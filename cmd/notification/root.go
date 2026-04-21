package notification

import (
	"fmt"

	"github.com/cicbyte/answer-cli/internal/client"
	"github.com/cicbyte/answer-cli/internal/common"
	"github.com/spf13/cobra"
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

func GetNotificationCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "notification",
		Short: "Manage notifications",
		Long: `Manage notifications.

Examples:
  answer-cli notification list
  answer-cli notification status
  answer-cli notification read --all
  answer-cli notification read <id>`,
	}
	cmd.AddCommand(listCmd, statusCmd, readCmd)
	return cmd
}
