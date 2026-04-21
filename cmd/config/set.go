package config

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/cicbyte/answer-cli/internal/common"
	"github.com/cicbyte/answer-cli/internal/models"
	"github.com/cicbyte/answer-cli/internal/utils"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var setCmd = &cobra.Command{
	Use:   "set <key> [value]",
	Short: "设置配置项的值",
	Long: `设置指定配置项的值。

敏感字段（如 token、api_key）如果不提供 value 参数，会以不回显方式交互式输入。

示例:
  answer-cli config set server.base_url https://answer.example.com
  answer-cli config set ai.model qwen2.5
  answer-cli config set ai.temperature 0.7
  answer-cli config set ai.api_key sk-xxx
  answer-cli config set ai.api_key          # 交互式输入（不回显）
  answer-cli config set log.level debug
  answer-cli config set output.format json`,
	Args: cobra.RangeArgs(1, 2),
	Run:  runSet,
}

// sensitiveKeys 敏感配置项集合
var sensitiveKeys = map[string]bool{
	"server.token": true,
	"ai.api_key":   true,
}

func runSet(cmd *cobra.Command, args []string) {
	key := args[0]

	// 检查 key 是否有效
	_, ok, _ := getConfigValue(key)
	if !ok {
		fmt.Printf("错误: 未知配置项 '%s'\n", key)
		fmt.Println("使用 'answer-cli config list' 查看所有配置项")
		os.Exit(1)
	}

	var value string

	if len(args) >= 2 {
		value = args[1]
	} else if sensitiveKeys[key] {
		// 敏感字段交互式输入（不回显）
		fmt.Printf("请输入 %s: ", key)
		raw, err := term.ReadPassword(int(os.Stdin.Fd()))
		fmt.Println()
		if err != nil {
			fmt.Println("错误: 读取输入失败")
			os.Exit(1)
		}
		value = string(raw)
	} else {
		// 普通字段交互式输入
		fmt.Printf("请输入 %s: ", key)
		reader := bufio.NewReader(os.Stdin)
		line, _ := reader.ReadString('\n')
		value = strings.TrimSpace(line)
	}

	if value == "" {
		fmt.Println("错误: 值不能为空")
		os.Exit(1)
	}

	appConfig := common.GetAppConfig()

	// 类型校验并设置值
	if err := setConfigValue(appConfig, key, value); err != nil {
		fmt.Printf("错误: %v\n", err)
		os.Exit(1)
	}

	// 保存配置
	utils.ConfigInstance.SaveConfig(appConfig)
	common.SetAppConfig(appConfig)

	fmt.Printf("%s 已更新\n", key)
}

// setConfigValue 设置配置值，包含类型校验
func setConfigValue(c *models.AppConfig, key, value string) error {
	switch key {
	case "server.base_url":
		c.Server.BaseURL = value
	case "server.token":
		c.Server.Token = value
	case "ai.provider":
		c.AI.Provider = value
	case "ai.base_url":
		c.AI.BaseURL = value
	case "ai.model":
		c.AI.Model = value
	case "ai.api_key":
		c.AI.ApiKey = value
	case "ai.max_tokens":
		v, err := strconv.Atoi(value)
		if err != nil {
			return fmt.Errorf("无效的整数值: %s", value)
		}
		c.AI.MaxTokens = v
	case "ai.temperature":
		v, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return fmt.Errorf("无效的浮点数值: %s", value)
		}
		c.AI.Temperature = v
	case "ai.timeout":
		v, err := strconv.Atoi(value)
		if err != nil {
			return fmt.Errorf("无效的整数值: %s", value)
		}
		c.AI.Timeout = v
	case "output.format":
		c.Output.Format = value
	case "log.level":
		c.Log.Level = value
	case "log.max_size":
		v, err := strconv.Atoi(value)
		if err != nil {
			return fmt.Errorf("无效的整数值: %s", value)
		}
		c.Log.MaxSize = v
	case "log.max_backups":
		v, err := strconv.Atoi(value)
		if err != nil {
			return fmt.Errorf("无效的整数值: %s", value)
		}
		c.Log.MaxBackups = v
	case "log.max_age":
		v, err := strconv.Atoi(value)
		if err != nil {
			return fmt.Errorf("无效的整数值: %s", value)
		}
		c.Log.MaxAge = v
	case "log.compress":
		v, err := strconv.ParseBool(value)
		if err != nil {
			return fmt.Errorf("无效的布尔值: %s (true/false)", value)
		}
		c.Log.Compress = v
	}
	return nil
}
