package hrm
import (
	// "fmt"
	"math"
)

/* By design, test cases are generated automatically to consider individual edge cases
for each level. A better approach would be to determine the exact test case input in
the game; this way, we get the same number of steps during execution time rather than
testing all possible combinations of expected inputs. */

/* 
Level 1: Mail Room
Drag commands into this area to build a program.

Your program should tell your worker to grab each thing from the INBOX, and drop
it into the OUTBOX.
*/

func Level1(d data) {
	*d.inbox = generateInputs(1, IntegerSlice(1, 3, 1))
	for _, x := range *d.inbox {
		*d.expected = append(*d.expected, x)
	}
}

/*
Level 2: Busy Mail Room
Grab each thing from the inbox, and drop each one into the OUTBOX.

You got a new command! You can drag JUMP's arrow to different lines within your
program.
*/
func Level2(d data) {
	for _, c := range "INITIALIZE" {
		*d.inbox = append(*d.inbox, CharVal(c))
	}
	for _, c := range "BOOTSEQUENCE" {
		*d.inbox = append(*d.inbox, CharVal(c))
	}
	for _, c := range "AUTOEXEC" {
		*d.inbox = append(*d.inbox, CharVal(c))
	}
	for _, x := range *d.inbox {
		*d.expected = append(*d.expected, x)
	}
}

/*
Level 3: Copy Floor
Ignore the INBOX for now, and just send the following 3 letters to the outbox:
B U G

The Facilities Management staff has placed some items over there on the carpet
for you. If only there were a way to pick them up...
*/
func Level3(d data) {
	for i := 0; i < 4; i += 1 {
		*d.inbox = append(*d.inbox, IntVal(-99))
	}
	for _, c := range "BUG" {
		*d.expected = append(*d.expected, CharVal(c))
	}
	*d.registers = []Value{
		CharVal('U'),
		CharVal('J'),
		CharVal('X'),
		CharVal('G'),
		CharVal('B'),
		CharVal('E'),
	}
}

/*
Level 4: Scrambler Handler
Grab the first TWO things from the INBOX and drop them into the OUTBOX in the
reverse order. Repeat until the INBOX is empty.

You got a new command! Feel free to COPYTO wherever you like on the carpet. It
will be cleaned later.
*/
func Level4(d data) {
	inbox := generateInputs(2, ALPHANUMERIC)
	for i := 0; i + 1 < len(inbox); i += 2 {
		*d.expected = append(*d.expected, inbox[i + 1], inbox[i])
	}
	*d.inbox = inbox
	allocateRegisters(3, d.registers)
}

/* Level 5: Coffee Time (Cutscene) */

/* 
Level 6: Rainy Summer
For each two things in the INBOX, add them together, and put the result
in the OUTBOX.

You got a new command! It ADDs the contents of a tile on the floor to whatever
value you're currently holding.
*/
func Level6(d data) {
	inbox := generateInputs(2, ALL_INTEGERS)
	for i := 0; i + 1 < len(inbox); i += 2 {
		sum := inbox[i].Int + inbox[i + 1].Int
		*d.expected = append(*d.expected, IntVal(sum))
	}
	*d.inbox = inbox
	allocateRegisters(3, d.registers)
}

/* 
Level 7: Zero Exterminator
Send all things that ARE NOT ZERO to the OUTBOX.

You got a new command! It jumps ONLY if the value you are holding is ZERO. Otherwise
it continues to the next line. */
func Level7(d data) {
	inbox := generateInputs(1, ALPHANUMERIC)
	for i := 0; i < len(inbox); i += 1 {
		if inbox[i].Type != VAL_INT || inbox[i].Int != 0 {
			*d.expected = append(*d.expected, inbox[i])
		}
	}
	*d.inbox = inbox
	allocateRegisters(9, d.registers)
}

/* 
Level 8: Tripler Room
For each thing in the INBOX, TRIPLE it. And OUTBOX the result.

Self-improvement tip: Where are we going with this? Please leave the high level
decisions to management.
*/
func Level8(d data) {
	inbox := generateInputs(1, ALL_INTEGERS)
	for i := 0; i < len(inbox); i += 1 {
		num := inbox[i].Int * 3
		*d.expected = append(*d.expected, IntVal(num))
	}
	*d.inbox = inbox
	allocateRegisters(3, d.registers)
}

/*
Level 9: Zero Preservation Initiative
Send only ZEROs to the OUTBOX. */
func Level9(d data) {
	inbox := generateInputs(1, ALPHANUMERIC)
	for i := 0; i < len(inbox); i += 1 {
		if inbox[i].Type == VAL_INT && inbox[i].Int == 0 {
			*d.expected = append(*d.expected, inbox[i])
		}
	}
	*d.inbox = inbox
	allocateRegisters(9, d.registers)
}

/*
Level 10: Octoplier Suite
For each thing in the INBOX, multiply it by 8, and put the result in the OUTBOX.

Using a bunch of ADD commands is easy, but WASTEFUL! Can you do it using only
3 ADD commands? Management is watching. */
func Level10(d data) {
	inbox := generateInputs(1, ALL_INTEGERS)
	for i := 0; i < len(inbox); i += 1 {
		num := inbox[i].Int * 8
		*d.expected = append(*d.expected, IntVal(num))
	}
	*d.inbox = inbox
	allocateRegisters(5, d.registers)
}

