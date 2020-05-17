package hrm

import (
	"fmt"
)

/* 
Human Resource Machine (HRM) Instruction Set
============================================
INBOX: Pick up the next thing from the INBOX.

OUTBOX: Put whatever you are holding into the outbox.

COPYFROM: Walk to a specific tile on the floor and pick up a copy of whatever
is there.

COPYTO: Copy whatever you are currently holding to a specific tile on the floor.

JUMP: Jump to a new location within your program. You can jump backward to
create loops, or jump forward to skip entire sections. The possibilities
are endless!

JUMP IF ZERO: Jump only if you are currently holding a ZERO. Otherwise continue
to the next line in your program.

JUMP IF NEGATIVE: Jump only if you are currently holding a negative number. Otherwise
continue to the next line in your program.

ADD: Add the contents of a specific tile to whatever you are currently holding.
The result goes back into your hands.

SUB: Subtract the contents of a specific tile on the floor FROM whatever you
are currently holding. The result goes back into your hands.

COMMENT: Use comments to leave helpful notes for yourself within your program.
Does not affect your program in any way, other than making it easier for you
to read!
*/

const (
	OP_HALT byte = iota
	OP_POP
	OP_CONSTANT
	OP_INBOX
	OP_OUTBOX
	OP_JUMP
	OP_JUMPZ
	OP_JUMPN
	OP_COPYFROM
	OP_COPYTO
	OP_ADD
	OP_SUB
	OP_COMMENT
	OP_NEGATE
)

type ValueType int
const (
	VAL_EMPTY ValueType = iota
	VAL_INT
	VAL_CHAR
	VAL_LABEL
)

type Value struct {
	Type ValueType
	Char rune
	Int int
	Label string
}
func (v Value) String() string {
	switch v.Type {
	case VAL_EMPTY:
		return "."
	case VAL_INT:
		return fmt.Sprintf("<Int %d>", v.Int)
	case VAL_CHAR:
		return fmt.Sprintf("<Char %s>", string(v.Char))
	case VAL_LABEL:
		return fmt.Sprintf("<Lbl %s>", v.Label)
	default:
		return fmt.Sprintf("<%v?>", v.Type)
	}
}

/* Create values using specified types. */
func EmptyVal() Value {
	return Value{Type: VAL_EMPTY}
}

func CharVal(c rune) Value {
	return Value{Type: VAL_CHAR, Char: c}
}

func IntVal(i int) Value {
	return Value{Type: VAL_INT, Int: i}
}

func LabelVal(l string) Value {
	return Value{Type: VAL_LABEL, Label: l}
}
