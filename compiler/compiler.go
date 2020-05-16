package hrm

import (
	"fmt"
	"strconv"
)

/* A parser is used to generate bytecode for the interpreter. */
type Parser struct {
	backpatch map[string][]Label
	current Token
	previous Token
	hasError bool
	errorState bool
	labels map[string]byte
	chunk *Chunk
	scanner *Scanner
	size int
}

type Label struct {
	offset int
	token Token
}

/* Handler for compile-time errors. */
func (p *Parser) raiseError(token Token, err string) {
	if p.errorState {
		return
	}
	fmt.Printf("[Ln %d:%d] Parsing error: %s\n", token.line, token.column, err)
	p.hasError = true
	p.errorState = true
}

/* Ignores upcoming tokens in an effort to recover from the error state. */
func (p *Parser) synchronize() {
	p.errorState = false
	for p.previous.Type != EOF {
		if p.previous.Type == NEWLINE {
			return
		}
		switch p.current.Type {
		case INBOX, OUTBOX, JUMP, JUMPZ:
			return
		// Otherwise, skip the current token
		}
		p.advance()
	}
}

/* Retrieves and sets the next available token. */
func (p *Parser) advance() {
	p.previous = p.current
	for {
		p.current = p.scanner.ScanToken()
		if p.current.Type != ERROR {
			break
		}
		// Report errors while scanning tokens
		p.raiseError(p.current, p.current.Literal)
		p.advance()
	}
}

/* Consumes the next available token. */
func (p *Parser) consume(tokenType TokenType, err string) {
	if (p.current.Type == tokenType) {
		p.advance()
		return
	}
	p.raiseError(p.current, err)
}

/* Matches the next token to a type, advancing the parser if
the match is successful. */
func (p *Parser) match(tokenType TokenType) bool {
	if !p.check(tokenType) {
		return false
	}
	p.advance()
	return true
}

/* Checks if the current token is a certain type. */
func (p *Parser) check(tokenType TokenType) bool {
	return p.current.Type == tokenType
}

/* Checks if all labels used are declared.
This is done by ensuring no more labels need to be backpatched. */
func (p *Parser) checkLabels() bool {
	if len(p.backpatch) > 0 {
		for label, data := range p.backpatch {
			p.raiseError(data[0].token, fmt.Sprintf("Unknown label '%s'.", label))
			return false
		}
	}
	return true
}

/* Patches labels which were used before declaration by setting the jump offset. */
func (p *Parser) patchLabel(label string) {
	if labels, ok := p.backpatch[label]; ok {
		for _, data := range labels {
			p.chunk.code[data.offset] = p.labels[label]
		}
		// Label no longer needs backpatching since it is declared
		delete(p.backpatch, label)
	}
}

/* Writes a bytecode instruction to the current chunk. */
func (p *Parser) emitByte(b byte) {
	p.chunk.Write(b, p.previous.line)
}

/* Writes 2 bytecode instructions to the current chunk. */
func (p *Parser) emitBytes(b1, b2 byte) {
	p.emitByte(b1)
	p.emitByte(b2)
}

/* Emits a HALT instruction. */
func (p *Parser) emitHalt() {
	p.emitByte(OP_HALT)
}

/* Emits a CONSTANT instruction followed by its index in the constant pool. */
func (p *Parser) emitConstant(v Value) {
	index := byte(p.chunk.AddConst(v))
	p.emitBytes(OP_CONSTANT, index)
}

/* Precedence levels specify the parsing order of productions. */
type Precedence int
const (
	PREC_NONE Precedence = iota
	PREC_UNARY
	PREC_PRIMARY
)

