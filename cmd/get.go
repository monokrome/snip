package cmd

import (
	"fmt"

	"github.com/matheuzgomes/Snip/internal/handler"
	"github.com/spf13/cobra"
)

var render bool
var jsonOutput bool

func init() {
	showCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Show more information about the notes")
	showCmd.Flags().BoolVarP(&render, "render", "r", false, "Render the note markdown")
	showCmd.Flags().BoolVarP(&jsonOutput, "json", "j", false, "Output JSONL with metadata")
}

var showCmd = &cobra.Command{
	Use:   "show [id]",
	Short: "Display the content of a specific note by ID",
	Long: `Display the full content of a specific note identified by its unique ID.

This command shows the note's title, content and tags in a readable format. Use the verbose
flag to see additional metadata like creation and modification timestamps.

Flags:
  --verbose, -v  Show detailed metadata (timestamps, ID, etc.)
  --render, -r   Render the note markdown content (default is false)
  --json, -j     Output JSONL with metadata (omit id for all notes)

Examples:
  snip show 1              # Display note with ID 1
  snip show 42             # Display note with ID 42  
  snip show 1 --verbose    # Show note 1 with full metadata
  snip show 1 -v           # Same as above (short flag)
  snip show 1 -r           # Render note 1 markdown content
  snip show --json         # All notes as JSONL
  snip show --json 10      # Single note as JSONL`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := executeWithHandler(func(h handler.Handler) error {
			idStr := ""
			if len(args) > 0 {
				idStr = args[0]
			}
			return h.GetNote(idStr, verbose, render, jsonOutput)
		}); err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	},
}
