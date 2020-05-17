package hrm

import (
	"fmt"
)

/* Slices in Go are dynamically resized. Internally,
they have a length and capacity. When the new length
would exceed the cap, memory is allocated to double
the capacity. */
type Chunk struct {
	count int
	code []byte
	constants []Value
	lines []int
}

/* Initialize chunks with a capacity of 8 (OpCodes). */
func (chunk *Chunk) Init() {
	chunk.code = make([]byte, 0, 8)
	chunk.lines = make([]int, 0, 8)
}

/* Writes an OpCode to the chunk */
func (chunk *Chunk) Write(data byte, line int) {
	chunk.code = append(chunk.code, data)
	chunk.lines = append(chunk.lines, line)
	chunk.count += 1
}

/* Adds a constant to the chunk constant pool.
Returns its index in the constant pool. */
func (chunk *Chunk) AddConst(data Value) int {
	chunk.constants = append(chunk.constants, data)
	return len(chunk.constants) - 1
}

/* Inspect the chunk and its contents for debugging. */
func (chunk *Chunk) Disassemble(name string) {
	fmt.Printf("[%s]\n", name)
	for offset := 0; offset < chunk.count; {
		offset = DisassembleInstruction(chunk, offset)
	}
}

/* Disassembles an instruction into a human readable format. */
func DisassembleInstruction(chunk *Chunk, offset int) int {
	fmt.Printf("%04d ", offset)
	if offset > 0 && chunk.lines[offset] == chunk.lines[offset - 1] {
		fmt.Printf("   | ")
	} else {
		fmt.Printf("%4d ", chunk.lines[offset])
	}
	instruction := chunk.code[offset]
	switch instruction {
	case OP_HALT:
		return simpleInstruction("HALT", offset)
	case OP_POP:
		return simpleInstruction("POP", offset)
	case OP_CONSTANT:
		return constantInstruction("CONSTANT", chunk, offset)
	case OP_INBOX:
		return simpleInstruction("INBOX", offset)
	case OP_OUTBOX:
		return simpleInstruction("OUTBOX", offset)
	case OP_JUMP:
		return byteInstruction("JUMP", chunk, offset)
	case OP_JUMPZ:
		return byteInstruction("JUMP IF ZERO", chunk, offset)
	case OP_JUMPN:
		return byteInstruction("JUMP IF NEGATIVE", chunk, offset)
	case OP_COPYFROM:
		return byteInstruction("COPYFROM", chunk, offset)
	case OP_COPYTO:
		return byteInstruction("COPYTO", chunk, offset)
	case OP_ADD:
		return byteInstruction("ADD", chunk, offset)
	case OP_NEGATE:
		return simpleInstruction("NEGATE", offset)
	default:
		fmt.Printf("Unknown opcode %d.\n", instruction)
		return offset + 1
	}
}

/* Simple instructions do not take any operands. */
func simpleInstruction(name string, offset int) int {
	fmt.Printf("%s\n", name)
	return offset + 1
}

/* Constant instructions are loaded from the chunk constant pool. */
func constantInstruction(name string, chunk *Chunk, offset int) int {
	constant := chunk.code[offset + 1]
	fmt.Printf("%-16s %4d '", name, constant)
	fmt.Printf("%v'\n", chunk.constants[constant])
	return offset + 2
}

/* Byte instructions take one byte argument. */
func byteInstruction(name string, chunk *Chunk, offset int) int {
	value := chunk.code[offset + 1]
	fmt.Printf("%-16s %4d\n", name, value)
	return offset + 2
}
