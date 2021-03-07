// +build ignore

package main

import (
	"fmt"

	"github.com/neonxp/unilex"
)

func main() {
	testJson := `
	{
		"key1": "value1",
		"key2": {
			"key3" : "value 3"
		},
		"key4": 123.321
	}`
	l := unilex.New(testJson)
	go l.Run(initJson)
	for ll := range l.Output {
		fmt.Println(ll)
	}
}

const (
	lObjectStart       unilex.LexType = "lObjectStart"
	lObjectEnd         unilex.LexType = "lObjectEnd"
	lObjectKey         unilex.LexType = "lObjectKey"
	lObjectValueString unilex.LexType = "lObjectValueString"
	lObjectValueNumber unilex.LexType = "lObjectValueNumber"
)

func initJson(l *unilex.Lexer) unilex.StateFunc {
	ignoreWhiteSpace(l)
	switch {
	case l.Accept("{"):
		l.Emit(lObjectStart)
		return stateInObject(true)
	case l.Peek() == unilex.EOF:
		return nil
	}
	return l.Errorf("Unknown token: %s", l.Peek())
}

func stateInObject(initial bool) unilex.StateFunc {
	return func(l *unilex.Lexer) unilex.StateFunc {
		// we in object, so we expect field keys and values
		ignoreWhiteSpace(l)
		if l.Accept("}") {
			l.Emit(lObjectEnd)
			if initial {
				return initJson
			}
			ignoreWhiteSpace(l)
			l.Accept(",")
			ignoreWhiteSpace(l)
			return stateInObject(initial)
		}
		if l.Peek() == unilex.EOF {
			return nil
		}
		if !unilex.ScanQuotedString(l, '"') {
			return l.Errorf("Unknown token: %s", l.Peek())
		}
		l.Emit(lObjectKey)
		ignoreWhiteSpace(l)
		if !l.Accept(":") {
			return l.Errorf("Expected ':'")
		}
		ignoreWhiteSpace(l)
		switch {
		case unilex.ScanQuotedString(l, '"'):
			l.Emit(lObjectValueString)
			ignoreWhiteSpace(l)
			l.Accept(",")
			l.Ignore()
			ignoreWhiteSpace(l)
			return stateInObject(initial)
		case unilex.ScanNumber(l):
			l.Emit(lObjectValueNumber)
			ignoreWhiteSpace(l)
			l.Accept(",")
			l.Ignore()
			ignoreWhiteSpace(l)
			return stateInObject(initial)
		case l.Accept("{"):
			l.Emit(lObjectStart)
			return stateInObject(false)
		}
		return l.Errorf("Unknown token")
	}
}

func ignoreWhiteSpace(l *unilex.Lexer) {
	l.AcceptWhile(" \n\t") //ignore whitespaces
	l.Ignore()
}
