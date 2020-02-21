package vm

import (
	"log"
	"strings"

	"github.com/cheshir/go-brainfuck/lexer"
	"github.com/cheshir/go-brainfuck/parser"
)

func ExampleVirtualMachine_Run() {
	source := strings.NewReader("++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++.") // Hello World!
	tokens, err := lexer.Tokenize(source)
	if err != nil {
		log.Fatalf("Lexer error: %v", err)
	}

	operations, err := parser.Parse(tokens)
	if err != nil {
		log.Fatalf("Parser error: %v", err)
	}

	vm := New(operations)
	if err := vm.Run(); err != nil {
		log.Fatalf("VM error: %v", err)
	}

	// Output: Hello World!
}
