package unilex

// ScanNumber simplest scanner that accepts decimal int and float.
func ScanNumber(l *Lexer) bool {
	l.AcceptWhile("0123456789")
	if l.AtStart() {
		// not found any digit
		return false
	}
	l.Accept(".")
	l.AcceptWhile("0123456789")
	return !l.AtStart()
}

// ScanAlphaNum returns true if next input token contains alphanum sequence that not starts from digit and not contains.
// spaces or special characters.
func ScanAlphaNum(l *Lexer) bool {
	digits := "0123456789"
	alpha := "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"
	if !l.Accept(alpha) {
		return false
	}
	l.AcceptWhile(alpha + digits)
	return true
}

// ScanQuotedString returns true if next input tokens is quoted string. Can be used with any type of quotes.
func ScanQuotedString(l *Lexer, quote rune) bool {
	start := l.Pos
	if l.Next() != quote {
		l.Back()
		return false
	}
	for {
		ch := l.Next()
		switch ch {
		case EOF:
			l.Pos = start // Return position to start
			return false  // Unclosed quote string?
		case '\\':
			l.Next() // Skip next char
		case quote:
			return true // Closing quote
		}
	}
}
