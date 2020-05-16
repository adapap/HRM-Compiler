package hrm
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
	for i := 0; i < len(inbox); i += 2 {
		*d.expected = append(*d.expected, inbox[i + 1], inbox[i])
	}
	*d.inbox = inbox
	allocateRegisters(3, d.registers)
}

/* Level 5: Coffee Time... */

/* 
Level 6: Rainy Summer
For each two things in the INBOX, add them together, and put the result
in the OUTBOX.

You got a new command! It ADDs the contents of a tile on the floor to whatever
value you're currently holding.
*/
func Level6(d data) {
	inbox := generateInputs(2, ALL_INTEGERS)
	for i := 0; i < len(inbox); i += 2 {
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

/* Level 9: Zero Preservation Initiative
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

/* Level 10: Octoplier Suite
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

/* Level 11: ... */
func Level11(d data) {
	
}

/* Level 12: Tetracontiplier */
func Level12(d data) {
	inbox := generateInputs(1, ALL_INTEGERS)
	for i := 0; i < len(inbox); i += 1 {
		num := inbox[i].Int * 40
		*d.expected = append(*d.expected, IntVal(num))
	}
	*d.inbox = inbox
	allocateRegisters(5, d.registers)
}

/* Level 13: ... */
func Level13(d data) {
	
}

/* Level 14: ... */
func Level14(d data) {
	
}

/* Level 15: ... */
func Level15(d data) {
	
}

/* Level 16: ... */
func Level16(d data) {
	
}

/* Level 17: ... */
func Level17(d data) {
	
}

/* Level 18: ... */
func Level18(d data) {
	
}

/* Level 19: ... */
func Level19(d data) {
	
}

/* Level 20: ... */
func Level20(d data) {
	
}
