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
	"github.com/cicbyte/answer-cli/internal/common"
	"github.com/cicbyte/answer-cli/internal/log"
	"github.com/cicbyte/answer-cli/internal/output"
	"github.com/cicbyte/answer-cli/internal/tui"
	"github.com/cicbyte/answer-cli/internal/utils"
	"github.com/spf13/cobra"
)

var globalFormat string

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
	if len(os.Args) == 1 {
		if err := tui.RunTUI(); err != nil {
			fmt.Fprintf(os.Stderr, "TUI 启动失败: %v\n", err)
			os.Exit(1)
		}
		return
	}
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
	rootCmd.AddCommand(mcp.GetMcpCommand())
	rootCmd.AddCommand(chat.GetChatCommand())

	rootCmd.PersistentFlags().StringVar(&globalFormat, "format", "table", "输出格式 (table|json|jsonl)")
	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		output.SetFormat(globalFormat)

		skipCommands := map[string]bool{"config": true, "auth": true, "mcp": true, "help": true, "completion": true}
		if skipCommands[cmd.Name()] {
			return nil
		}
		if cmd.HasParent() && skipCommands[cmd.Parent().Name()] {
			return nil
		}

		cfg := common.GetAppConfig()
		if cfg.GetServerURL() == "" || cfg.GetServerToken() == "" {
			fmt.Fprintln(os.Stderr, "欢迎使用 answer-cli！")
			if cfg.GetServerURL() == "" {
				fmt.Fprintln(os.Stderr, "  尚未配置服务器，请先登录：")
			} else {
				fmt.Fprintln(os.Stderr, "  尚未登录，请先执行登录：")
			}
			fmt.Fprintln(os.Stderr, "  answer-cli auth login")
			fmt.Fprintln(os.Stderr, "  登录时会交互式引导配置服务器地址、邮箱和密码")
			os.Exit(1)
		}
		return nil
	}
}
