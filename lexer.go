package unilex

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// EOF const.
const EOF rune = -1

// Lexer holds current scanner state.
type Lexer struct {
	Input  string     // Input string.
	Start  int        // Start position of current lexem.
	Pos    int        // Pos at input string.
	Output chan Lexem // Lexems channel.
	width  int        // Width of last rune.
	states stateStack // Stack of states to realize PrevState.
}

// New returns new scanner for input string.
func New(input string) *Lexer {
	return &Lexer{
		Input:  input,
		Start:  0,
		Pos:    0,
		Output: make(chan Lexem, 2),
		width:  0,
	}
}

// Run lexing.
func (l *Lexer) Run(init StateFunc) {
	for state := init; state != nil; {
		state = state(l)
	}
	close(l.Output)
}

// PopState returns previous state function.
func (l *Lexer) PopState() StateFunc {
	return l.states.Pop()
}

// PushState pushes state before going deeper states.
func (l *Lexer) PushState(s StateFunc) {
	l.states.Push(s)
}

// Emit current lexem to output.
func (l *Lexer) Emit(typ LexType) {
	l.Output <- Lexem{
		Type:  typ,
		Value: l.Input[l.Start:l.Pos],
		Start: l.Start,
		End:   l.Pos,
	}
	l.Start = l.Pos
}

// Errorf produces error lexem and stops scanning.
func (l *Lexer) Errorf(format string, args ...interface{}) StateFunc {
	l.Output <- Lexem{
		Type:  LexError,
		Value: fmt.Sprintf(format, args...),
		Start: l.Start,
		End:   l.Pos,
	}
	return nil
}

// Next rune from input.
func (l *Lexer) Next() (r rune) {
	if int(l.Pos) >= len(l.Input) {
		l.width = 0
		return EOF
	}
	r, l.width = utf8.DecodeRuneInString(l.Input[l.Pos:])
	l.Pos += l.width
	return r
}

// Back move position to previos rune.
func (l *Lexer) Back() {
	l.Pos -= l.width
}

// Ignore previosly buffered text.
func (l *Lexer) Ignore() {
	l.Start = l.Pos
	l.width = 0
}

// Peek rune at current position without moving position.
func (l *Lexer) Peek() (r rune) {
	r = l.Next()
	l.Back()
	return r
}

// Accept any rune from valid string. Returns true if next rune was in valid string.
func (l *Lexer) Accept(valid string) bool {
	if strings.ContainsRune(valid, l.Next()) {
		return true
	}
	l.Back()
	return false
}

// AcceptString returns true if given string was at position.
func (l *Lexer) AcceptString(s string, caseInsentive bool) bool {
	input := l.Input
	if caseInsentive {
		input = strings.ToLower(input)
		s = strings.ToLower(s)
	}
	if strings.HasPrefix(input, s) {
		l.width = 0
		l.Pos += len(s)
		return true
	}
	return false
}

// AcceptAnyOf substrings. Retuns true if any of substrings was found.
func (l *Lexer) AcceptAnyOf(s []string, caseInsentive bool) bool {
	for _, substring := range s {
		if l.AcceptString(substring, caseInsentive) {
			return true
		}
	}
	return false
}

// AcceptWhile passing symbols from input while they at `valid` string.
func (l *Lexer) AcceptWhile(valid string) {
	for l.Accept(valid) {
	}
}

// AcceptWhileNot passing symbols from input while they NOT in `invalid` string.
func (l *Lexer) AcceptWhileNot(invalid string) {
	for !strings.ContainsRune(invalid, l.Next()) {
	}
	l.Back()
}

// AtStart returns true if current lexem not empty
func (l *Lexer) AtStart() bool {
	return l.Pos == l.Start
}
