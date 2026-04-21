package comment

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

func GetCommentCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "comment",
		Short: "Manage comments",
		Long: `Manage comments on questions and answers.

Examples:
  answer-cli comment list <object-id>
  answer-cli comment get <id>
  answer-cli comment add <object-id> --text="Nice answer!"
  answer-cli comment update <id> --text="Updated text"
  answer-cli comment delete <id>`,
	}
	cmd.AddCommand(listCmd, getCmd, addCmd, updateCmd, deleteCmd)
	return cmd
}
