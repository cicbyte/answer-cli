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
  question_create  创建问题
  question_update  更新问题
  answer_list      列出问题的回答
  answer_create    创建回答
  answer_update    更新回答
  comment_add      添加评论
  tag_search       搜索标签`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := runMcpServer(); err != nil {
				fmt.Fprintf(cmd.ErrOrStderr(), "Error: %v\n", err)
				return err
			}
			return nil
		},
	}
}
