package mcp

import (
	"fmt"

	"github.com/spf13/cobra"
)

func GetMcpCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "mcp",
		Short: "启动 MCP Server",
		Long: `以 stdio 模式启动 MCP Server，让 AI 客户端能直接搜索和操作 Answer 数据。

注册的 Tools:
  question_search  搜索问题
  question_get     获取问题详情
  answer_list      列出问题的回答
  tag_search       搜索标签
  tag_get          获取标签详情
  user_search      搜索用户`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := runMcpServer(); err != nil {
				fmt.Fprintf(cmd.ErrOrStderr(), "Error: %v\n", err)
				return err
			}
			return nil
		},
	}
}
