package config

import (
	"github.com/spf13/cobra"
)

func GetConfigCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "管理应用配置",
		Long: `管理 answer-cli 应用配置（服务器、AI、输出格式、日志等参数）。

示例:
  answer-cli config list
  answer-cli config get server.base_url
  answer-cli config set ai.model qwen2.5`,
	}
	cmd.AddCommand(listCmd, getCmd, setCmd)
	return cmd
}
