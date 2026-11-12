package cli

import (
	"fmt"
	"os"
)

func Execute() error {
	if len(os.Args) < 2 {
		return fmt.Errorf("expected 'save', 'find', or 'serve-mcp' subcommands")
	}

	switch os.Args[1] {
	case "save":
		return handleSave(os.Args[2:])
	case "find":
		return handleFind(os.Args[2:])
	default:
		return fmt.Errorf("unknown subcommand: %s", os.Args[1])
	}
}
