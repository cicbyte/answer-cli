package ai

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/cicbyte/answer-cli/internal/client"
	"github.com/cicbyte/answer-cli/internal/models"
	"github.com/sashabaranov/go-openai"
	"go.uber.org/zap"
)

type Agent struct {
	llmClient *openai.Client
	apiClient *client.Client
	model     string
	serverURL string
}

func NewAgent(llmClient *openai.Client, apiClient *client.Client, model string, serverURL string) *Agent {
	return &Agent{llmClient: llmClient, apiClient: apiClient, model: model, serverURL: serverURL}
}

func (a *Agent) buildSystemPrompt() string {
	now := time.Now().Format("2006-01-02 15:04:05")
	return fmt.Sprintf(`你是一个 Apache Answer Q&A 社区的智能助手。你的目标是为用户提供完整、有价值的回答。

当前时间：%s

## 可用工具

- question_search：按关键词搜索问题。**此工具会自动获取最相关问题的详情和回答**，你只需调用一次即可获得完整信息
- question_get：获取问题的完整详情（含正文内容），通常不需要单独调用，question_search 已自动获取
- answer_list：获取问题的所有回答，通常不需要单独调用，question_search 已自动获取
- tag_search：搜索标签
- user_search：搜索用户
- open_question：在浏览器中打开问题页面（仅当用户明确要求，或回答内容包含图片时使用）

## 回答策略

1. **搜索即回答**：调用 question_search 后，工具已经自动为你获取了最相关问题的详情和所有回答。直接基于返回的完整内容组织回答即可
2. **完整回答**：基于搜索到的完整内容（问题描述 + 回答）进行总结和回答，不要只说"请查看问题ID xxx"
3. **结构化输出**：使用 Markdown 格式化回答，包括标题、列表、代码块等
4. **引用来源**：回答时引用问题标题和回答者，方便用户溯源
5. **追问确认**：如果用户的问题模糊，先搜索可能相关的内容，基于搜索结果给出最可能的回答，然后询问是否需要更多信息

## 重要

- question_search 返回的内容已包含问题详情和回答列表，你可以直接用于回答
- 如果搜索结果为空，尝试换关键词重新搜索，或根据你的知识给出建议
- 不要重复调用 question_search，一次搜索即可`, now)
}

func (a *Agent) getTools() []openai.Tool {
	return []openai.Tool{
		{
			Type: "function",
			Function: &openai.FunctionDefinition{
				Name:        "question_search",
				Description: "Search questions by keyword. This tool automatically fetches the top result's detail and answers - you get complete info in one call.",
				Parameters: map[string]any{
					"type": "object",
					"properties": map[string]any{
						"keyword": map[string]any{"type": "string", "description": "Search keyword"},
						"tag":     map[string]any{"type": "string", "description": "Filter by tag slug"},
						"order":   map[string]any{"type": "string", "description": "Sort: newest|active|hot|score"},
						"limit":   map[string]any{"type": "integer", "description": "Max results (default 5)"},
					},
				},
			},
		},
		{
			Type: "function",
			Function: &openai.FunctionDefinition{
				Name:        "question_get",
				Description: "Get full question detail by ID (usually not needed, question_search already fetches this)",
				Parameters: map[string]any{
					"type": "object",
					"properties": map[string]any{
						"question_id": map[string]any{"type": "string", "description": "Question ID"},
					},
					"required": []string{"question_id"},
				},
			},
		},
		{
			Type: "function",
			Function: &openai.FunctionDefinition{
				Name:        "answer_list",
				Description: "List answers for a question (usually not needed, question_search already fetches this)",
				Parameters: map[string]any{
					"type": "object",
					"properties": map[string]any{
						"question_id": map[string]any{"type": "string", "description": "Question ID"},
						"limit":       map[string]any{"type": "integer", "description": "Max results (default 10)"},
					},
					"required": []string{"question_id"},
				},
			},
		},
		{
			Type: "function",
			Function: &openai.FunctionDefinition{
				Name:        "tag_search",
				Description: "Search tags by prefix",
				Parameters: map[string]any{
					"type": "object",
					"properties": map[string]any{
						"query": map[string]any{"type": "string", "description": "Tag name prefix"},
					},
					"required": []string{"query"},
				},
			},
		},
		{
			Type: "function",
			Function: &openai.FunctionDefinition{
				Name:        "user_search",
				Description: "Search users by username",
				Parameters: map[string]any{
					"type": "object",
					"properties": map[string]any{
						"username": map[string]any{"type": "string", "description": "Username to search"},
					},
					"required": []string{"username"},
				},
			},
		},
		{
			Type: "function",
			Function: &openai.FunctionDefinition{
				Name:        "open_question",
				Description: "Open question page in browser. Only call when user explicitly asks to open, or when answer content contains images that user may want to see",
				Parameters: map[string]any{
					"type": "object",
					"properties": map[string]any{
						"question_id": map[string]any{"type": "string", "description": "Question ID"},
					},
					"required": []string{"question_id"},
				},
			},
		},
	}
}

