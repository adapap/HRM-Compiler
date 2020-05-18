package hrm
import (
	"fmt"
	"os"
)

/* Returns a slice of integers [start..stop] by step. */
func IntegerSlice(start, stop, step int) []Value {
	result := make([]Value, 0)
	for i := start; i != stop + step; i += step {
		result = append(result, IntVal(i))
	}
	return result
}

/* Returns a slice of runes [start..stop]. */
func RuneSlice(start, stop rune) []Value {
	result := make([]Value, 0)
	for i := start; i <= stop; i += 1 {
		result = append(result, CharVal(i))
	}
	return result
}

// Built-in test inputs
var POSITIVE_INTEGERS = IntegerSlice(0, 10, 1)
var ALL_INTEGERS = IntegerSlice(-10, 10, 1)
var LARGE_INTEGERS = IntegerSlice(-99, 99, 1)
var ALPHABET = RuneSlice('a', 'z')
var ALPHANUMERIC = append(ALL_INTEGERS, ALPHABET...)

/* Allocates the registers with a set amount of empty values. */
func allocateRegisters(n int, registers *[]Value) {
	for i := 0; i < n; i += 1 {
		*registers = append(*registers, EmptyVal())
	}
}

/* Generates the cartesian product of an iterable with n repeats. */
func product(n int, iterable []Value) [][]Value {
	indices := make([]int, n)
	var result [][]Value
	for indices != nil {
		p := make([]Value, n)
		for i, x := range indices {
			p[i] = iterable[x]
		}
		for i := len(indices) - 1; i >= 0; i -= 1 {
			indices[i] += 1
			if indices[i] < len(iterable) {
				break
			}
			indices[i] = 0
			if i <= 0 {
				indices = nil
				break
			}
		}
		result = append(result, p)
	}
	return result
}

/* Generates a read-only channel which has the inputs
for running the VM and testing levels. Additionally sends
values to inbox channel. */
func generateInputs(n int, data ...[]Value) []Value {
	entries := make([]Value, 0)
	for _, collection := range data {
		for _, item := range collection {
			entries = append(entries, item)
		}
	}
	inputs := make([]Value, 0)
	for _, xs := range product(n, entries) {
		for _, x := range xs {
			inputs = append(inputs, x)
		}
	}
	return inputs
}

type data struct {
	inbox *[]Value
	expected *[]Value
	registers *[]Value
	goal *INFO
}
type levelFn func(data)
var Level map[int]levelFn = map[int]levelFn{
	1: Level1,
	2: Level2,
	3: Level3,
	4: Level4,
	6: Level6,
	7: Level7,
	8: Level8,
	9: Level9,
	10: Level10,
	11: Level11,
	12: Level12,
	13: Level13,
	14: Level14,
	16: Level16,
	17: Level17,
	19: Level19,
	20: Level20,
	21: Level21,
	22: Level22,
	23: Level23,
	24: Level24,
	25: Level25,
	26: Level26,
	27: Level27,
	28: Level28,
	29: Level29,
	30: Level30,
}

func TestLevel(level int, source string, debug bool) bool {
	// Simulate each level in Go
	inbox := make([]Value, 0)
	outbox := make([]Value, 0)
	expected := make([]Value, 0)
	registers := make([]Value, 0)
	goal := INFO{}
	data := data{&inbox, &expected, &registers, &goal}
	if test, ok := Level[level]; ok {
		test(data)
	} else {
		fmt.Printf("No test written for level %d.\n", level)
		os.Exit(2)
	}
	var vm VM
	vm.Init(debug, inbox, &outbox, registers)
	state, info := vm.Interpret(source)
	if state != INTERPRET_OK {
		return false
	}
	fmt.Printf("Steps: %-4d Size: %-4d\n", info.steps, info.size)
	// Assert that all outbox values are expected
	if len(expected) < len(outbox) {
		fmt.Printf("Too many values in OUTBOX.\n")
		return false
	}
	if len(expected) > len(outbox) {
		fmt.Printf("Not enough stuff in the OUTBOX! " +
		"Management expected a total of %d items, not %d!\n",
		len(expected), len(outbox))
		return false
	}
	fmt.Printf("Expecting INBOX (%d values) -> OUTBOX (%d values)...\n", len(inbox), len(expected))
	for i, expVal := range expected {
		if outVal := outbox[i]; expVal != outVal {
			fmt.Printf("Bad outbox! Management expected %v, " +
			"but you outboxed %v.\n", expVal, outVal)
			return false
		}
	}
	fmt.Printf("Level %d test passed.", level)
	return true
}
