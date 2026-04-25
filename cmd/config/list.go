package config

import (
	"fmt"
	"strconv"

	"github.com/cicbyte/answer-cli/internal/common"
	"github.com/cicbyte/answer-cli/internal/output"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "列出所有配置项及当前值",
	Long: `列出所有配置项及当前值。敏感字段（如 token、api_key）会显示脱敏后的值。

示例:
  answer-cli config list`,
	Run: runList,
}

func runList(cmd *cobra.Command, args []string) {
	appConfig := common.GetAppConfig()

	// 构建配置项列表
	type configEntry struct {
		key       string
		section   string
		value     string
		sensitive bool
	}

	entries := []configEntry{
		// Server
		{key: "server.base_url", section: "Server", value: appConfig.Server.BaseURL},
		{key: "server.token", section: "Server", value: appConfig.Server.Token, sensitive: true},
		// AI
		{key: "ai.provider", section: "AI", value: appConfig.AI.Provider},
		{key: "ai.base_url", section: "AI", value: appConfig.AI.BaseURL},
		{key: "ai.model", section: "AI", value: appConfig.AI.Model},
		{key: "ai.api_key", section: "AI", value: appConfig.AI.ApiKey, sensitive: true},
		{key: "ai.max_tokens", section: "AI", value: strconv.Itoa(appConfig.AI.MaxTokens)},
		{key: "ai.temperature", section: "AI", value: fmt.Sprintf("%.2f", appConfig.AI.Temperature)},
		{key: "ai.timeout", section: "AI", value: strconv.Itoa(appConfig.AI.Timeout)},
		// Output
		{key: "output.format", section: "Output", value: appConfig.Output.Format},
		// Log
		{key: "log.level", section: "Log", value: appConfig.Log.Level},
		{key: "log.max_size", section: "Log", value: strconv.Itoa(appConfig.Log.MaxSize)},
		{key: "log.max_backups", section: "Log", value: strconv.Itoa(appConfig.Log.MaxBackups)},
		{key: "log.max_age", section: "Log", value: strconv.Itoa(appConfig.Log.MaxAge)},
		{key: "log.compress", section: "Log", value: strconv.FormatBool(appConfig.Log.Compress)},
	}

	// 显示
	headers := []string{"KEY", "VALUE"}
	rows := make([][]string, 0, len(entries))
	currentSection := ""

	for _, e := range entries {
		// 插入分组标题行
		if e.section != currentSection {
			currentSection = e.section
			rows = append(rows, []string{fmt.Sprintf("[%s]", currentSection), ""})
		}

		displayVal := e.value
		if e.sensitive {
			displayVal = maskValue(e.value)
		}
		if displayVal == "" {
			displayVal = "(未设置)"
		}

		rows = append(rows, []string{e.key, displayVal})
	}

	output.PrintTable(headers, rows)
}

func maskValue(value string) string {
	if value == "" {
		return ""
	}
	runes := []rune(value)
	if len(runes) <= 8 {
		return "******"
	}
	return string(runes[:4]) + "..." + string(runes[len(runes)-4:])
}
