package lexer

import (
	"github.com/lukebrobbs/Interpreter/token"
)

// Lexer reads a string
type Lexer struct {
	input        string
	position     int  // Current position in input (points to current char)
	readPosition int  // Current reading position in input (after current char)
	ch           byte // current char under examination
}

// New creates a new Lexer
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	l.ch = l.peekChar()
	l.position = l.readPosition
	l.readPosition++
}

// NextToken returns a token representing the current character
// It also calls readChar() to move Lexer onto the next char
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.eatWhitespace()

	switch string(l.ch) {
	case "=":
		if string(l.peekChar()) == "=" {
			tok = l.makeTwoCharToken(token.EQ)
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case ";":
		tok = newToken(token.SEMICOLON, l.ch)
	case "(":
		tok = newToken(token.LPAREN, l.ch)
	case ")":
		tok = newToken(token.RPAREN, l.ch)
	case "{":
		tok = newToken(token.LBRACE, l.ch)
	case "}":
		tok = newToken(token.RBRACE, l.ch)
	case ",":
		tok = newToken(token.COMMA, l.ch)
	case "+":
		tok = newToken(token.PLUS, l.ch)
	case "-":
		tok = newToken(token.MINUS, l.ch)
	case "!":
		if string(l.peekChar()) == "=" {
			tok = l.makeTwoCharToken(token.NOT_EQ)
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case "/":
		tok = newToken(token.SLASH, l.ch)
	case "*":
		tok = newToken(token.ASTERISX, l.ch)
	case "<":
		tok = newToken(token.LT, l.ch)
	case ">":
		tok = newToken(token.GT, l.ch)
	case string(byte(0)):
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		}
		tok = newToken(token.ILLEGAL, l.ch)
	}
	l.readChar()
	return tok
}

func (l *Lexer) eatWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) makeTwoCharToken(tt token.TokenType) token.Token {
	ch := l.ch
	l.readChar()
	literal := string(ch) + string(l.ch)
	return token.Token{Type: tt, Literal: literal}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