/*
Level 11: Sub Hallway
For each two things in the INBOX, first subtract the 1st from the 2nd
and put the result in the OUTBOX. AND THEN, subtract the 2nd from the
1st and put the result in the OUTBOX. Repeat.

You got a new command! SUBtracts the contents of a tile on the floor
FROM whatever value you're currently holding. */
func Level11(d data) {
	/* todo not working, works in game */
	inbox := generateInputs(2, ALL_INTEGERS)
	for i := 0; i + 1 < len(inbox); i += 2 {
		diff := inbox[i].Int - inbox[i + 1].Int
		rdiff := inbox[i + 1].Int - inbox[i].Int
		*d.expected = append(*d.expected, IntVal(rdiff), IntVal(diff))
	}
	*d.inbox = inbox
	allocateRegisters(3, d.registers)
}

/*
Level 12: Tetracontiplier
For each thing in the INBOX, multiply it by 40,
and put the result in the OUTBOX. */
func Level12(d data) {
	inbox := generateInputs(1, ALL_INTEGERS)
	for i := 0; i < len(inbox); i += 1 {
		num := inbox[i].Int * 40
		*d.expected = append(*d.expected, IntVal(num))
	}
	*d.inbox = inbox
	allocateRegisters(5, d.registers)
}

/*
Level 13: Equalization Room
Get two things from the INBOX. If they are EQUAL, put ONE of them in the OUTBOX.
Discard non-equal pairs. Repeat!

You got... COMMENTS! You can use them, if you like, to mark sections of your program. */
func Level13(d data) {
	inbox := generateInputs(2, ALL_INTEGERS)
	for i := 0; i + 1 < len(inbox); i += 2 {
		a := inbox[i]
		b := inbox[i + 1]
		if a.Int == b.Int {
			*d.expected = append(*d.expected, IntVal(a.Int))
		}
	}
	*d.inbox = inbox
	allocateRegisters(3, d.registers)
}

/*
Level 14: Maximization Room
Grab TWO things from the INBOX, and put only the BIGGER of the two in the OUTBOX.
If they are equal, just pick either one. Repeat!

You got a new command! Jumps only if the thing you're holding is negative.
(Less than zero). Otherwise continues to the next line. */
func Level14(d data) {
	inbox := generateInputs(2, ALL_INTEGERS)
	for i := 0; i + 1 < len(inbox); i += 2 {
		num := int(math.Max(float64(inbox[i].Int), float64(inbox[i + 1].Int)))
		*d.expected = append(*d.expected, IntVal(num))
	}
	*d.inbox = inbox
	allocateRegisters(3, d.registers)
}

/* Level 15: Employee Morale Insertion (Cutscene) */

/*
Level 16: Absolute Positivity
Send each thing from the INBOX to the OUTBOX. BUT, if a number is negative,
first remove its negative sign. */
func Level16(d data) {
	inbox := generateInputs(1, ALL_INTEGERS)
	for i := 0; i < len(inbox); i += 1 {
		num := int(math.Abs(float64(inbox[i].Int)))
		*d.expected = append(*d.expected, IntVal(num))
	}
	*d.inbox = inbox
	allocateRegisters(3, d.registers)
}

/*
Level 17: Exclusive Lounge
For each TWO things in the INBOX:

Send a 0 to the OUTBOX if they have the same sign. (Both positive or both negative.) 

Send a 1 to the OUTBOX if their signs are different. Repeat until the INBOX is empty. */
func Level17(d data) {
	inbox := generateInputs(2, ALL_INTEGERS)
	for i := 0; i + 1 < len(inbox); i += 2 {
		a := inbox[i].Int
		b := inbox[i + 1].Int
		var num int
		if math.Signbit(float64(a)) == math.Signbit(float64(b)) {
			num = 0
		} else {
			num = 1
		}
		*d.expected = append(*d.expected, IntVal(num))
	}
	*d.inbox = inbox
	*d.registers = []Value{
		EmptyVal(),
		EmptyVal(),
		EmptyVal(),
		EmptyVal(),
		IntVal(0),
		IntVal(1),
	}
}

/* Level 18: Sabbatical Beach Paradise (Cutscene) */

/*
Level 19: Countdown
For each number in the INBOX, send that number to the OUTBOX,
followed by all numbers down to (or up to) zero. It's a countdown!

You got new commands! They add ONE or subtract ONE from an item on the floor.
The result is given back to you, and for your convenience, also written right
back on the floor. BUMP! */
func Level19(d data) {
	inbox := generateInputs(1, ALL_INTEGERS)
	for i := 0; i < len(inbox); i += 1 {
		num := inbox[i].Int
		for num != 0 {
			*d.expected = append(*d.expected, IntVal(num))
			if num < 0 {
				num += 1
			} else {
				num -= 1
			}
		}
		*d.expected = append(*d.expected, IntVal(num))
	}
	*d.inbox = inbox
	allocateRegisters(10, d.registers)
}

