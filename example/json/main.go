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
		"key4": 123.321,
		"key5": [
			1,
			2,
			[
				3,
				4,
				5,
				{
					"key6": "value6"
				}
			]
		]
	}`
	l := unilex.New(testJson)
	go l.Run(initJson)
	for ll := range l.Output {
		fmt.Println(ll)
	}
}

const (
	lObjectStart unilex.LexType = iota
	lObjectEnd
	lObjectKey
	lObjectValue
	lArrayStart
	lArrayEnd
	lString
	lNumber
)

func initJson(l *unilex.Lexer) unilex.StateFunc {
	ignoreWhiteSpace(l)
	switch {
	case l.Accept("{"):
		l.Emit(lObjectStart)
		return stateInObject
	case l.Peek() == unilex.EOF:
		return nil
	}
	return l.Errorf("Unknown token: %s", string(l.Peek()))
}

func stateInObject(l *unilex.Lexer) unilex.StateFunc {
	// we in object, so we expect field keys and values
	ignoreWhiteSpace(l)
	if l.Accept("}") {
		l.Emit(lObjectEnd)
		// If meet close object return to previous state (including initial)
		return l.PopState()
	}
	ignoreWhiteSpace(l)
	l.Accept(",")
	ignoreWhiteSpace(l)
	if !unilex.ScanQuotedString(l, '"') {
		return l.Errorf("Unknown token: %s", string(l.Peek()))
	}
	l.Emit(lObjectKey)
	ignoreWhiteSpace(l)
	if !l.Accept(":") {
		return l.Errorf("Expected ':'")
	}
	ignoreWhiteSpace(l)
	l.Emit(lObjectValue)
	switch {
	case unilex.ScanQuotedString(l, '"'):
		l.Emit(lString)
		ignoreWhiteSpace(l)
		l.Accept(",")
		l.Ignore()
		ignoreWhiteSpace(l)
		return stateInObject
	case unilex.ScanNumber(l):
		l.Emit(lNumber)
		ignoreWhiteSpace(l)
		l.Accept(",")
		l.Ignore()
		ignoreWhiteSpace(l)
		return stateInObject
	case l.Accept("{"):
		l.Emit(lObjectStart)
		l.PushState(stateInObject)
		return stateInObject
	case l.Accept("["):
		l.Emit(lArrayStart)
		l.PushState(stateInObject)
		return stateInArray
	}
	return l.Errorf("Unknown token: %s", string(l.Peek()))
}

func stateInArray(l *unilex.Lexer) unilex.StateFunc {
	ignoreWhiteSpace(l)
	l.Accept(",")
	ignoreWhiteSpace(l)
	switch {
	case unilex.ScanQuotedString(l, '"'):
		l.Emit(lString)
	case unilex.ScanNumber(l):
		l.Emit(lNumber)
	case l.Accept("{"):
		l.Emit(lObjectStart)
		l.PushState(stateInArray)
		return stateInObject
	case l.Accept("["):
		l.Emit(lArrayStart)
		l.PushState(stateInArray)
		return stateInArray
	case l.Accept("]"):
		l.Emit(lArrayEnd)
		return l.PopState()
	}
	return stateInArray
}

func ignoreWhiteSpace(l *unilex.Lexer) {
	l.AcceptWhile(" \n\t") //ignore whitespaces
	l.Ignore()
}