func (a *Agent) executeTool(ctx context.Context, name string, args string) string {
	switch name {
	case "question_search":
		var params struct {
			Keyword string `json:"keyword"`
			Tag     string `json:"tag"`
			Order   string `json:"order"`
			Limit   int    `json:"limit"`
		}
		if err := json.Unmarshal([]byte(args), &params); err != nil {
			return fmt.Sprintf("error: %v", err)
		}
		if params.Limit <= 0 {
			params.Limit = 5
		}

		if params.Keyword != "" {
			order := params.Order
			if order == "" {
				order = "relevance"
			}
			result, err := a.apiClient.Search.Search(ctx, &models.SearchReq{
				Query: params.Keyword, Page: 1, Size: params.Limit, Order: order,
			})
			if err != nil {
				return fmt.Sprintf("error: %v", err)
			}
			// 自动获取 top 结果的详情和回答
			return a.enrichSearchResults(ctx, result)
		}

		result, err := a.apiClient.Question.Page(ctx, &models.QuestionListReq{
			Page: 1, PageSize: params.Limit, Order: params.Order, Tag: params.Tag,
		})
		if err != nil {
			return fmt.Sprintf("error: %v", err)
		}
		return formatQuestionListResults(result)

	case "question_get":
		var params struct {
			QuestionID string `json:"question_id"`
		}
		if err := json.Unmarshal([]byte(args), &params); err != nil {
			return fmt.Sprintf("error: %v", err)
		}
		return a.getQuestionWithAnswers(ctx, params.QuestionID)

	case "answer_list":
		var params struct {
			QuestionID string `json:"question_id"`
			Limit      int    `json:"limit"`
		}
		if err := json.Unmarshal([]byte(args), &params); err != nil {
			return fmt.Sprintf("error: %v", err)
		}
		if params.Limit <= 0 {
			params.Limit = 10
		}
		return a.getAnswersForQuestion(ctx, params.QuestionID, params.Limit)

	case "tag_search":
		var params struct {
			Query string `json:"query"`
		}
		if err := json.Unmarshal([]byte(args), &params); err != nil {
			return fmt.Sprintf("error: %v", err)
		}
		tags, err := a.apiClient.Tag.Search(ctx, params.Query)
		if err != nil {
			return fmt.Sprintf("error: %v", err)
		}
		data, _ := json.Marshal(tags)
		return string(data)

	case "user_search":
		var params struct {
			Username string `json:"username"`
		}
		if err := json.Unmarshal([]byte(args), &params); err != nil {
			return fmt.Sprintf("error: %v", err)
		}
		users, err := a.apiClient.User.SearchUsers(ctx, params.Username)
		if err != nil {
			return fmt.Sprintf("error: %v", err)
		}
		data, _ := json.Marshal(users)
		return string(data)

	case "open_question":
		var params struct {
			QuestionID string `json:"question_id"`
		}
		if err := json.Unmarshal([]byte(args), &params); err != nil {
			return fmt.Sprintf("error: %v", err)
		}
		url := fmt.Sprintf("%s/questions/%s", a.serverURL, params.QuestionID)
		openBrowser(url)
		return fmt.Sprintf("已打开: %s", url)

	default:
		return fmt.Sprintf("unknown tool: %s", name)
	}
}

