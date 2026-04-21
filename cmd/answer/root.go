package answer

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

func GetAnswerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "answer",
		Short: "管理回答",
		Long: `管理 Apache Answer 社区中的回答。

示例:
  answer-cli answer list <question-id>
  answer-cli answer get <id>
  answer-cli answer create <question-id> --content="..."
  answer-cli answer update <id> --content="..."
  answer-cli answer delete <id>
  answer-cli answer accept <answer-id> --question-id=<question-id>`,
	}
	cmd.AddCommand(getListCommand())
	cmd.AddCommand(getGetCommand())
	cmd.AddCommand(getCreateCommand())
	cmd.AddCommand(getUpdateCommand())
	cmd.AddCommand(getDeleteCommand())
	cmd.AddCommand(getAcceptCommand())
	return cmd
}
