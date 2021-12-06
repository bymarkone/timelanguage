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
	DOT       = "DOT"
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
	ARROW     = "ARROW"
	DUALARROW = "DUAL_ARROW"
	PLUS      = "PLUS"
)
