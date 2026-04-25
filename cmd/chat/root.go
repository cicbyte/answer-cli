package chat

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/cicbyte/answer-cli/internal/ai"
	"github.com/cicbyte/answer-cli/internal/client"
	"github.com/cicbyte/answer-cli/internal/common"
	"github.com/cicbyte/answer-cli/internal/output"
	"github.com/spf13/cobra"
)

func GetChatCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "chat [question]",
		Short: "AI 对话，基于 Answer 社区数据回答问题",
		Long: `使用 AI 对话助手回答关于 Answer 社区的问题。
AI 会自动搜索相关的问题、回答、标签和用户信息来生成回答。

无参数进入交互式对话模式，提供问题参数则单次问答。

示例:
  answer-cli chat "Go 语言如何处理错误？"
  answer-cli chat`,
		RunE: func(cmd *cobra.Command, args []string) error {
			nonStream, _ := cmd.Flags().GetBool("non-stream")
			svc, err := newAIService()
			if err != nil {
				return err
			}
			if len(args) > 0 {
				return ask(svc, cmd.Context(), args[0], nonStream)
			}
			return runInteractive(svc, cmd.Context(), nonStream)
		},
	}

	cmd.Flags().Bool("non-stream", false, "使用非流式输出")

	return cmd
}

func ask(svc *ai.AIService, ctx context.Context, question string, nonStream bool) error {
	_, err := askWithHistory(svc, ctx, question, nil, nonStream)
	return err
}

func askWithHistory(svc *ai.AIService, ctx context.Context, question string, history []ai.ChatMessage, nonStream bool) (string, error) {
	if nonStream {
		resp, err := svc.Ask(ctx, question, history)
		if err != nil {
			return "", fmt.Errorf("AI 请求失败: %w", err)
		}
		fmt.Print(output.RenderMarkdown(resp.Answer))
		if resp.PromptTokens > 0 || resp.CompletionTokens > 0 {
			fmt.Printf("\n%s\nTokens: %d prompt + %d completion | Model: %s\n",
				output.Dim("---"), resp.PromptTokens, resp.CompletionTokens, resp.Model)
		}
		return resp.Answer, nil
	}

	start := time.Now()
	var buf strings.Builder
	var promptTokens, completionTokens int

	err := svc.AskStream(ctx, question, history, func(event ai.StreamEvent) {
		switch event.Type {
		case "content":
			buf.WriteString(event.Content)
		case "tool_call":
			fmt.Printf("  %s %s\n", output.Dim("▸"), event.Tool)
		case "tool_result":
			fmt.Printf("  %s %s\n", output.Dim("✓"), event.Content)
		case "done":
			promptTokens = event.PromptTokens
			completionTokens = event.CompletionTokens
		case "error":
			fmt.Printf("  %s\n", output.Dim("✗ "+event.Content))
		}
	})

	if err != nil {
		fmt.Println()
		return "", fmt.Errorf("AI 请求失败: %w", err)
	}

	raw := buf.String()
	if raw == "" {
		return "", nil
	}

	fmt.Print(output.RenderMarkdown(raw))

	if promptTokens > 0 || completionTokens > 0 {
		elapsed := time.Since(start)
		fmt.Printf("\n%s\n", output.Dim(fmt.Sprintf("Tokens: %d + %d · %.1fs",
			promptTokens, completionTokens, elapsed.Seconds())))
	}

	return raw, nil
}

func runInteractive(svc *ai.AIService, ctx context.Context, nonStream bool) error {
	fmt.Println(output.Bold(" AI 对话模式"))
	fmt.Println(output.Dim("  输入问题开始对话，/quit 退出，/clear 清除上下文"))
	fmt.Println()

	var history []ai.ChatMessage
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(output.Bold("  user > "))
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if line == "/quit" || line == "/exit" || line == "/q" {
			fmt.Println(output.Dim("  再见!"))
			break
		}
		if line == "/clear" {
			history = nil
			fmt.Println(output.Dim("  对话上下文已清除"))
			fmt.Println()
			continue
		}

		resp, err := askWithHistory(svc, ctx, line, history, nonStream)
		if err != nil {
			fmt.Fprintf(os.Stderr, "  错误: %v\n", err)
			fmt.Println()
			continue
		}

		history = append(history, ai.ChatMessage{Role: "user", Content: line})
		history = append(history, ai.ChatMessage{Role: "assistant", Content: resp})
		fmt.Println()
	}

	return nil
}

func newAIService() (*ai.AIService, error) {
	cfg := common.GetAppConfig()
	if cfg == nil || cfg.GetAIProvider() == "" {
		return nil, fmt.Errorf("AI 未配置，请先运行: answer-cli config set ai.api_key <key>")
	}

	apiClient := client.NewClient(&client.Config{
		BaseURL: cfg.GetServerURL(),
		Token:   cfg.GetServerToken(),
	})

	return ai.NewAIService(
		cfg.GetAIProvider(),
		cfg.GetAIBaseURL(),
		cfg.GetAIModel(),
		cfg.GetAIAPIKey(),
		apiClient,
		cfg.GetServerURL(),
	), nil
}
