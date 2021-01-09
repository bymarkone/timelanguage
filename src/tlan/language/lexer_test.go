package language

import "testing"

func TestAnything(t *testing.T) {
	input := `
AI
- Math
  - Bachelors Degree [1501-1012]
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
		{DASH, "-"},
		{IDENT, "Math"},
		{LEVEL, "  "},
		{DASH, "-"},
		{IDENT, "Bachelors Degree"},
		{LSB, "["},
		{IDENT, "1501"},
		{DASH, "-"},
		{IDENT, "1012"},
		{RSB, "]"},
		{LEVEL, "  "},
		{DASH, "-"},
		{LP, "("},
		{IDENT, "Cambridge Part III"},
		{RP, ")"},
		{DASH, "-"},
		{IDENT, "Foundations"},
		{DASH, "-"},
		{IDENT, "Books"},
		{DASH, "-"},
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
