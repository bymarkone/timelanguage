package language

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	DASH      = "DASH"
	STAR      = "STAR"
	COMMA     = "COMMA"
	IDENT     = "IDENT"
	STRING    = "STRING"
	LEVEL     = "LEVEL"
	SEMICOLON = "SEMICOLON"
	EOF       = "EOF"
	ILLEGAL   = "ILLEGAL"
	LP        = "LEFT_PARENTHESIS"
	RP        = "RIGHT_PARENTHESIS"
	LSB       = "LEFT_SQUARE_BRACKETS"
	RSB       = "RIGHT_SQUARE_BRACKETS"
)
