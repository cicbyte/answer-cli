package ai

import (
	"context"
	"strings"

	"github.com/cicbyte/answer-cli/internal/client"
	"github.com/sashabaranov/go-openai"
)

type AIService struct {
	llmClient *openai.Client
	model     string
	baseURL   string
	apiClient *client.Client
}

func NewAIService(provider, baseURL, model, apiKey string, apiClient *client.Client) *AIService {
	if provider == "ollama" {
		if !strings.HasSuffix(baseURL, "/v1") {
			baseURL = strings.TrimSuffix(baseURL, "/") + "/v1"
		}
	}

	config := openai.DefaultConfig(apiKey)
	config.BaseURL = baseURL

	return &AIService{
		llmClient: openai.NewClientWithConfig(config),
		model:     model,
		baseURL:   baseURL,
		apiClient: apiClient,
	}
}

type ChatMessage struct {
	Role    string
	Content string
}

type StreamCallback func(text string)

type AskResponse struct {
	Answer           string
	Model            string
	PromptTokens     int
	CompletionTokens int
}

func (s *AIService) AskStream(ctx context.Context, question string, history []ChatMessage, cb StreamCallback) error {
	agent := NewAgent(s.llmClient, s.apiClient, s.model)
	return agent.AskStream(ctx, question, history, cb)
}

func (s *AIService) Ask(ctx context.Context, question string, history []ChatMessage) (*AskResponse, error) {
	agent := NewAgent(s.llmClient, s.apiClient, s.model)
	result, err := agent.Ask(ctx, question, history)
	if err != nil {
		return nil, err
	}
	return result, nil
}
