package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/cicbyte/answer-cli/internal/common"
	"github.com/spf13/cobra"
)

var getShowFlag bool

var getCmd = &cobra.Command{
	Use:   "get <key>",
	Short: "查看配置项的值",
	Long: `查看指定配置项的当前值。

敏感字段（如 token、api_key）默认显示为脱敏值，使用 --show 查看明文。

示例:
  answer-cli config get server.base_url
  answer-cli config get ai.model
  answer-cli config get server.token
  answer-cli config get server.token --show`,
	Args: cobra.ExactArgs(1),
	Run:  runGet,
}

func init() {
	getCmd.Flags().BoolVar(&getShowFlag, "show", false, "显示敏感字段的明文值")
}

func runGet(cmd *cobra.Command, args []string) {
	key := args[0]

	value, ok, sensitive := getConfigValue(key)
	if !ok {
		fmt.Printf("错误: 未知配置项 '%s'\n", key)
		fmt.Println("使用 'answer-cli config list' 查看所有配置项")
		os.Exit(1)
	}

	if value == "" {
		fmt.Printf("%s: (未设置)\n", key)
		return
	}

	if sensitive && !getShowFlag {
		fmt.Printf("%s: %s\n", key, maskValue(value))
		fmt.Println("使用 --show 查看明文")
		return
	}

	fmt.Printf("%s: %s\n", key, value)
}

// getConfigValue 根据配置项键获取值。
// 返回 (值, 是否存在, 是否敏感)。
func getConfigValue(key string) (string, bool, bool) {
	appConfig := common.GetAppConfig()

	switch key {
	case "server.base_url":
		return appConfig.Server.BaseURL, true, false
	case "server.token":
		return appConfig.Server.Token, true, true
	case "ai.provider":
		return appConfig.AI.Provider, true, false
	case "ai.base_url":
		return appConfig.AI.BaseURL, true, false
	case "ai.model":
		return appConfig.AI.Model, true, false
	case "ai.api_key":
		return appConfig.AI.ApiKey, true, true
	case "ai.max_tokens":
		return strconv.Itoa(appConfig.AI.MaxTokens), true, false
	case "ai.temperature":
		return fmt.Sprintf("%.2f", appConfig.AI.Temperature), true, false
	case "ai.timeout":
		return strconv.Itoa(appConfig.AI.Timeout), true, false
	case "output.format":
		return appConfig.Output.Format, true, false
	case "log.level":
		return appConfig.Log.Level, true, false
	case "log.max_size":
		return strconv.Itoa(appConfig.Log.MaxSize), true, false
	case "log.max_backups":
		return strconv.Itoa(appConfig.Log.MaxBackups), true, false
	case "log.max_age":
		return strconv.Itoa(appConfig.Log.MaxAge), true, false
	case "log.compress":
		return strconv.FormatBool(appConfig.Log.Compress), true, false
	default:
		return "", false, false
	}
}
