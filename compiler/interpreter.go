package hrm

import (
	"fmt"
)

/* A virtual machine stores a chunk of data and executes it. For HRM,
a register-based VM will be used with 12 registers. Inbox is a read-only
channel, outbox is a write-only channel. */
type VM struct {
	chunk *Chunk
	debug bool
	hand Value
	inbox []Value
	ip int
	outbox *[]Value
	registers []Value
	stack []Value
	stackTop int
	steps int
}

/* Handler for runtime errors. */
func (vm *VM) raiseError(format string, err ...interface{}) {
	if len(err) > 0 {
		fmt.Printf(format, err...)
		return
	}
	instruction := vm.ip - 1
	line := vm.chunk.lines[instruction]
	fmt.Printf("[Ln %d] Runtime Error: %s", line, format)
	vm.stackTop = 0
}

/* Initializes the virtual machine. */
func (vm *VM) Init(
		debug bool,
		inbox []Value,
		outbox *[]Value,
		registers []Value,
	) {
	vm.debug = debug
	vm.inbox = inbox
	vm.outbox = outbox
	vm.registers = registers
}

/* An enum representing the state of executing the VM's instructions. */
type INTERPRET_STATE int
const (
	INTERPRET_OK INTERPRET_STATE = iota
	INTERPRET_COMPILE_ERROR
	INTERPRET_RUNTIME_ERROR
)

const (
	EMPTY_TILE_ERROR = "Empty value! You can't %s with an empty tile on the floor! " +
		"Try writing something to that tile first."
	EMPTY_HAND_ERROR = "Empty value! You can't %s with empty hands!"
	NAN_ERROR = "Value is not a number, cannot %s!"
)

/* Reads the next byte in the chunk. */
func (vm *VM) readByte() byte {
	val := vm.chunk.code[vm.ip]
	vm.ip += 1
	return val
}

/* Retrieves the next constant value from the chunk's constant pool. */
func (vm *VM) readConstant() Value {
	index := vm.readByte()
	return vm.chunk.constants[index]
}

/* Reads a register by its index. */
func (vm *VM) readRegister() int {
	value := vm.pop()
	return value.Int
}

/* Pushes a value onto the stack. */
func (vm *VM) push(value Value) {
	vm.stack = append(vm.stack, value)
	vm.stackTop += 1
}

/* Pops a value from the stack. */
func (vm *VM) pop() Value {
	vm.stackTop -= 1
	top := vm.stack[vm.stackTop]
	vm.stack = vm.stack[:vm.stackTop]
	return top
}

/* Peeks n values from the top of the stack. */
func (vm *VM) peek(n int) Value {
	return vm.stack[vm.stackTop - 1 - n]
}

/* Picks up an item, discarding anything currently held. */
func (vm *VM) take(value Value) bool {
	if value.Type == VAL_EMPTY {
		return false
	}
	vm.hand = value
	return true
}

/* Picks up an item from the given register, if possible. */
func (vm *VM) takeRegister(register int, opcode string) bool {
	ok := vm.take(vm.registers[register])
	if !ok {
		vm.raiseError(EMPTY_TILE_ERROR, opcode)
	}
	return ok
}

/* Places an item, returning the item and status.
Status is a bool that is false if no item is held. */
func (vm *VM) drop() (Value, bool) {
	if vm.hand.Type == VAL_EMPTY {
		return Value{}, false
	}
	value := vm.hand
	vm.hand = Value{}
	return value, true
}

/* Copies an item to the given register, if possible. */
func (vm *VM) copyRegister(register int, opcode string) bool {
	value := vm.hand
	if value.Type == VAL_EMPTY {
		vm.raiseError(EMPTY_HAND_ERROR, opcode)
		return false
	} else {
		vm.registers[register] = value
	}
	return true
}

/* Checks the value at a given register and returns it if not empty. */
func (vm *VM) checkRegister(register int, opcode string) (Value, bool) {
	if register < 0 || register >= len(vm.registers) {
		vm.raiseError("There are only %d slots available on this floor!", len(vm.registers))
		return Value{}, false
	}
	value := vm.registers[register]
	if value.Type == VAL_EMPTY {
		vm.raiseError(EMPTY_TILE_ERROR, opcode)
		return Value{}, false
	}
	return value, true
}

