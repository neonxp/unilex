package unilex

// StateFunc represents function that scans lexems and returns new state function or nil if lexing completed.
type StateFunc func(*Lexer) StateFunc

type stateStack []StateFunc

func (ss *stateStack) Push(s StateFunc) {
	*ss = append(*ss, s)
}

func (ss *stateStack) Pop() (s StateFunc) {
	if len(*ss) == 0 {
		return nil
	}
	*ss, s = (*ss)[:len(*ss)-1], (*ss)[len(*ss)-1]
	return s
}
