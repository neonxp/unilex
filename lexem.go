package unilex

// Lexem represents part of parsed string.
type Lexem struct {
	Type  LexType // Type of Lexem.
	Value string  // Value of Lexem.
	Start int     // Start position at input string.
	End   int     // End position at input string.
}

// LexType represents type of current lexem.
type LexType string

// Some std lexem types
const (
	// LError represents lexing error.
	LError LexType = "ERROR"
	// LEOF represents end of input.
	LEOF LexType = "EOF"
)
