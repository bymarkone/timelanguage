package parser

import "testing"

func TestAnything(t *testing.T) {
  input := `
- Change the world
  - What I build
  - How I live
`

  cases := []struct {
    expectedType    TokenType
    expectedLiteral string
  }{
    {ITEM, "-"},
    {IDENT, "Change the world"},
    {FIRST, "  "},
    {ITEM, "-"},
    {IDENT, "What I build"},
    {FIRST, "  "},
    {ITEM, "-"},
    {IDENT, "How I live"},
  }

  lexer := New(input)

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

