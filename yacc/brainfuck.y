%{

package main

import (
	"flag"
	"io/ioutil"
	"log"
	"unicode"
)

const heapSize = 4096

var heap = make([]int, heapSize, heapSize)
var cursor = 0
var sourceFile string

func init() {
	flag.StringVar(&sourceFile, "s", "", "Path to the brainfuck source file")
	flag.Parse()
}

%}

%union{}

%token '>' '<' '+' '-' '.' SPACE COMMENT

%%

line	: /* empty */
	| line expr
	| line COMMENT
	| line SPACE;

expr	:    '>'
		{
			if cursor < heapSize - 1 {
				cursor++
			}
		}
	|    '<'
		{
			if cursor >  0 {
				cursor--
			}
		}
	|    '+'
		{
			if heap[cursor] < 255 {
				heap[cursor]++
			}
		}
	|    '-'
		{
			if heap[cursor] > 0 {
				heap[cursor]--
			}
		}
	|    '.'
		{
			print(heap[cursor]);
		}
	;
%%

func main() {
	if sourceFile == "" {
		log.Fatal("Source file must be specified. Check help: -h")
	}

	source, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		log.Fatal("Failed to read source file: ", err.Error())
	}

	yyParse(&yyLex{source: source})
}

type yyLex struct {
	source []byte
	cursor int
}

var isComment bool

func (l *yyLex) Lex(lval *yySymType) int {
	if l.cursor == len(l.source) {
		return 0
	}

	c := rune(l.source[l.cursor])
	l.cursor++

	switch c {
	case '#':
		isComment = true
	case '\n':
		isComment = false
	}

	if isComment {
		return COMMENT
	}

	if unicode.IsSpace(c) {
		return SPACE
	}

	return int(c)
}

func (l *yyLex) Error(s string) {
	log.Printf("syntax error: %s\n", s)
}