/* Parses a statement. */
func (p *Parser) statement() {
	switch {
	case p.match(NEWLINE):
		return
	case p.match(INBOX):
		p.inbox()
		p.size += 1
	case p.match(OUTBOX):
		p.outbox()
		p.size += 1
	case p.match(JUMP):
		p.jump()
		p.size += 1
	case p.match(JUMPZ):
		p.jumpIfZero()
		p.size += 1
	case p.match(COPYFROM):
		p.copyfrom()
		p.size += 1
	case p.match(COPYTO):
		p.copyto()
		p.size += 1
	case p.match(ADD):
		p.add()
		p.size += 1
	case p.match(LABEL):
		p.labelDeclaration()
	case p.match(EOF):
		return
	default:
		p.exprStatement()
	}
}

/* Parses an expression statement, which evaluates without affecting
the VM stack. */
func (p *Parser) exprStatement() {
	p.expression()
	// Immediately ignore the last expression result
	p.emitByte(OP_POP)
}

/* Parses a label declaration. Jump offsets are handled by the compiler. */
func (p *Parser) labelDeclaration() {
	label := p.previous.Literal
	p.consume(COLON, "Expected ':' after label declaration.")
	if _, ok := p.labels[label]; ok {
		p.raiseError(p.previous, fmt.Sprintf("Label '%s' already used.", label))
	} else {
		p.labels[label] = byte(p.chunk.count)
		p.patchLabel(label)
	}
}

/* Parses an INBOX instruction. */
func (p *Parser) inbox() {
	p.emitByte(OP_INBOX)
}

/* Parses an OUTBOX instruction. */
func (p *Parser) outbox() {
	p.emitByte(OP_OUTBOX)
}

/* Parses a JUMP [label] instruction. */
func (p *Parser) jump() {
	p.emitByte(OP_JUMP)
	p.primary()
}

/* Parses a JUMPZ instruction. */
func (p *Parser) jumpIfZero() {
	p.emitByte(OP_JUMPZ)
	p.primary()
}

/* Parses a COPYFROM [addr] instruction. */
func (p *Parser) copyfrom() {
	p.primary()
	p.emitByte(OP_COPYFROM)
}

/* Parses a COPYTO [addr] instruction. */
func (p *Parser) copyto() {
	p.primary()
	p.emitByte(OP_COPYTO)
}

/* Parses an ADD [addr] instruction. */
func (p *Parser) add() {
	p.primary()
	p.emitByte(OP_ADD)
}

/* Parses an expression. */
func (p *Parser) expression() {
	p.unary()
}

/* Parses a unary expression. */
func (p *Parser) unary() {
	switch {	
	case p.match(MINUS):
		p.unary()
		p.emitByte(OP_NEGATE)
	default:
		p.primary()
	}
}

/* Parses a primary expression (literal). */
func (p *Parser) primary() {
	switch {
	case p.match(INT):
		p.number()
	case p.match(LABEL):
		p.label()
	default:
		p.raiseError(p.current, fmt.Sprintf("Unexpected token '%s'.", p.current))
		p.advance()
	}
}

/* Parses a number. */
func (p *Parser) number() {
	num, _ := strconv.Atoi(p.previous.Literal)
	p.emitConstant(IntVal(num))
}

/* Parses a label (identifier), but emits the jump offset. */
func (p *Parser) label() {
	label := p.previous.Literal
	if offset, ok := p.labels[label]; ok {
		p.emitByte(byte(offset))
	} else {
		data := Label{offset: p.chunk.count, token: p.previous}
		p.backpatch[label] = append(p.backpatch[label], data)
		p.emitByte(0)
	}
}


/* Compiles the source code into a chunk. */
func (vm *VM) Compile(source string, chunk *Chunk) (int, bool) {
	var scanner Scanner
	var parser Parser
	scanner.Init(source)
	parser.scanner = &scanner
	parser.chunk = chunk
	parser.backpatch = map[string][]Label{}
	parser.labels = map[string]byte{}
	parser.advance()
	for parser.previous.Type != EOF {
		parser.statement()
		if parser.errorState {
			parser.synchronize()
		}
	}
	parser.consume(EOF, "Expected EOF.")
	parser.emitHalt()
	parser.checkLabels()
	return parser.size, !parser.hasError
}
