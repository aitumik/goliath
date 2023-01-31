package lexer

import (
	"github.com/aitumik/interpreter/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	char         byte
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhiteSpace()

	switch l.char {
	case '+':
		tok = newToken(token.PLUS, l.char)
	case '=':
		ch := l.peekChar()
		if ch == '=' {
			tok = token.Token{Type: token.EQ, Literal: string(l.char) + string(l.char)}
			l.readChar()
		} else {
			tok = newToken(token.ASSIGN, l.char)
		}
	case '-':
		tok = newToken(token.MINUS, l.char)
	case '(':
		tok = newToken(token.LPAREN, l.char)
	case ')':
		tok = newToken(token.RPAREN, l.char)
	case '{':
		tok = newToken(token.LBRACE, l.char)
	case '}':
		tok = newToken(token.RBRACE, l.char)
	case ';':
		tok = newToken(token.SEMICOLON, l.char)
	case ',':
		tok = newToken(token.COMMA, l.char)
	case '*':
		tok = newToken(token.ASTERISK, l.char)
	case '/':
		tok = newToken(token.SLASH, l.char)
	case '!':
		ch := l.peekChar()
		if ch == '=' {
			a := l.char
			l.readChar()
			tok = token.Token{Type: token.NEQ, Literal: string(a) + string(l.char)}
		} else {
			tok = newToken(token.BANG, l.char)
		}
	case '<':
		tok = newToken(token.LT, l.char)
	case '>':
		tok = newToken(token.GT, l.char)
	case 0:
		tok = newToken(token.EOF, l.char)
	default:
		if isLetter(l.char) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isNumber(l.char) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.char)
		}
	}
	l.readChar()

	return tok
}

func (l *Lexer) readIdentifier() string {
	pos := l.position

	for isLetter(l.char) {
		l.readChar()
	}

	return l.input[pos:l.position]
}

func New(str string) *Lexer {
	l := &Lexer{input: str}
	l.readChar()
	return l
}

func newToken(tokenType token.TokenType, char byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(char)}
}
