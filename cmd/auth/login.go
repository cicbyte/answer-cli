package auth

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/cicbyte/answer-cli/internal/client"
	"github.com/cicbyte/answer-cli/internal/common"
	"github.com/cicbyte/answer-cli/internal/log"
	"github.com/cicbyte/answer-cli/internal/utils"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"golang.org/x/term"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "登录 Answer 服务器",
	Long: `登录 Apache Answer 服务器，保存认证令牌到本地配置。

如果未配置服务器地址，会交互式提示输入。

示例:
  answer-cli auth login
  answer-cli auth login --url=https://answer.example.com`,
	Run: runLogin,
}

var loginURL string

func init() {
	loginCmd.Flags().StringVarP(&loginURL, "url", "u", "", "服务器 URL（如 https://answer.example.com）")
}

func runLogin(cmd *cobra.Command, args []string) {
	appConfig := common.GetAppConfig()
	baseURL := loginURL

	// 如果未通过 flag 指定，尝试从配置中读取
	if baseURL == "" {
		baseURL = appConfig.GetServerURL()
		if baseURL != "" {
			fmt.Printf("使用已配置的服务器: %s\n", baseURL)
		}
	}

	// 如果仍然为空，交互式提示
	if baseURL == "" {
		fmt.Print("请输入服务器 URL: ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		baseURL = strings.TrimSpace(input)
	}

	if baseURL == "" {
		fmt.Println("错误: 服务器地址不能为空")
		os.Exit(1)
	}

	baseURL = strings.TrimSuffix(baseURL, "/")

	// 提示输入邮箱
	fmt.Print("请输入邮箱: ")
	reader := bufio.NewReader(os.Stdin)
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	if email == "" {
		fmt.Println("错误: 邮箱不能为空")
		os.Exit(1)
	}

	// 提示输入密码（不回显）
	fmt.Print("请输入密码: ")
	passwordBytes, err := term.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	if err != nil {
		fmt.Printf("读取密码失败: %v\n", err)
		os.Exit(1)
	}
	password := string(passwordBytes)

	if password == "" {
		fmt.Println("错误: 密码不能为空")
		os.Exit(1)
	}

	// 创建客户端（登录时不需要 token）
	c := client.NewClient(&client.Config{
		BaseURL: baseURL,
	})

	// 调用登录 API
	authService := client.NewAuthService(c)
	resp, err := authService.Login(context.Background(), email, password)
	if err != nil {
		log.Error("登录失败", zap.Error(err))
		fmt.Printf("登录失败: %v\n", err)
		os.Exit(1)
	}

	// 保存 token 和 base_url 到配置
	appConfig.Server.BaseURL = baseURL
	appConfig.Server.Token = resp.AccessToken
	utils.ConfigInstance.SaveConfig(appConfig)
	common.SetAppConfig(appConfig)

	// 更新 client 的 token 以便后续调用
	c.SetToken(resp.AccessToken)

	fmt.Println()
	fmt.Println("登录成功!")
	fmt.Printf("  用户名: %s\n", resp.Username)
	if resp.DisplayName != "" {
		fmt.Printf("  显示名: %s\n", resp.DisplayName)
	}
	fmt.Printf("  邮箱: %s\n", resp.EMail)
	fmt.Printf("  服务器: %s\n", baseURL)
	fmt.Println("  配置已保存")
}
