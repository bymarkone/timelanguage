package language

import (
	"strings"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() Token {
	var tok Token

	l.skipNewLine()

	switch l.ch {
	case '"':
		tok.Type = STRING
		tok.Literal = l.readString()
	case '-':
		tok = newToken(DASH, l.ch)
	case '*':
		tok = newToken(STAR, l.ch)
	case '(':
		tok = newToken(LP, l.ch)
	case ')':
		tok = newToken(RP, l.ch)
	case '[':
		tok = newToken(LSB, l.ch)
	case ',':
		tok = newToken(COMMA, l.ch)
	case ']':
		tok = newToken(RSB, l.ch)
	case '+':
		tok = newToken(PLUS, l.ch)
	case '>':
		l.readChar()
		if l.ch == '>' {
			tok = Token{Type: DUALARROW, Literal: ">>"}
			l.readChar()
			return tok
		} else {
			tok = newToken(ARROW, l.ch)
			return tok
		}
	case ' ':
		if isSpace(l.peekChar()) {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = Token{Type: LEVEL, Literal: literal}
		} else {
			l.readChar()
			return l.NextToken()
		}
	case 0:
		tok.Literal = ""
		tok.Type = EOF
	default:
		if isLetter(l.ch) || isNumber(l.ch) || isSlash(l.ch) || isColon(l.ch) ||
			isComma(l.ch) || isDot(l.ch) || isBang(l.ch) || isHash(l.ch) {
			tok.Type = IDENT
			tok.Literal = l.readIdentifier()
			return tok
		} else {
			tok = newToken(ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

//private
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return l.input[position:l.position]
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for {
		l.readChar()
		if !(isLetter(l.ch) || isNumber(l.ch) || isSpace(l.ch) || isSlash(l.ch) || isColon(l.ch) ||
			isComma(l.ch) || isDot(l.ch) || isBang(l.ch) || isHash(l.ch)) {
			break
		}
	}
	return strings.TrimSpace(l.input[position:l.position])
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) skipNewLine() {
	for l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

//local
func newToken(tokenType TokenType, ch byte) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isNumber(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isSpace(ch byte) bool {
	return ch == ' '
}

func isSlash(ch byte) bool {
	return ch == '\\' || ch == '/'
}

func isColon(ch byte) bool {
	return ch == ':'
}

func isComma(ch byte) bool {
	return ch == ','
}

func isDot(ch byte) bool {
	return ch == '.'
}

func isBang(ch byte) bool {
	return ch == '!'
}

func isHash(ch byte) bool {
	return ch == '#'
}
