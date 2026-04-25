package auth

import (
	"context"
	"fmt"
	"os"

	"github.com/cicbyte/answer-cli/internal/client"
	"github.com/cicbyte/answer-cli/internal/common"
	"github.com/cicbyte/answer-cli/internal/output"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "查看认证状态",
	Long: `查看当前认证状态和用户信息。

示例:
  answer-cli auth status`,
	Run: runStatus,
}

func runStatus(cmd *cobra.Command, args []string) {
	appConfig := common.GetAppConfig()

	if appConfig.GetServerURL() == "" {
		fmt.Println("未配置服务器。请先运行 'answer-cli auth login' 登录。")
		return
	}

	fmt.Printf("服务器: %s\n", appConfig.GetServerURL())

	if appConfig.GetServerToken() == "" {
		fmt.Println("状态: 未认证")
		fmt.Println("\n请运行 'answer-cli auth login' 进行认证。")
		return
	}

	// 创建客户端查询当前用户
	c := client.NewClient(&client.Config{
		BaseURL: appConfig.GetServerURL(),
		Token:   appConfig.GetServerToken(),
	})

	authService := client.NewAuthService(c)
	user, err := authService.GetCurrentUser(context.Background())
	if err != nil {
		fmt.Println("状态: 认证失败")
		fmt.Printf("错误: %v\n", err)
		fmt.Println("\n请运行 'answer-cli auth login' 重新认证。")
		os.Exit(1)
	}

	fmt.Println("状态: 已认证 ✓")

	fmt.Println()
	fmt.Printf("  用户名  %s\n", user.Username)
	fmt.Printf("  显示名  %s\n", user.DisplayName)
	if user.EMail != "" {
		fmt.Printf("  邮箱    %s\n", user.EMail)
	}
	fmt.Printf("  等级    %d\n", user.Rank)
	fmt.Printf("  提问    %d  回答 %d  关注 %d\n", user.QuestionCount, user.AnswerCount, user.FollowCount)
	if user.IsAdmin {
		fmt.Printf("  管理员  true\n")
	}
	if user.Bio != "" {
		fmt.Printf("  简介    %s\n", output.Truncate(user.Bio, 60))
	}
	if user.Website != "" {
		fmt.Printf("  网站    %s\n", user.Website)
	}
	if user.Location != "" {
		fmt.Printf("  位置    %s\n", user.Location)
	}
}
