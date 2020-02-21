package main

import (
	"bytes"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"strings"

	"github.com/cheshir/go-brainfuck/lexer"
	"github.com/cheshir/go-brainfuck/parser"
	"github.com/cheshir/go-brainfuck/vm"
)

var sourceFile string
var source string
var outputFile string

func init() {
	flag.StringVar(&source, "s", "", "Brainfuck source")
	flag.StringVar(&sourceFile, "f", "", "Path to brainfuck source file")
	flag.StringVar(&outputFile, "o", "", "Path to store generated code")
	flag.Parse()
}

func main() {
	var input io.RuneScanner

	if source != "" {
		input = strings.NewReader(source)
	}

	if sourceFile != "" {
		sourceCode, err := ioutil.ReadFile(sourceFile)
		if err != nil {
			log.Fatal(err)
		}

		input = bytes.NewReader(sourceCode)
	}

	tokens, err := lexer.Tokenize(input)
	if err != nil {
		log.Fatal(err)
	}

	parsed, err := parser.Parse(tokens)
	if err != nil {
		log.Fatal(err)
	}

	if outputFile != "" {
		println("Code generation is not implemented yet. Switching to interpretation mode...")
	}

	program := vm.New(parsed)
	if err := program.Run(); err != nil {
		log.Fatal(err)
	}
}
