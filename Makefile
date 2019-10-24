compile:
	goyacc -o main.go brainfuck.y && go build
