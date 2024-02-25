package main

import (
	"context"
	openai "github.com/sashabaranov/go-openai"
)

func fetchChatCompletion(prompt string) (openai.ChatCompletionResponse, error) {
	client := openai.NewClient(Config.OpenAIKey)

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)
	return resp, err
}
