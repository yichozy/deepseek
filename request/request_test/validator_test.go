package request_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yichozy/deepseek"
	"github.com/yichozy/deepseek/request"
)

func TestValidateChatCompletionsRequest(t *testing.T) {
	req := &request.ChatCompletionsRequest{
		Messages: []*request.Message{
			{
				Role:    request.RoleUser,
				Content: "Hello",
			},
		},
		Model:            deepseek.DEEPSEEK_CHAT_MODEL,
		FrequencyPenalty: 1,
		MaxTokens:        1,
		PresencePenalty:  1,
		ResponseFormat: &request.ResponseFormat{
			Type: request.ResponseFormatText,
		},
		Stop:   []string{"MOD1"},
		Stream: true,
		StreamOptions: &request.StreamOptions{
			IncludeUsage: true,
		},
		Temperature: 2.0,
		TopP:        nil, // TODO: VN -- pass non nil
	}
	err := request.ValidateChatCompletionsRequest(req)
	assert.NoError(t, err)
	fmt.Println(err)
}

func TestValidateChatCompletionsRequestWithTopP(t *testing.T) {
	req := &request.ChatCompletionsRequest{
		Messages: []*request.Message{
			{
				Role:    request.RoleUser,
				Content: "Hello",
			},
		},
		Model:            deepseek.DEEPSEEK_CHAT_MODEL,
		FrequencyPenalty: 1,
		MaxTokens:        1,
		PresencePenalty:  1,
		ResponseFormat: &request.ResponseFormat{
			Type: request.ResponseFormatText,
		},
		Stop:   []string{"MOD1"},
		Stream: true,
		StreamOptions: &request.StreamOptions{
			IncludeUsage: true,
		},
		Temperature: 2.0,
		TopP:        float32Ptr(0.5),
	}
	err := request.ValidateChatCompletionsRequest(req)
	assert.NoError(t, err)
	fmt.Println(err)
}

func float32Ptr(f float32) *float32 {
	return &f
}
