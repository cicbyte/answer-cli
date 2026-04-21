package chat

import (
	"context"
	"fmt"

	"github.com/cicbyte/answer-cli/internal/ai"
	"github.com/cicbyte/answer-cli/internal/client"
	"github.com/cicbyte/answer-cli/internal/common"
	"github.com/spf13/cobra"
)

func GetChatCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "chat [question]",
		Short: "AI 对话，基于 Answer 社区数据回答问题",
		Long: `使用 AI 对话助手回答关于 Answer 社区的问题。
AI 会自动搜索相关的问题、回答、标签和用户信息来生成回答。

示例:
  answer-cli chat "Go 语言如何处理错误？"
  answer-cli chat --non-stream "什么是 REST API？"`,
		Args: cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			nonStream, _ := cmd.Flags().GetBool("non-stream")
			return runChat(cmd.Context(), args[0], nonStream)
		},
	}
}

func init() {
	GetChatCommand().Flags().Bool("non-stream", false, "使用非流式输出")
}

func runChat(ctx context.Context, question string, nonStream bool) error {
	cfg := common.GetAppConfig()
	if cfg == nil || cfg.GetAIProvider() == "" {
		return fmt.Errorf("AI 未配置，请先运行: answer-cli config set ai.api_key <key>")
	}

	apiClient := client.NewClient(&client.Config{
		BaseURL: cfg.GetServerURL(),
		Token:   cfg.GetServerToken(),
	})

	aiService := ai.NewAIService(
		cfg.GetAIProvider(),
		cfg.GetAIBaseURL(),
		cfg.GetAIModel(),
		cfg.GetAIAPIKey(),
		apiClient,
	)

	if nonStream {
		resp, err := aiService.Ask(ctx, question, nil)
		if err != nil {
			return fmt.Errorf("AI 请求失败: %w", err)
		}
		fmt.Println(resp.Answer)
		if resp.PromptTokens > 0 || resp.CompletionTokens > 0 {
			fmt.Printf("\n---\nTokens: %d prompt + %d completion | Model: %s\n",
				resp.PromptTokens, resp.CompletionTokens, resp.Model)
		}
	} else {
		err := aiService.AskStream(ctx, question, nil, func(text string) {
			fmt.Print(text)
		})
		if err != nil {
			return fmt.Errorf("AI 请求失败: %w", err)
		}
		fmt.Println()
	}

	return nil
}
