package memo

import "github.com/spf13/cobra"

func GetMemoCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "memo",
		Short: "管理备忘录",
		Long: `管理备忘录 - 创建、列出、查看、编辑和删除备忘录。

示例:
  answer-cli memo list
  answer-cli memo create --content="Hello, world!"
  answer-cli memo get <memo-id>
  answer-cli memo update <memo-id> --content="Updated content"
  answer-cli memo delete <memo-id>`,
	}
	cmd.AddCommand(getListCommand())
	cmd.AddCommand(getGetCommand())
	cmd.AddCommand(getCreateCommand())
	cmd.AddCommand(getUpdateCommand())
	cmd.AddCommand(getDeleteCommand())
	cmd.AddCommand(getStatsCommand())
	return cmd
}
