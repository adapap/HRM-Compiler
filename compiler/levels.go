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
	*d.inbox = generateInputs(1, IntegerSlice(1, 4, 1))
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
... */
func Level19(d data) {
	
}

/*
Level 20: Multiplication Workshop
... */
func Level20(d data) {
	
}

/*
Level 20: Multiplication Workshop
... */
func Level21(d data) {
	
}

/*
Level 20: Multiplication Workshop
... */
func Level22(d data) {
	
}

/*
Level 20: Multiplication Workshop
... */
func Level23(d data) {
	
}

/*
Level 20: Multiplication Workshop
... */
func Level24(d data) {
	
}

/*
Level 20: Multiplication Workshop
... */
func Level25(d data) {
	
}

/*
Level 20: Multiplication Workshop
... */
func Level26(d data) {
	
}

/*
Level 20: Multiplication Workshop
... */
func Level27(d data) {
	
}

/*
Level 20: Multiplication Workshop
... */
func Level28(d data) {
	
}

/*
Level 20: Multiplication Workshop
... */
func Level29(d data) {
	
}

/*
Level 20: Multiplication Workshop
... */
func Level30(d data) {
	
}
