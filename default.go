// Package reggen generates text based on regex definitions
package reggen

import "fmt"

const defaultLimit = 258

// Generate string by regexp.
func Generate(regex string) (string, error) {
	g, err := NewGenerator(regex)
	if err != nil {
		return "", fmt.Errorf("failed to create generator: %w", err)
	}

	return g.Generate(defaultLimit), nil
}
