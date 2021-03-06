package main

// Helper functions to convert infix notation to RPN and calculates expression result.

import (
	"fmt"
	"log"
	"strconv"

	"github.com/neonxp/unilex"
)

func infixToRPNotation(l *unilex.Lexer) []unilex.Lexem {
	output := []unilex.Lexem{}
	stack := lexemStack{}
parseLoop:
	for ll := range l.Output { // Read lexems from Lexer output channel, convert starts as soon as first lexems scanned!
		fmt.Printf("Lexem: %+v\n", ll)

		switch {
		case ll.Type == "NUMBER", ll.Type == "OPERATOR" && ll.Value == "!":
			output = append(output, ll)
		case ll.Type == "LP":
			stack.Push(ll)
		case ll.Type == "RP":
			for {
				cl := stack.Pop()
				if cl.Type == "LP" {
					break
				}
				if cl.Type == unilex.LEOF {
					log.Fatalf("No pair for parenthesis at %d", ll.Start)
				}
				output = append(output, cl)
			}
		case ll.Type == "OPERATOR":
			for {
				if stack.Head().Type == "OPERATOR" && (opPriority[stack.Head().Value] > opPriority[ll.Value]) {
					output = append(output, stack.Pop())
					continue
				}
				break
			}
			stack.Push(ll)
		case ll.Type == unilex.LEOF:
			break parseLoop
		}
	}

	for stack.Head().Type != unilex.LEOF {
		output = append(output, stack.Pop())
	}

	return output
}

func calculateRPN(rpnLexems []unilex.Lexem) string {
	stack := lexemStack{}
	for _, op := range rpnLexems {
		if op.Type == "NUMBER" {
			stack.Push(op)
		} else {
			switch op.Value {
			case "+":
				a1, _ := strconv.ParseFloat(stack.Pop().Value, 64)
				a2, _ := strconv.ParseFloat(stack.Pop().Value, 64)
				stack.Push(unilex.Lexem{Type: "NUMBER", Value: strconv.FormatFloat(a2+a1, 'f', -1, 64)})
			case "-":
				a1, _ := strconv.ParseFloat(stack.Pop().Value, 64)
				a2, _ := strconv.ParseFloat(stack.Pop().Value, 64)
				stack.Push(unilex.Lexem{Type: "NUMBER", Value: strconv.FormatFloat(a2-a1, 'f', -1, 64)})
			case "*":
				a1, _ := strconv.ParseFloat(stack.Pop().Value, 64)
				a2, _ := strconv.ParseFloat(stack.Pop().Value, 64)
				stack.Push(unilex.Lexem{Type: "NUMBER", Value: strconv.FormatFloat(a2*a1, 'f', -1, 64)})
			case "/":
				a1, _ := strconv.ParseFloat(stack.Pop().Value, 64)
				a2, _ := strconv.ParseFloat(stack.Pop().Value, 64)
				stack.Push(unilex.Lexem{Type: "NUMBER", Value: strconv.FormatFloat(a2/a1, 'f', -1, 64)})
			}
		}
	}
	return stack.Head().Value
}
