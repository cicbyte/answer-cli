package user

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

func GetUserCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user",
		Short: "User information",
		Long: `View user information.

Examples:
  answer-cli user info
  answer-cli user info --username=john
  answer-cli user search "keyword"`,
	}
	cmd.AddCommand(infoCmd, searchCmd)
	return cmd
}
