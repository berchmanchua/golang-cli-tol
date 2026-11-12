package cli

import (
	"flag"
	"fmt"

	"github.com/berchmanchua/golang-cli-tol/internal/db"
	"github.com/berchmanchua/golang-cli-tol/internal/llm"
)

func handleSave(args []string) error {
	saveCmd := flag.NewFlagSet("save", flag.ExitOnError)
	cmdText := saveCmd.String("cmd", "", "The command to save")
	descText := saveCmd.String("desc", "", "Description of the command")

	saveCmd.Parse(args)

	if *cmdText == "" || *descText == "" {
		return fmt.Errorf("both --cmd and --desc are required")
	}

	fmt.Println("Initializing AI provider...")
	embedder, err := llm.NewEmbedderFromEnv()
	if err != nil {
		return fmt.Errorf("failed to initialize AI provider: %v", err)
	}

	fmt.Printf("Generating embeddings for description: %s\n", *descText)
	embedding, err := embedder.GetEmbedding(*descText)
	if err != nil {
		return fmt.Errorf("failed to get embedding: %v", err)
	}

	fmt.Println("Connecting to SQLite database...")
	database, err := db.NewDatabase("tol.db")
	if err != nil {
		return err
	}
	defer database.Close()

	fmt.Printf("Saving command: %s\n", *cmdText)
	err = database.InsertCommand(*cmdText, *descText, embedding)
	if err != nil {
		return fmt.Errorf("failed to save command: %v", err)
	}

	fmt.Println("Successfully saved to tol database!")
	return nil
}
