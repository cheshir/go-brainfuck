package lexer

import (
	"strings"
	"testing"
)

func TestAnalyze(t *testing.T) {
	tt := []struct {
		Input    string
		Expected []Token
	}{
		{
			Input:    "><",
			Expected: []Token{ShiftRight, ShiftLeft},
		},
		{
			Input: "",
		},
		{
			Input: `
,  # Read input
++ # Add 2
.  # Output result
`,
			Expected: []Token{GetChar, Increment, Increment, PutChar},
		},
	}

	for _, tc := range tt {
		input := strings.NewReader(tc.Input)
		tokens, _ := Tokenize(input)

		if len(tc.Expected) != len(tokens) {
			t.Errorf("%s Actual result is not equal to expected: \nActual: %v\nExpected:%v", tc.Input, tokens, tc.Expected)
		}

		for i, tok := range tokens {
			if tok != tc.Expected[i] {
				t.Errorf("%s Actual result is not equal to expected: \nActual: %v\nExpected:%v", tc.Input, tokens, tc.Expected)
			}
		}
	}
}
