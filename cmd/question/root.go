package question

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

func GetQuestionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "question",
		Short: "管理问题",
		Long: `管理 Apache Answer 社区中的问题。

示例:
  answer-cli question list
  answer-cli question get <id>
  answer-cli question create --title="How to..." --content="..."
  answer-cli question update <id> --title="New title"
  answer-cli question delete <id>
  answer-cli question close <id>
  answer-cli question reopen <id>`,
	}
	cmd.AddCommand(getListCommand())
	cmd.AddCommand(getGetCommand())
	cmd.AddCommand(getCreateCommand())
	cmd.AddCommand(getUpdateCommand())
	cmd.AddCommand(getDeleteCommand())
	cmd.AddCommand(getCloseCommand())
	cmd.AddCommand(getReopenCommand())
	return cmd
}
