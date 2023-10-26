package language

import (
	"testing"
)

func TestAnything(t *testing.T) {
	input := `
First Category
AI
- Math >> Mathematician
  - Bachelors Degree [Unary. 15/01/21-10/12/22]
  - (Cambridge Part III)
- Foundations [Unary. 05:00-10:00]
- Parent :: Child
* Books 
- (Research)
  + Follow other list, but not too "eagerly"
`
	cases := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{IDENT, "First Category"},
		{IDENT, "AI"},
		{DASH, "-"},
		{IDENT, "Math"},
		{DUALARROW, ">>"},
		{IDENT, "Mathematician"},
		{LEVEL, "  "},
		{DASH, "-"},
		{IDENT, "Bachelors Degree"},
		{LSB, "["},
		{IDENT, "Unary"},
		{DOT, "."},
		{IDENT, "15/01/21"},
		{DASH, "-"},
		{IDENT, "10/12/22"},
		{RSB, "]"},
		{LEVEL, "  "},
		{DASH, "-"},
		{LP, "("},
		{IDENT, "Cambridge Part III"},
		{RP, ")"},
		{DASH, "-"},
		{IDENT, "Foundations"},
		{LSB, "["},
		{IDENT, "Unary"},
		{DOT, "."},
		{IDENT, "05:00"},
		{DASH, "-"},
		{IDENT, "10:00"},
		{RSB, "]"},
		{DASH, "-"},
		{IDENT, "Parent"},
		{DUALCOLON, "::"},
		{IDENT, "Child"},
		{STAR, "*"},
		{IDENT, "Books"},
		{DASH, "-"},
		{LP, "("},
		{IDENT, "Research"},
		{RP, ")"},
		{LEVEL, "  "},
		{PLUS, "+"},
		{IDENT, "Follow other list, but not too \"eagerly\""},
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
