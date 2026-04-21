package server

import "github.com/spf13/cobra"

func GetServerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "管理 Memos 服务器",
		Long: `管理 Memos 服务器配置。

支持添加、列出、删除服务器和设置默认服务器。

示例:
  answer-cli server add
  answer-cli server add --name=my-server --url=https://memos.example.com --token=xxx
  answer-cli server list
  answer-cli server default my-server
  answer-cli server remove my-server`,
	}
	cmd.AddCommand(getAddCommand())
	cmd.AddCommand(getListCommand())
	cmd.AddCommand(getDefaultCommand())
	cmd.AddCommand(getRemoveCommand())
	return cmd
}
