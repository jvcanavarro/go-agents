package brain

import (
	"context"
	"database/sql"
	"log"

	"google.golang.org/genai"
)

type Agent struct {
	client *genai.Client
	db     *sql.DB
	model  string
	config *genai.GenerateContentConfig
	memory []*genai.Content
	Debug  bool
}

func (a *Agent) SendMessage(ctx context.Context, input string) (*genai.GenerateContentResponse, error) {
	a.memory = append(a.memory, genai.NewContentFromText(input, genai.RoleUser))
	resp, err := a.client.Models.GenerateContent(ctx, a.model, a.memory, a.config)
	if err != nil {
		return nil, err
	}

	a.memory = append(a.memory, genai.NewContentFromText(resp.Text(), genai.RoleModel))
	return resp, nil
}

func (a *Agent) LogDebug(msg ...any) {
	if a.Debug {
		// a.LogDebug("teste", "valor1", "")
		log.Println("Bom dia Mestre,")
		log.Println(msg...)
	}
}
