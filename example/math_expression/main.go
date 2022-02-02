package main

import (
	"fmt"

	"github.com/neonxp/unilex"
)

const (
	LP unilex.LexType = iota
	RP
	NUMBER
	OPERATOR
)

func main() {

	l := unilex.New("10 * (20.0 + 30.0)")

	go l.Run(lexExpression) // Start lexer

	// Read infix expression lexems from lexer and convert them to RPN (reverse polish notation)
	rpn := infixToRPNotation(l)
	fmt.Println("RPN:", rpn)

	// Calculate RPN
	result := calculateRPN(rpn)
	fmt.Println("Result:", result)
}

func lexExpression(l *unilex.Lexer) unilex.StateFunc {
	l.AcceptWhile(" \t")
	l.Ignore() // Ignore whitespaces

	switch {
	case l.Accept("("):
		l.Emit(LP)
	case l.Accept(")"):
		l.Emit(RP)
	case unilex.ScanNumber(l):
		l.Emit(NUMBER)
	case l.Accept("+-*/^!"):
		l.Emit(OPERATOR)
	case l.Peek() == unilex.EOF:
		return nil
	default:
		return l.Errorf("Unexpected symbol")
	}

	return lexExpression
}
