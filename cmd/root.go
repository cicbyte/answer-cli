package cmd

import (
	"fmt"
	"os"

	"github.com/cicbyte/answer-cli/cmd/answer"
	"github.com/cicbyte/answer-cli/cmd/auth"
	"github.com/cicbyte/answer-cli/cmd/chat"
	"github.com/cicbyte/answer-cli/cmd/comment"
	"github.com/cicbyte/answer-cli/cmd/config"
	"github.com/cicbyte/answer-cli/cmd/mcp"
	"github.com/cicbyte/answer-cli/cmd/notification"
	"github.com/cicbyte/answer-cli/cmd/question"
	"github.com/cicbyte/answer-cli/cmd/search"
	"github.com/cicbyte/answer-cli/cmd/tag"
	"github.com/cicbyte/answer-cli/cmd/user"
	"github.com/cicbyte/answer-cli/cmd/vote"
	"github.com/cicbyte/answer-cli/internal/common"
	"github.com/cicbyte/answer-cli/internal/log"
	"github.com/cicbyte/answer-cli/internal/utils"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "answer-cli",
	Short: "Apache Answer 命令行工具",
	Long: `answer-cli - Apache Answer Q&A 社区的命令行工具。

通过 CLI 直接操作问答、评论、标签等，也支持 AI 对话和 MCP Server。

示例:
  answer-cli question list           # 列出问题
  answer-cli question get <id>       # 查看问题详情
  answer-cli auth login              # 登录服务器
  answer-cli search "关键词"          # 搜索`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	if err := utils.InitAppDirs(); err != nil {
		fmt.Printf("初始化目录失败: %v\n", err)
		os.Exit(1)
	}
	common.SetAppConfig(utils.ConfigInstance.LoadConfig())
	if err := log.Init(utils.ConfigInstance.GetLogPath()); err != nil {
		fmt.Printf("日志初始化失败: %v\n", err)
		os.Exit(1)
	}

	rootCmd.AddCommand(auth.GetAuthCommand())
	rootCmd.AddCommand(config.GetConfigCommand())
	rootCmd.AddCommand(question.GetQuestionCommand())
	rootCmd.AddCommand(answer.GetAnswerCommand())
	rootCmd.AddCommand(comment.GetCommentCommand())
	rootCmd.AddCommand(tag.GetTagCommand())
	rootCmd.AddCommand(search.GetSearchCommand())
	rootCmd.AddCommand(notification.GetNotificationCommand())
	rootCmd.AddCommand(user.GetUserCommand())
	rootCmd.AddCommand(vote.GetVoteCommand())
	rootCmd.AddCommand(mcp.GetMcpCommand())
	rootCmd.AddCommand(chat.GetChatCommand())
}