/*
Level 20: Multiplication Workshop
For each two things in the INBOX, multiply them, and OUTBOX the result.
Don't worry about negative numbers for now.

You got... LABELS! They can help you remember the purpose of each tile on the
floor. Just tap any tile on the floor to edit. */
func Level20(d data) {
	inbox := generateInputs(2, POSITIVE_INTEGERS)
	for i := 0; i + 1 < len(inbox); i += 2 {
		a := inbox[i].Int
		b := inbox[i + 1].Int
		*d.expected = append(*d.expected, IntVal(a * b))
	}
	*d.inbox = inbox
	*d.registers = []Value{
		EmptyVal(),
		EmptyVal(),
		EmptyVal(),
		EmptyVal(),
		EmptyVal(),
		EmptyVal(),
		EmptyVal(),
		EmptyVal(),
		EmptyVal(),
		IntVal(0),
	}
}

/*
Level 21: Zero Terminated Sum
The INBOX is filled with ZERO terminated strings! What's that? Ask me. Your Boss.

Add together all the numbers in each string. When you reach the end of a string
(marked by a ZERO), put your sum in the OUTBOX. Reset and repeat for each string. */
func Level21(d data) {
	inbox := generateInputs(1, ALL_INTEGERS, []Value{IntVal(0)}, POSITIVE_INTEGERS, []Value{IntVal(0)})
	sum := 0
	for i := 0; i < len(inbox); i += 1 {
		if inbox[i].Int != 0 {
			sum += inbox[i].Int
		} else {
			*d.expected = append(*d.expected, IntVal(sum))
			sum = 0
		}
	}
	*d.inbox = inbox
	*d.registers = []Value{
		EmptyVal(),
		EmptyVal(),
		EmptyVal(),
		EmptyVal(),
		EmptyVal(),
		IntVal(0),
	}
}

/*
Level 22: Fibonacci Visitor
For each thing in the INBOX, send to the OUTBOX the full Fibonacci Sequence
up to, but not exceeding that value. For example, if INBOX is 10, OUTBOX should
be 1 1 2 3 5 8. What's a Fibonacci Sequence? Ask your boss, or a friendly
search box.

1 1 2 3 5 8 13 21 34 55 89... */
func Level22(d data) {
	inbox := generateInputs(1, POSITIVE_INTEGERS)
	for i := 0; i < len(inbox); i += 1 {
		n := inbox[i].Int
		for a, b := 0, 1; b <= n; {
			*d.expected = append(*d.expected, IntVal(b))
			sum := a + b
			a = b
			b = sum
		}
	}
	*d.inbox = inbox
	*d.registers = []Value{
		EmptyVal(),
		EmptyVal(),
		EmptyVal(),
		EmptyVal(),
		EmptyVal(),
		EmptyVal(),
		EmptyVal(),
		EmptyVal(),
		EmptyVal(),
		IntVal(0),
	}
}

/*
Level 23: The Littlest Number
For each two things in the INBOX, how many times does the second fully fit into
the first? Don't worry about negative numbers, divide by zero, or remainders.

Self improvement tip: This might be a good time to practice copying and pasting
from a previous assignment! */
func Level23(d data) {
	
}

/*
Level 24: Mod Module
For each two things in the INBOX, OUTBOX the remainder that would result if you had
divided the first by the second. Don't worry, you don't actually have to divide.
And don't worry about negative numbers for now. */
func Level24(d data) {
	inbox := generateInputs(1, POSITIVE_INTEGERS)
	for i := 0; i + 1 < len(inbox); i += 2 {
		num := inbox[i].Int % inbox[i + 1].Int
		*d.expected = append(*d.expected, IntVal(num))
	}
	*d.inbox = inbox
	allocateRegisters(10, d.registers)
}

/*
Level 25: Cumulative Countdown
... */
func Level25(d data) {
	
}

/*
Level 26: Small Divide
... */
func Level26(d data) {
	
}

/* Level 27: Midnight Petroleum (Cutscene) */

/*
Level 28: Three Sort
... */
func Level28(d data) {
	
}

/*
Level 29: Storage Floor
... */
func Level29(d data) {
	
}

/*
Level 30: String Storage Floor
... */
func Level30(d data) {
	
}

/*
Level 31: String Reverse
... */
func Level31(d data) {
	
}

/*
Level 32: Inventory Report
... */
func Level32(d data) {
	
}

/* Level 33: Where's Carol? (Cutscene) */

/*
Level 34: Vowel Incinerator
... */
func Level34(d data) {
	
}

/*
Level 35: Duplicate Removal
... */
func Level35(d data) {
	
}

/*
Level 36: Alphabetizer
... */
func Level36(d data) {
	
}

/*
Level 37: Scavenger Chain
... */
func Level37(d data) {
	
}

/*
Level 38: Digit Exploder
... */
func Level38(d data) {
	
}

/*
Level 39: Re-Coordinator
... */
func Level39(d data) {
	
}

/*
Level 40: Prime Factory
... */
func Level40(d data) {
	
}

/*
Level 41: Sorting Floor
... */
func Level41(d data) {
	
}

/* Level 42: End Program. Congratulations. */
