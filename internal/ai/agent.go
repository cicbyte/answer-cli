package ai

import (
	"context"
	"encoding/json"
	"fmt"
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
}

func NewAgent(llmClient *openai.Client, apiClient *client.Client, model string) *Agent {
	return &Agent{llmClient: llmClient, apiClient: apiClient, model: model}
}

func (a *Agent) buildSystemPrompt() string {
	now := time.Now().Format("2006-01-02 15:04:05")
	return fmt.Sprintf(`You are an AI assistant for an Apache Answer Q&A community. You can search questions, answers, tags, and users to help answer questions.

Current time: %s

Available tools:
- question_search: Search questions by keyword, tag, or order
- question_get: Get full question detail by ID
- answer_list: List answers for a question
- tag_search: Search tags
- user_search: Search users

Guidelines:
- Always search for relevant information before answering
- When referencing specific questions or answers, include their IDs
- Provide concise, accurate answers based on the data found
- If no relevant information is found, say so clearly
- Respond in the same language as the user's question`, now)
}

func (a *Agent) getTools() []openai.Tool {
	return []openai.Tool{
		{
			Type: "function",
			Function: &openai.FunctionDefinition{
				Name:        "question_search",
				Description: "Search questions by keyword, tag, or sort order",
				Parameters: map[string]any{
					"type": "object",
					"properties": map[string]any{
						"keyword": map[string]any{"type": "string", "description": "Search keyword"},
						"tag":     map[string]any{"type": "string", "description": "Filter by tag slug"},
						"order":   map[string]any{"type": "string", "description": "Sort: newest|active|hot|score|unanswered"},
						"limit":   map[string]any{"type": "integer", "description": "Max results (default 10)"},
					},
				},
			},
		},
		{
			Type: "function",
			Function: &openai.FunctionDefinition{
				Name:        "question_get",
				Description: "Get full question detail by ID",
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
				Description: "List answers for a specific question",
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
			params.Limit = 10
		}
		result, err := a.apiClient.Question.Page(ctx, &models.QuestionListReq{
			Page: 1, PageSize: params.Limit, Order: params.Order, Tag: params.Tag,
		})
		if err != nil {
			return fmt.Sprintf("error: %v", err)
		}
		data, _ := json.Marshal(result)
		return string(data)

	case "question_get":
		var params struct {
			QuestionID string `json:"question_id"`
		}
		if err := json.Unmarshal([]byte(args), &params); err != nil {
			return fmt.Sprintf("error: %v", err)
		}
		q, err := a.apiClient.Question.Get(ctx, params.QuestionID)
		if err != nil {
			return fmt.Sprintf("error: %v", err)
		}
		data, _ := json.Marshal(q)
		return string(data)

	case "answer_list":
		var params struct {
			QuestionID string `json:"question_id"`
			Limit     int    `json:"limit"`
		}
		if err := json.Unmarshal([]byte(args), &params); err != nil {
			return fmt.Sprintf("error: %v", err)
		}
		if params.Limit <= 0 {
			params.Limit = 10
		}
		result, err := a.apiClient.Answer.Page(ctx, &models.AnswerListReq{
			QuestionID: params.QuestionID, Page: 1, PageSize: params.Limit,
		})
		if err != nil {
			return fmt.Sprintf("error: %v", err)
		}
		data, _ := json.Marshal(result)
		return string(data)

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

	default:
		return fmt.Sprintf("unknown tool: %s", name)
	}
}

func (a *Agent) Ask(ctx context.Context, question string, history []ChatMessage) (*AskResponse, error) {
	messages := []openai.ChatCompletionMessage{
		{Role: "system", Content: a.buildSystemPrompt()},
	}
	for _, msg := range history {
		messages = append(messages, openai.ChatCompletionMessage{Role: msg.Role, Content: msg.Content})
	}
	messages = append(messages, openai.ChatCompletionMessage{Role: "user", Content: question})

	resp, err := a.llmClient.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model:    a.model,
		Messages: messages,
		Tools:    a.getTools(),
	})
	if err != nil {
		return nil, err
	}

	if len(resp.Choices) == 0 || len(resp.Choices[0].Message.ToolCalls) == 0 {
		answer := ""
		if len(resp.Choices) > 0 {
			answer = resp.Choices[0].Message.Content
		}
		return &AskResponse{Answer: answer, Model: a.model}, nil
	}

	assistantMsgs := []openai.ChatCompletionMessage{
		{Role: "assistant", Content: resp.Choices[0].Message.Content},
	}
	messages = append(messages, assistantMsgs...)

	for _, toolCall := range resp.Choices[0].Message.ToolCalls {
		toolResult := a.executeTool(ctx, toolCall.Function.Name, toolCall.Function.Arguments)
		messages = append(messages, openai.ChatCompletionMessage{
			Role:       "tool",
			Content:    toolResult,
			ToolCallID: toolCall.ID,
		})
	}

	resp2, err := a.llmClient.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model:    a.model,
		Messages: messages,
	})
	if err != nil {
		return nil, err
	}

	answer := ""
	if len(resp2.Choices) > 0 {
		answer = resp2.Choices[0].Message.Content
	}

	return &AskResponse{
		Answer:           answer,
		Model:            a.model,
		PromptTokens:     resp.Usage.PromptTokens,
		CompletionTokens: resp2.Usage.CompletionTokens,
	}, nil
}

func (a *Agent) AskStream(ctx context.Context, question string, history []ChatMessage, cb StreamCallback) error {
	messages := []openai.ChatCompletionMessage{
		{Role: "system", Content: a.buildSystemPrompt()},
	}
	for _, msg := range history {
		messages = append(messages, openai.ChatCompletionMessage{Role: msg.Role, Content: msg.Content})
	}
	messages = append(messages, openai.ChatCompletionMessage{Role: "user", Content: question})

	resp, err := a.llmClient.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model:    a.model,
		Messages: messages,
		Tools:    a.getTools(),
	})
	if err != nil {
		return err
	}

	if len(resp.Choices) == 0 || len(resp.Choices[0].Message.ToolCalls) == 0 {
		if len(resp.Choices) > 0 && resp.Choices[0].Message.Content != "" {
			cb(resp.Choices[0].Message.Content)
		}
		return nil
	}

	assistantMsgs := []openai.ChatCompletionMessage{
		{Role: "assistant", Content: resp.Choices[0].Message.Content},
	}
	messages = append(messages, assistantMsgs...)

	for _, toolCall := range resp.Choices[0].Message.ToolCalls {
		toolResult := a.executeTool(ctx, toolCall.Function.Name, toolCall.Function.Arguments)
		zap.L().Debug("tool call", zap.String("name", toolCall.Function.Name))
		messages = append(messages, openai.ChatCompletionMessage{
			Role:       "tool",
			Content:    toolResult,
			ToolCallID: toolCall.ID,
		})
	}

	stream, err := a.llmClient.CreateChatCompletionStream(ctx, openai.ChatCompletionRequest{
		Model:    a.model,
		Messages: messages,
	})
	if err != nil {
		return err
	}
	defer stream.Close()

	for {
		resp, err := stream.Recv()
		if err != nil {
			break
		}
		if len(resp.Choices) > 0 && len(resp.Choices[0].Delta.Content) > 0 {
			cb(resp.Choices[0].Delta.Content)
		}
	}
	return nil
}
