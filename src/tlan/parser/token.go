package parser

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ITEM      = "ITEM"
	IDENT     = "IDENT"
	STRING    = "STRING"
	LEVEL     = "LEVEL"
	SEMICOLON = "SEMICOLON"
	EOF       = "EOF"
	ILLEGAL   = "ILLEGAL"
)
