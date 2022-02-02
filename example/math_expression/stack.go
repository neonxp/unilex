package main

// Simple lexem stack implementation.

import "github.com/neonxp/unilex"

type lexemStack []unilex.Lexem

func (ls *lexemStack) Head() (l unilex.Lexem) {
	if len(*ls) == 0 {
		return unilex.Lexem{Type: unilex.LexEOF}
	}
	return (*ls)[len(*ls)-1]
}

func (ls *lexemStack) Push(l unilex.Lexem) {
	*ls = append(*ls, l)
}

func (ls *lexemStack) Pop() (l unilex.Lexem) {
	if len(*ls) == 0 {
		return unilex.Lexem{Type: unilex.LexEOF}
	}
	*ls, l = (*ls)[:len(*ls)-1], (*ls)[len(*ls)-1]
	return l
}
