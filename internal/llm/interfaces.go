package llm

type Embedder interface {
	GetEmbedding(text string) ([]float32, error)
}
