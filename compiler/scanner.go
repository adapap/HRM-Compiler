package hrm

import (
	"fmt"
)
/* A token represents a single lexeme. */
type TokenType int
type Token struct {
	Type TokenType
	Literal string
	column int
	line int
}
func (t Token) String() string {
	return fmt.Sprintf("%s", t.Literal)
}

/* Enum of token types for reference. */
const (
	// Literals
	INT TokenType = iota
	LABEL
	
	// Operators + Punctuation
	MINUS
	COLON
	
	// Keywords
	INBOX
	OUTBOX
	JUMP
	JUMPZ
	JUMPN
	COPYFROM
	COPYTO
	ADD
	SUB
	
	// Other
	NEWLINE
	EOF
	ERROR
)

/* Scanner produces tokens from a string of source code. */
type Scanner struct {
	column int
	char byte
	current int
	line int
	source string
	start int
}

/* Initialize the scanner to start at the beginning of the source. */
func (s *Scanner) Init(source string) {
	s.line = 1
	s.column = 1
	s.source = source
	s.char = source[s.current]
}

/* Scan for the next token in the stream.
Tokens are produced as needed, and are not stored. */
func (s *Scanner) ScanToken() Token {
	s.skipWhitespace()
	s.start = s.current
	token := Token{
		line: s.line,
		column: s.column,
	}
	switch s.char {
	case '-':
		if s.peek(1) == '-' {
			for !(s.char == '\n' || s.char == 0) {
				s.advance()
			}
			s.advance()
			s.line += 1
			s.column = 1
			return s.ScanToken()
		} else {
			s.column += 1
			token.Type = MINUS
			token.Literal = "-"
		}
	case ':':
		s.column += 1
		token.Type = COLON
		token.Literal = ":"
	case '\n':
		s.line += 1
		s.column = 1
		token.Type = NEWLINE
		token.Literal = "NEWLINE"
	case 0:
		token.Type = EOF
		token.Literal = "EOF"
	default:
		switch {
		case isDigit(s.char):
			return s.scanInteger(token)
		case isAlpha(s.char):
			return s.scanIdentifier(token)
		default:
			s.column += 1
			token.Literal = fmt.Sprintf("Unexpected character '%s'.", string(s.char))
			token.Type = ERROR
		}
	}
	s.advance()
	return token
}

/* Advances the scanner to the next character. */
func (s *Scanner) advance() {
	s.current += 1
	s.column += 1
	if s.current >= len(s.source) {
		s.char = 0
		return
	}
	s.char = s.source[s.current]
}

/* Peeks n characters ahead. */
func (s *Scanner) peek(n int) byte {
	return s.source[s.current + n]
}

/* Advances the scanner, skipping all whitespace encountered. */
func (s *Scanner) skipWhitespace() {
	for isWhitespace(s.char) {
		s.column += 1
		s.advance()
	}
}

/* Scans for the next integer, producing a token. */
func (s *Scanner) scanInteger(t Token) Token {
	literal := ""
	for isDigit(s.char) {
		s.column += 1
		literal += string(s.char)
		s.advance()
	}
	t.Type = INT
	t.Literal = literal
	return t
}

/* Scans for the next identifier or keyword, producing a token. */
func (s *Scanner) scanIdentifier(t Token) Token {
	literal := ""
	for isIdentifier(s.char) {
		s.column += 1
		literal += string(s.char)
		s.advance()
	}
	t.Literal = literal
	t.Type = s.matchKeyword()
	return t
}

/* Checks if a lexeme matches the characters for a keyword.
Returns the keyword type if match is successful, otherwise
it is an identifier. */
func (s *Scanner) matchKeyword() TokenType {
	keyword := s.source[s.start:s.current]
	switch keyword {
	case "ADD":
		return ADD
	case "COPYFROM":
		return COPYFROM
	case "COPYTO":
		return COPYTO
	case "INBOX":
		return INBOX
	case "JUMP":
		return JUMP
	case "JUMPN":
		return JUMPN
	case "JUMPZ":
		return JUMPZ
	case "OUTBOX":
		return OUTBOX
	case "SUB":
		return SUB
	}
	return LABEL
}

/* Checks if a character is considered whitespace. */
func isWhitespace(char byte) bool {
	switch char {
	case '\r', ' ', '\t':
		return true
	}
	return false
}

/* Checks if a character is a digit. */
func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}

/* Checks if a character is alphabetic. */
func isAlpha(char byte) bool {
	return ('a' <= char && char <= 'z') || ('A' <= char && char <= 'Z')
}

/* Checks if a character is a valid identifier. */
func isIdentifier(char byte) bool {
	return isAlpha(char) || isDigit(char)
}
