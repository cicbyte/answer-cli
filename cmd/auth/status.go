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

	fmt.Println("状态: 已认证")
	fmt.Println()

	// 以表格形式显示用户信息
	headers := []string{"字段", "值"}
	rows := [][]string{
		{"用户名", user.Username},
		{"显示名", user.DisplayName},
		{"邮箱", user.EMail},
		{"等级", fmt.Sprintf("%d", user.Rank)},
		{"提问数", fmt.Sprintf("%d", user.QuestionCount)},
		{"回答数", fmt.Sprintf("%d", user.AnswerCount)},
		{"关注数", fmt.Sprintf("%d", user.FollowCount)},
		{"管理员", fmt.Sprintf("%v", user.IsAdmin)},
	}

	if user.Bio != "" {
		rows = append(rows, []string{"简介", output.Truncate(user.Bio, 60)})
	}
	if user.Website != "" {
		rows = append(rows, []string{"网站", user.Website})
	}
	if user.Location != "" {
		rows = append(rows, []string{"位置", user.Location})
	}

	output.PrintTable(headers, rows)
}
