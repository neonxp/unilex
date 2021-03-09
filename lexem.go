package unilex

// Lexem represents part of parsed string.
type Lexem struct {
	Type  LexType // Type of Lexem.
	Value string  // Value of Lexem.
	Start int     // Start position at input string.
	End   int     // End position at input string.
}

// LexType represents type of current lexem.
type LexType int

// Some std lexem types
const (
	// LEOF represents end of input.
	LexEOF LexType = -1
	// LError represents lexing error.
	LexError LexType = -2
)
