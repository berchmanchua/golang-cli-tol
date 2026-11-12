package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// ClaudeVoyageClient uses Voyage AI, Anthropic's official embedding partner.
type ClaudeVoyageClient struct {
	APIKey string
	Model  string
}

func NewClaudeVoyageClient(apiKey, model string) *ClaudeVoyageClient {
	return &ClaudeVoyageClient{APIKey: apiKey, Model: model}
}

type voyageRequest struct {
	Input           []string `json:"input"`
	Model           string   `json:"model"`
	OutputDimension int      `json:"output_dimension"`
}

type voyageResponse struct {
	Data []struct {
		Embedding []float32 `json:"embedding"`
	} `json:"data"`
}

func (c *ClaudeVoyageClient) GetEmbedding(text string) ([]float32, error) {
	reqBody := voyageRequest{
		Input:           []string{text},
		Model:           c.Model,
		OutputDimension: 768, // Voyage 3 supports Matryoshka truncation to match our 768 baseline
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "https://api.voyageai.com/v1/embeddings", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.APIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("voyage ai (claude partner) error: status %d", resp.StatusCode)
	}

	var result voyageResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("no embedding returned from Voyage AI")
	}

	return result.Data[0].Embedding, nil
}
