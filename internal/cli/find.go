package cli

import (
	"fmt"
	"strings"

	"github.com/berchmanchua/golang-cli-tol/internal/db"
	"github.com/berchmanchua/golang-cli-tol/internal/llm"
)

func handleFind(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("expected search query")
	}
	query := strings.Join(args, " ")

	embedder, err := llm.NewEmbedderFromEnv()
	if err != nil {
		return fmt.Errorf("failed to initialize AI provider: %v", err)
	}

	fmt.Printf("Searching for: %s\n", query)
	embedding, err := embedder.GetEmbedding(query)
	if err != nil {
		return fmt.Errorf("failed to get embedding: %v", err)
	}

	database, err := db.NewDatabase("tol.db")
	if err != nil {
		return err
	}
	defer database.Close()

	if (embedding == nil) || (len(embedding) == 0) {
		return fmt.Errorf("embedding is empty")
	}

	result, err := database.SearchKNN(embedding, 1)
	if err != nil {
		return fmt.Errorf("search failed: %v", err)
	}

	if result.Command == "" {
		fmt.Println("No matches found.")
		return nil
	}

	fmt.Printf("\nMatch found:\nDescription: %s\nCommand: %s\n", result.Description, result.Command)
	return nil
}
