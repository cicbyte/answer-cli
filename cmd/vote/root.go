package vote

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

func GetVoteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vote",
		Short: "Vote on questions and answers",
		Long: `Vote on questions and answers.

Examples:
  answer-cli vote up <object-id>
  answer-cli vote down <object-id>`,
	}
	cmd.AddCommand(upCmd, downCmd)
	return cmd
}
