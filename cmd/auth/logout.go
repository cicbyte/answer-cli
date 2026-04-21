package auth

import (
	"context"
	"fmt"
	"os"

	"github.com/cicbyte/answer-cli/internal/client"
	"github.com/cicbyte/answer-cli/internal/common"
	"github.com/cicbyte/answer-cli/internal/log"
	"github.com/cicbyte/answer-cli/internal/utils"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "登出 Answer 服务器",
	Long: `登出 Apache Answer 服务器，清除本地已保存的认证令牌。

示例:
  answer-cli auth logout`,
	Run: runLogout,
}

func runLogout(cmd *cobra.Command, args []string) {
	appConfig := common.GetAppConfig()

	if appConfig.GetServerURL() == "" {
		fmt.Println("错误: 未配置服务器地址，请先登录")
		os.Exit(1)
	}

	if appConfig.GetServerToken() == "" {
		fmt.Println("当前未认证，无需登出")
		return
	}

	// 创建客户端调用登出 API
	c := client.NewClient(&client.Config{
		BaseURL: appConfig.GetServerURL(),
		Token:   appConfig.GetServerToken(),
	})

	authService := client.NewAuthService(c)
	err := authService.Logout(context.Background())
	if err != nil {
		// 登出 API 调用失败时仍然清除本地 token
		log.Warn("登出 API 调用失败，仍清除本地令牌", zap.Error(err))
	}

	// 清除本地 token
	appConfig.Server.Token = ""
	utils.ConfigInstance.SaveConfig(appConfig)
	common.SetAppConfig(appConfig)

	fmt.Println("已登出")
}
