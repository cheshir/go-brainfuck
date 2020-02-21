package vm

import (
	"fmt"
	"os"

	"github.com/cheshir/go-brainfuck/parser"
)

const heapSize = 30000

type VirtualMachine struct {
	operations []*parser.Operation
	memory     []byte
	position   int
}

func New(operations []*parser.Operation) *VirtualMachine {
	return &VirtualMachine{
		operations: operations,
		memory:     make([]byte, heapSize),
	}
}

func (vm *VirtualMachine) Run() error {
	err := vm.run(vm.operations)
	println() // Add new line to output.

	return err
}

func (vm *VirtualMachine) run(operations []*parser.Operation) error {
	for _, op := range operations {
		if err := vm.exec(op); err != nil {
			return err
		}
	}

	return nil
}

func (vm *VirtualMachine) exec(op *parser.Operation) error {
	switch op.Type {
	case parser.Shift:
		return vm.execShift(op)
	case parser.Add:
		return vm.execAdd(op)
	case parser.GetChar:
		return vm.execGetChar(op)
	case parser.PutChar:
		return vm.execPutChar(op)
	case parser.Loop:
		return vm.execLoop(op)
	}

	return fmt.Errorf("unexpected operation type %v", op.Type)
}

func (vm *VirtualMachine) execShift(op *parser.Operation) error {
	vm.position += op.Value.(int)
	if vm.position < 0 || vm.position >= len(vm.memory) {
		return fmt.Errorf("index %d is out of memory", vm.position)
	}

	return nil
}

func (vm *VirtualMachine) execAdd(op *parser.Operation) error {
	vm.memory[vm.position] += byte(op.Value.(int))

	return nil
}

func (vm *VirtualMachine) execGetChar(op *parser.Operation) error {
	print("(enter char)> ")

	input := make([]byte, 1)
	_, err := os.Stdin.Read(input)
	if err != nil {
		return err
	}

	vm.memory[vm.position] = input[0]

	return nil
}

func (vm *VirtualMachine) execPutChar(_ *parser.Operation) error {
	fmt.Print(string([]byte{vm.memory[vm.position]})) // fmt needed to catch result in tests.

	return nil
}

func (vm *VirtualMachine) execLoop(op *parser.Operation) error {
	operations := op.Value.([]*parser.Operation)

	for vm.memory[vm.position] != 0 {
		if err := vm.run(operations); err != nil {
			return err
		}
	}

	return nil
}
