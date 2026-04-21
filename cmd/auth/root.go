package auth

import "github.com/spf13/cobra"

func GetAuthCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auth",
		Short: "认证管理",
		Long: `认证管理 - 登录、登出和查看认证状态。

示例:
  answer-cli auth login --url=https://memos.example.com
  answer-cli auth status
  answer-cli auth logout`,
	}
	cmd.AddCommand(getLoginCommand())
	cmd.AddCommand(getLogoutCommand())
	cmd.AddCommand(getStatusCommand())
	return cmd
}
