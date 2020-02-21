package parser

import (
	"github.com/cheshir/go-brainfuck/lexer"
	"github.com/pkg/errors"
)

type OperationType uint

const (
	Shift OperationType = iota + 1
	Add
	GetChar
	PutChar
	Loop
)

type Operation struct {
	Type  OperationType
	Value interface{}
}

func newOperation(token lexer.Token) (*Operation, error) {
	switch token {
	case lexer.ShiftLeft:
		return &Operation{
			Type:  Shift,
			Value: -1,
		}, nil
	case lexer.ShiftRight:
		return &Operation{
			Type:  Shift,
			Value: 1,
		}, nil
	case lexer.Increment:
		return &Operation{
			Type:  Add,
			Value: 1,
		}, nil
	case lexer.Decrement:
		return &Operation{
			Type:  Add,
			Value: -1,
		}, nil
	case lexer.PutChar:
		return &Operation{
			Type: PutChar,
		}, nil
	case lexer.GetChar:
		return &Operation{
			Type: GetChar,
		}, nil
	case lexer.LoopEnd:
		return &Operation{
			Type: Loop,
		}, nil
	}

	return nil, errors.Errorf("got unknown token %v", token)
}
