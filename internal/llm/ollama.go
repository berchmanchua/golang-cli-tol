package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type OllamaClient struct {
	Host  string
	Model string
}

func NewOllamaClient(host, model string) *OllamaClient {
	return &OllamaClient{Host: host, Model: model}
}

type ollamaEmbedRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
}

type ollamaEmbedResponse struct {
	Embedding []float32 `json:"embedding"`
}

func (c *OllamaClient) GetEmbedding(text string) ([]float32, error) {
	reqBody := ollamaEmbedRequest{
		Model:  c.Model,
		Prompt: text,
	}
	
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/api/embeddings", c.Host)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	var result ollamaEmbedResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Embedding, nil
}
