package parser

import "testing"

func TestAnything(t *testing.T) {
  input := `
- Change the world: "Because the world needs to be changed"
  - What I build
  - How I live
`
  cases := []struct {
    expectedType    TokenType
    expectedLiteral string
  }{
    {ITEM, "-"},
    {IDENT, "Change the world"},
    {SEMICOLON, ":"},
    {STRING, "Because the world needs to be changed"},
    {LEVEL, "  "},
    {ITEM, "-"},
    {IDENT, "What I build"},
    {LEVEL, "  "},
    {ITEM, "-"},
    {IDENT, "How I live"},
  }

  lexer := NewLexer(input)

  for i, tt := range cases {
    tok := lexer.NextToken()
    if tok.Type != tt.expectedType {
      t.Fatalf("tests[%d] - token type is wrong, expected %q, got %q (literal %q)", i, tt.expectedType, tok.Type, tok.Literal)
    }
    if tok.Literal != tt.expectedLiteral {
      t.Fatalf("tests[%d] - literal is wrong, expected %q, got %q", i, tt.expectedLiteral, tok.Literal)
    }
  }

}

