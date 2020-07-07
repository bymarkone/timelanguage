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
  FIRST     = "FIRST"
  SEMICOLON = "SEMICOLON"
  EOL       = "EOL"
  EOF       = "EOF"
  ILLEGAL   = "ILLEGAL"
)
