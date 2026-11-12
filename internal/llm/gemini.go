package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type GeminiClient struct {
	APIKey string
	Model  string
}

func NewGeminiClient(apiKey, model string) *GeminiClient {
	return &GeminiClient{APIKey: apiKey, Model: model}
}

type geminiRequest struct {
	Content struct {
		Parts []struct {
			Text string `json:"text"`
		} `json:"parts"`
	} `json:"content"`
	OutputDimensionality int `json:"outputDimensionality"`
}

type geminiResponse struct {
	Embedding struct {
		Values []float32 `json:"values"`
	} `json:"embedding"`
}

func (c *GeminiClient) GetEmbedding(text string) ([]float32, error) {
	reqBody := geminiRequest{
		OutputDimensionality: 768, // Forces Gemini to output exactly 768 dimensions
	}
	reqBody.Content.Parts = append(reqBody.Content.Parts, struct {
		Text string `json:"text"`
	}{Text: text})

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models/%s:embedContent?key=%s", c.Model, c.APIKey)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("gemini api error: status %d", resp.StatusCode)
	}

	var result geminiResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if len(result.Embedding.Values) == 0 {
		return nil, fmt.Errorf("no embedding returned")
	}

	return result.Embedding.Values, nil
}