/* Executes the VM's instructions, 1 code at a time.
This is the most performance-critical part of the machine. */
func (vm *VM) run() INTERPRET_STATE {
	for {
		instruction := vm.readByte();
		if vm.debug {
			DisassembleInstruction(vm.chunk, vm.ip - 1)
			fmt.Printf("Regs    : %v\n", vm.registers)
			fmt.Printf("Stack   : %v\n", vm.stack)
			fmt.Printf("Hand	: %v\n\n", vm.hand)
		}
		switch instruction {
		case OP_HALT:
			return INTERPRET_OK
		case OP_POP:
			vm.pop();
		case OP_CONSTANT:
			constant := vm.readConstant()
			vm.push(constant)
		case OP_INBOX:
			if len(vm.inbox) > 0 {
				vm.take(vm.inbox[0])
				vm.inbox = vm.inbox[1:]
				vm.steps += 1
			} else {
				return INTERPRET_OK
			}
		case OP_OUTBOX:
			value, ok := vm.drop()
			if !ok {
				vm.raiseError(EMPTY_HAND_ERROR, "OUTBOX")
				return INTERPRET_RUNTIME_ERROR
			}
			*vm.outbox = append(*vm.outbox, value)
			vm.steps += 1
		case OP_JUMP:
			offset := vm.readByte()
			vm.ip = int(offset)
			vm.steps += 1
		case OP_JUMPZ:
			offset := vm.readByte()
			value := vm.hand
			if value.Type == VAL_INT && value.Int == 0 {
				vm.ip = int(offset)
				vm.steps += 1
			}
		case OP_JUMPN:
			offset := vm.readByte()
			value := vm.hand
			if value.Type == VAL_INT && value.Int < 0 {
				vm.ip = int(offset)
				vm.steps += 1
			}
		case OP_COPYFROM:
			register := vm.readRegister()
			ok := vm.takeRegister(register, "COPYFROM")
			if !ok {
				return INTERPRET_RUNTIME_ERROR
			}
			vm.steps += 1
		case OP_COPYTO:
			register := vm.readRegister()
			ok := vm.copyRegister(register, "COPYTO")
			if !ok {
				return INTERPRET_RUNTIME_ERROR
			}
			vm.steps += 1
		case OP_ADD:
			register := vm.readRegister()
			value, ok := vm.checkRegister(register, "ADD")
			if !ok {
				return INTERPRET_RUNTIME_ERROR
			}
			if value.Type != VAL_INT {
				vm.raiseError(NAN_ERROR, "ADD")
				return INTERPRET_RUNTIME_ERROR
			}
			vm.hand.Int += value.Int
			vm.steps += 1
		case OP_SUB:
			register := vm.readRegister()
			value, ok := vm.checkRegister(register, "SUB")
			if !ok {
				return INTERPRET_RUNTIME_ERROR
			}
			if value.Type != VAL_INT {
				vm.raiseError(NAN_ERROR, "SUB")
				return INTERPRET_RUNTIME_ERROR
			}
			vm.hand.Int -= value.Int
			vm.steps += 1
		case OP_BUMPUP:
			register := vm.readRegister()
			value, ok := vm.checkRegister(register, "BUMP+")
			if !ok {
				return INTERPRET_RUNTIME_ERROR
			}
			if value.Type != VAL_INT {
				vm.raiseError(NAN_ERROR, "BUMP+")
				return INTERPRET_RUNTIME_ERROR
			}
			vm.hand.Int = value.Int + 1
			vm.copyRegister(register, "BUMP+")
			vm.steps += 1
		case OP_BUMPDN:
			register := vm.readRegister()
			value, ok := vm.checkRegister(register, "BUMP-")
			if !ok {
				return INTERPRET_RUNTIME_ERROR
			}
			if value.Type != VAL_INT {
				vm.raiseError(NAN_ERROR, "BUMP-")
				return INTERPRET_RUNTIME_ERROR
			}
			vm.hand.Int = value.Int - 1
			vm.copyRegister(register, "BUMP-")
			vm.steps += 1
		case OP_NEGATE:
			fmt.Printf("OP_NEGATE\n")
			if vm.peek(0).Type != VAL_INT {
				vm.raiseError(NAN_ERROR, "NEGATE")
				return INTERPRET_RUNTIME_ERROR
			}
			vm.push(IntVal(-vm.pop().Int))
		default:
			fmt.Printf("Unknown opcode %d.", instruction)
			return INTERPRET_RUNTIME_ERROR
		}
	}
}

type INFO struct {
	steps int
	size int
}
/* Interprets the source string of code. */
func (vm *VM) Interpret(source string) (INTERPRET_STATE, INFO) {
	var chunk Chunk
	chunk.Init()
	size, ok := vm.Compile(source, &chunk)
	if !ok {
		return INTERPRET_COMPILE_ERROR, INFO{}
	}
	vm.chunk = &chunk
	vm.ip = 0
	vm.steps = 0
	result := vm.run()
	info := INFO{vm.steps, size}
	return result, info
}