// enrichSearchResults 搜索后自动获取 top 1 问题的详情和回答
func (a *Agent) enrichSearchResults(ctx context.Context, result *models.SearchResp) string {
	if result == nil || len(result.List) == 0 {
		return "搜索结果为空，没有找到相关问题。请尝试换一个关键词搜索。"
	}

	var b strings.Builder
	fmt.Fprintf(&b, "找到 %d 个相关结果：\n\n", result.Count)

	// 格式化搜索结果列表
	for i, item := range result.List {
		obj := item.Object
		if obj == nil {
			continue
		}
		fmt.Fprintf(&b, "%d. [%s] %s (ID: %s)", i+1, item.ObjectType, obj.Title, obj.ID)
		if item.ObjectType == "answer" && obj.QuestionID != "" {
			fmt.Fprintf(&b, " → 问题ID: %s", obj.QuestionID)
		}
		b.WriteString("\n")
		if obj.Excerpt != "" {
			fmt.Fprintf(&b, "   摘要: %s\n", obj.Excerpt)
		}
	}

	// 找到第一个 question 类型的结果
	var topQuestionID string
	for _, item := range result.List {
		if item.ObjectType == "question" && item.Object != nil {
			topQuestionID = item.Object.ID
			break
		}
		if item.ObjectType == "answer" && item.Object != nil && item.Object.QuestionID != "" {
			topQuestionID = item.Object.QuestionID
			break
		}
	}

	if topQuestionID != "" {
		fmt.Fprintf(&b, "\n=== 最相关问题详情 (ID: %s) ===\n\n", topQuestionID)
		detail := a.getQuestionWithAnswers(ctx, topQuestionID)
		b.WriteString(detail)
	}

	return b.String()
}

func (a *Agent) getQuestionWithAnswers(ctx context.Context, questionID string) string {
	var b strings.Builder

	q, err := a.apiClient.Question.Get(ctx, questionID)
	if err != nil {
		fmt.Fprintf(&b, "获取问题详情失败: %v\n", err)
		return b.String()
	}

	author := "匿名"
	if q.UserInfo != nil && q.UserInfo.DisplayName != "" {
		author = q.UserInfo.DisplayName
	}

	fmt.Fprintf(&b, "## 问题: %s\n", q.Title)
	fmt.Fprintf(&b, "作者: %s | 投票: %d | 回答数: %d\n\n", author, q.VoteCount, q.AnswerCount)
	if q.Content != "" {
		fmt.Fprintf(&b, "### 问题描述\n\n%s\n\n", q.Content)
	}
	if len(q.Tags) > 0 {
		tags := make([]string, 0, len(q.Tags))
		for _, t := range q.Tags {
			tags = append(tags, t.DisplayName)
		}
		fmt.Fprintf(&b, "标签: %s\n\n", strings.Join(tags, ", "))
	}

	answers := a.getAnswersForQuestion(ctx, questionID, 10)
	if answers != "" {
		b.WriteString(answers)
	}

	return b.String()
}

func (a *Agent) getAnswersForQuestion(ctx context.Context, questionID string, limit int) string {
	result, err := a.apiClient.Answer.Page(ctx, &models.AnswerListReq{
		QuestionID: questionID, Page: 1, PageSize: limit,
	})
	if err != nil {
		return fmt.Sprintf("获取回答失败: %v\n", err)
	}
	return formatAnswerList(result)
}

