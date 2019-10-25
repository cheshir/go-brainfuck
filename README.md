# Yet Another Brainfuck Compiler

Brainfuck compiler based on the YACC.

## Motivation

Just to try YACC tool.

## Usage

Install:

`go get -u github.com/cheshir/go-brainfuck`

Prepare input (brainfuck source code):

`echo "++.>.<.>." > input.bf`

Execute:

`go-brainfuck -s input.bf`

## TODO

* ASCII instead of raw bytes output.
* Loops
* Compiler (right now utility works as an interpreter)
