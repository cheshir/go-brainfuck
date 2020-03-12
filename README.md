# Yet Another Brainfuck Compiler (Transpiler)

## About

There are two modes: transpiler from brainfuck source code to go and interpreter.

Implemented in raw go without yacc, lex and other stuff.

## Why?

For fun.

## Usage

First of all:

`go get -u github.com/cheshir/go-brainfuck`

### Brainfuck to go code  

Transpile bf file to go:

`go-brainfuck -f /path/to/source.bf -o ./compiled/compiled.go`

Or transpile a little piece of bf code:

`go-brainfuck -s "+>++>+++" -o ./compiled/compiled.go`

Then compile go file:

`go build ./compiled -o compiled-bf`

Run your bf program:

`./compiled-bf`

### Brainfuck interpreter

Also, you can simply run your brainfuck program without compiling.

From file:

`go-brainfuck -f /path/to/source.bf`

Without file:

`go-brainfuck -s "+>++>+++"`
