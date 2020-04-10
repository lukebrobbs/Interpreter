package lexer

import (
	"github.com/lukebrobbs/interpreter/token"
)

type Lexer struct {
	input        string
	position     int  // Current position in input (points to current char)
	readPosition int  // Current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch string(l.ch) {
	case "=":
		tok = newToken(token.ASSIGN, l.ch)
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
	case string(byte(0)):
		tok.Literal = ""
		tok.Type = token.EOF
	}
	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