func formatAnswerList(result *models.AnswerListResp) string {
	if result == nil || len(result.List) == 0 {
		return ""
	}
	var b strings.Builder
	fmt.Fprintf(&b, "### 回答列表 (%d 条)\n\n", result.Count)
	for i, a := range result.List {
		author := "匿名"
		if a.UserInfo != nil && a.UserInfo.DisplayName != "" {
			author = a.UserInfo.DisplayName
		}
		accepted := ""
		if a.Accepted == 1 {
			accepted = " [已采纳]"
		}
		fmt.Fprintf(&b, "#### 回答 %d%s\n", i+1, accepted)
		fmt.Fprintf(&b, "作者: %s | 投票: %d\n\n", author, a.VoteCount)
		if a.Content != "" {
			fmt.Fprintf(&b, "%s\n\n", a.Content)
		}
	}
	return b.String()
}

func formatQuestionListResults(result *models.QuestionPageResp) string {
	if result == nil || len(result.List) == 0 {
		return "问题列表为空。"
	}
	var b strings.Builder
	fmt.Fprintf(&b, "共 %d 个问题：\n\n", result.Count)
	for i, q := range result.List {
		author := "匿名"
		if q.Operator != nil {
			author = q.Operator.DisplayName
		}
		fmt.Fprintf(&b, "%d. %s (ID: %s)\n", i+1, q.Title, q.ID)
		fmt.Fprintf(&b, "   作者: %s, 投票: %d, 回答数: %d, 评论: %d\n",
			author, q.VoteCount, q.AnswerCount, q.CommentCount)
	}
	return b.String()
}

func (a *Agent) buildMessages(question string, history []ChatMessage) []openai.ChatCompletionMessage {
	messages := []openai.ChatCompletionMessage{
		{Role: "system", Content: a.buildSystemPrompt()},
	}
	for _, msg := range history {
		messages = append(messages, openai.ChatCompletionMessage{Role: msg.Role, Content: msg.Content})
	}
	messages = append(messages, openai.ChatCompletionMessage{Role: "user", Content: question})
	return messages
}

func (a *Agent) Ask(ctx context.Context, question string, history []ChatMessage) (*AskResponse, error) {
	messages := a.buildMessages(question, history)
	tools := a.getTools()
	var totalUsage openai.Usage

	for range 5 {
		resp, err := a.llmClient.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
			Model:    a.model,
			Messages: messages,
			Tools:    tools,
		})
		if err != nil {
			return nil, err
		}

		totalUsage.PromptTokens += resp.Usage.PromptTokens
		totalUsage.CompletionTokens += resp.Usage.CompletionTokens

		if len(resp.Choices) == 0 || len(resp.Choices[0].Message.ToolCalls) == 0 {
			answer := ""
			if len(resp.Choices) > 0 {
				answer = resp.Choices[0].Message.Content
			}
			return &AskResponse{Answer: answer, Model: a.model,
				PromptTokens: totalUsage.PromptTokens, CompletionTokens: totalUsage.CompletionTokens}, nil
		}

		choice := resp.Choices[0]
		assistantMsg := choice.Message
		if assistantMsg.Content == "" && len(assistantMsg.ToolCalls) > 0 {
			assistantMsg.Content = " "
		}
		messages = append(messages, assistantMsg)

		for _, tc := range choice.Message.ToolCalls {
			result := a.executeTool(ctx, tc.Function.Name, tc.Function.Arguments)
			messages = append(messages, openai.ChatCompletionMessage{
				Role:       "tool",
				Content:    result,
				ToolCallID: tc.ID,
			})
		}
	}

	return nil, fmt.Errorf("agent exceeded max iterations")
}

func (a *Agent) AskStream(ctx context.Context, question string, history []ChatMessage, cb StreamCallback) error {
	messages := a.buildMessages(question, history)
	return a.streamLoop(ctx, messages, cb)
}

