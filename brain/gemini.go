package brain

import (
	"context"
	"os"

	"google.golang.org/genai"
)

func NewAgent(ctx context.Context) (*Agent, error) {
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey: os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})

	if err != nil {
		return nil, err
	}

	return &Agent{
		client: client,
		model : "gemini-2.0-flash",
		config: &genai.GenerateContentConfig{
			SystemInstruction: genai.NewContentFromText(agentPrompt, genai.RoleModel),
		},
		memory: []*genai.Content{},
	} , nil
}