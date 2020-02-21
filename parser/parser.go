package parser

import (
	"errors"
	"fmt"

	"github.com/cheshir/go-brainfuck/lexer"
)

func Parse(tokens []lexer.Token) ([]*Operation, error) {
	tmpStacks := make([]*stack, 0) // For nested loops stacks.
	currentStack := newStack()

	for _, token := range tokens {
		if token == lexer.LoopStart {
			tmpStacks = append(tmpStacks, currentStack)
			currentStack = newStack()

			continue
		}

		op, err := newOperation(token)
		if err != nil {
			return nil, err
		}

		if op.Type == Loop {
			if len(tmpStacks) == 0 {
				return nil, errors.New("trying to close not started loop")
			}

			op.Value = currentStack.data
			// Return previous stack.
			currentStack = tmpStacks[len(tmpStacks)-1]
			tmpStacks = tmpStacks[:len(tmpStacks)-1]
		}

		currentStack.Push(op)
	}

	if len(tmpStacks) != 0 {
		return nil, fmt.Errorf("%d loop[s] was not closed", len(tmpStacks))
	}

	parsed := optimize(currentStack.data)

	return parsed, nil
}

// Just for example implemented only squashing same operations.
func optimize(operations []*Operation) []*Operation {
	result := newStack()
	optimizeNested(result, operations)

	return result.data
}

func optimizeNested(result *stack, operations []*Operation) {
	for _, op := range operations {
		prev, _ := result.Top()
		if prev == nil {
			result.Push(op)
			continue
		}

		if op.Type == prev.Type {
			switch op.Type {
			case Shift, Add:
				// Squashing.
				prev.Value = prev.Value.(int) + op.Value.(int)

				// Annihilation.
				if prev.Value.(int) == 0 {
					_, _ = result.Pop()
				}

				continue
			}
		}

		if op.Type == Loop {
			newBody := newStack()
			optimizeNested(newBody, op.Value.([]*Operation))
			op.Value = newBody.data
		}

		result.Push(op)
	}
}
