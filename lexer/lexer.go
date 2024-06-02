package lexer

import "github.com/lavish-gambhir/qu-interpreter/token"

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	currentChar := l.ch
	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, currentChar)
	case ';':
		tok = newToken(token.SEMICOLON, currentChar)
	case '(':
		tok = newToken(token.LPAREN, currentChar)
	case ')':
		tok = newToken(token.RPAREN, currentChar)
	case ',':
		tok = newToken(token.COMMA, currentChar)
	case '+':
		tok = newToken(token.PLUS, currentChar)
	case '}':
		tok = newToken(token.RBRACE, currentChar)
	case '{':
		tok = newToken(token.LBRACE, currentChar)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}
	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readChar() {
	if l.readPosition > len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}
