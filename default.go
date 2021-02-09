// Package reggen generates text based on regex definitions
package reggen

const defaultLimit = 258

// Generate string by regexp
func Generate(regex string) (string, error) {
	g, err := NewGenerator(regex)
	if err != nil {
		return "", err
	}
	return g.Generate(defaultLimit), nil
}
