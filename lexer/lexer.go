package lexer

import (
	"io"
)

type Token int

const (
	ShiftLeft  Token = iota + 1 // <
	ShiftRight                  // >
	Increment                   // +
	Decrement                   // -
	GetChar                     // ,
	PutChar                     // .
	LoopStart                   // [
	LoopEnd                     // ]
)

var symbolToTokensMap = map[rune]Token{
	'<': ShiftLeft,
	'>': ShiftRight,
	'+': Increment,
	'-': Decrement,
	',': GetChar,
	'.': PutChar,
	'[': LoopStart,
	']': LoopEnd,
}

func Tokenize(input io.RuneScanner) ([]Token, error) {
	var tokens []Token

	for {
		symbol, _, err := input.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}

			return nil, err
		}

		token, ok := symbolToTokensMap[symbol]
		if !ok {
			continue // Skip unknown symbols.
		}

		tokens = append(tokens, token)
	}

	return tokens, nil
}
