package unilex

// StateFunc represents function that scans lexems and returns new state function or nil if lexing completed.
type StateFunc func(*Lexer) StateFunc