func (a *Agent) streamLoop(ctx context.Context, messages []openai.ChatCompletionMessage, cb StreamCallback) error {
	tools := a.getTools()
	var totalUsage openai.Usage

	for range 5 {
		stream, err := a.llmClient.CreateChatCompletionStream(ctx, openai.ChatCompletionRequest{
			Model:    a.model,
			Messages: messages,
			Tools:    tools,
		})
		if err != nil {
			cb(StreamEvent{Type: "error", Content: fmt.Sprintf("请求失败: %v", err)})
			return err
		}

		var assistantContent string
		toolCallMap := make(map[int]*openai.ToolCall)

		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				cb(StreamEvent{Type: "error", Content: fmt.Sprintf("流式读取失败: %v", err)})
				return err
			}

			if resp.Usage != nil {
				totalUsage.PromptTokens += resp.Usage.PromptTokens
				totalUsage.CompletionTokens += resp.Usage.CompletionTokens
			}

			if len(resp.Choices) == 0 {
				continue
			}

			delta := resp.Choices[0].Delta

			if delta.Content != "" {
				assistantContent += delta.Content
			}

			for _, tc := range delta.ToolCalls {
				idx := 0
				if tc.Index != nil {
					idx = *tc.Index
				}
				if _, ok := toolCallMap[idx]; !ok {
					toolCallMap[idx] = &openai.ToolCall{
						ID:   tc.ID,
						Type: tc.Type,
						Function: openai.FunctionCall{
							Name:      tc.Function.Name,
							Arguments: tc.Function.Arguments,
						},
					}
				} else {
					toolCallMap[idx].Function.Arguments += tc.Function.Arguments
					if tc.ID != "" {
						toolCallMap[idx].ID = tc.ID
					}
				}
			}
		}

		assistantMsg := openai.ChatCompletionMessage{Role: "assistant"}
		if len(toolCallMap) > 0 {
			tcs := make([]openai.ToolCall, 0, len(toolCallMap))
			for j := range len(toolCallMap) {
				tcs = append(tcs, *toolCallMap[j])
			}
			assistantMsg.ToolCalls = tcs
		}
		if assistantContent == "" && len(toolCallMap) > 0 {
			assistantContent = " "
		}
		assistantMsg.Content = assistantContent
		messages = append(messages, assistantMsg)

		if len(toolCallMap) == 0 {
			if assistantContent != "" {
				cb(StreamEvent{Type: "content", Content: assistantContent})
			}
			cb(StreamEvent{
				Type:             "done",
				PromptTokens:     totalUsage.PromptTokens,
				CompletionTokens: totalUsage.CompletionTokens,
			})
			return nil
		}

		for j := range len(toolCallMap) {
			tc := toolCallMap[j]
			cb(StreamEvent{Type: "tool_call", Tool: tc.Function.Name})
			zap.L().Debug("tool call", zap.String("name", tc.Function.Name))

			result := a.executeTool(ctx, tc.Function.Name, tc.Function.Arguments)
			summary := toolSummary(tc.Function.Name)
			cb(StreamEvent{Type: "tool_result", Tool: tc.Function.Name, Content: summary})

			messages = append(messages, openai.ChatCompletionMessage{
				Role:       "tool",
				Content:    result,
				ToolCallID: tc.ID,
			})
		}
	}

	return fmt.Errorf("agent exceeded max iterations")
}

func toolSummary(name string) string {
	switch name {
	case "question_search":
		return "搜索并获取详情"
	case "question_get":
		return "获取问题详情"
	case "answer_list":
		return "获取回答列表"
	case "tag_search":
		return "搜索标签"
	case "user_search":
		return "搜索用户"
	case "open_question":
		return "已打开浏览器"
	default:
		return "完成"
	}
}

func openBrowser(url string) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	case "darwin":
		cmd = exec.Command("open", url)
	default:
		cmd = exec.Command("xdg-open", url)
	}
	_ = cmd.Start()
}
