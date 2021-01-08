package parser

import "testing"

func TestAnything(t *testing.T) {
	input := `
AI
- Math
  - Bachelors Degree
  - (Cambridge Part III)
- Foundations
- Books
- (Research)
`
	cases := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{IDENT, "AI"},
		{ITEM, "-"},
		{IDENT, "Math"},
		{LEVEL, "  "},
		{ITEM, "-"},
		{IDENT, "Bachelors Degree"},
		{LEVEL, "  "},
		{ITEM, "-"},
		{LP, "("},
		{IDENT, "Cambridge Part III"},
		{RP, ")"},
		{ITEM, "-"},
		{IDENT, "Foundations"},
		{ITEM, "-"},
		{IDENT, "Books"},
		{ITEM, "-"},
		{LP, "("},
		{IDENT, "Research"},
		{RP, ")"},
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
