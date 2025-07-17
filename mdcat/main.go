package main

import (
	"fmt"
	"io"
	"os"

	"github.com/charmbracelet/glamour"
)

func mdPrint(r io.Reader) error {
	content, err := io.ReadAll(r)
	if err != nil {
		return fmt.Errorf("failed to read input: %w", err)
	}

	renderer, err := glamour.NewTermRenderer(glamour.WithAutoStyle())
	if err != nil {
		return fmt.Errorf("failed to create renderer: %w", err)
	}

	out, err := renderer.Render(string(content))
	if err != nil {
		return fmt.Errorf("failed to render markdown: %w", err)
	}

	fmt.Print(out)
	return nil
}

func main() {
	args := os.Args[1:]

	// If no files provided, read from stdin
	if len(args) == 0 {
		if err := mdPrint(os.Stdin); err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}
		return
	}

	for _, filename := range args {
		var file io.Reader
		if filename == "-" {
			file = os.Stdin
		} else {
			f, err := os.Open(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error opening file %s: %v\n", filename, err)
				continue
			}
			defer f.Close()
			file = f
		}

		if err := mdPrint(file); err != nil {
			fmt.Fprintf(os.Stderr, "Error rendering %s: %v\n", filename, err)
		}
	}
}
