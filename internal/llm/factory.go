package llm

import (
	"fmt"
	"os"
)

// NewEmbedderFromEnv initializes the correct LLM provider based on environment variables.
func NewEmbedderFromEnv() (Embedder, error) {
	provider := os.Getenv("TOL_AI_PROVIDER")
	if provider == "" {
		provider = "ollama" // default
	}

	switch provider {
	case "gemini":
		apiKey := os.Getenv("GEMINI_API_KEY")
		if apiKey == "" {
			return nil, fmt.Errorf("GEMINI_API_KEY environment variable is required")
		}
		model := os.Getenv("GEMINI_MODEL")
		if model == "" {
			model = "text-embedding-004"
		}
		return NewGeminiClient(apiKey, model), nil

	case "claude", "anthropic", "voyage":
		// Anthropic officially partners with Voyage AI for embeddings as they do not host their own.
		apiKey := os.Getenv("VOYAGE_API_KEY") 
		if apiKey == "" {
			return nil, fmt.Errorf("VOYAGE_API_KEY environment variable is required for Claude/Anthropic ecosystem embeddings")
		}
		model := os.Getenv("VOYAGE_MODEL")
		if model == "" {
			model = "voyage-3"
		}
		return NewClaudeVoyageClient(apiKey, model), nil

	case "openai":
		apiKey := os.Getenv("OPENAI_API_KEY")
		if apiKey == "" {
			return nil, fmt.Errorf("OPENAI_API_KEY environment variable is required")
		}
		model := os.Getenv("OPENAI_MODEL")
		if model == "" {
			model = "text-embedding-3-small"
		}
		return NewOpenAIClient(apiKey, model), nil

	case "ollama":
		host := os.Getenv("OLLAMA_HOST")
		if host == "" {
			host = "http://localhost:11434"
		}
		model := os.Getenv("OLLAMA_MODEL")
		if model == "" {
			model = "nomic-embed-text"
		}
		return NewOllamaClient(host, model), nil

	default:
		return nil, fmt.Errorf("unsupported AI provider: %s", provider)
	}
}
