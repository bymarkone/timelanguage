package parser

import "testing"

func TestTreeCreation(t *testing.T) {
	input := `
AI
- Math
  - Bachelors Degree
- Foundations
- Books
- Research
`
	cases := []struct {
		item        string
		description string
		level       int
		category    string
		children 		int
	}{
		{"Math", "", 1, "AI", 1},
		{"Foundations", "", 1, "AI", 0},
		{"Books", "", 1, "AI", 0},
		{"Research", "", 1, "AI", 0},
	}

	lexer := NewLexer(input)
	parser := NewParser(lexer)
	items := parser.Parse()

	for i, tt := range cases {
		if items[i].Name.TokenLiteral() != tt.item {
			t.Fatalf("Expecting %s got %s", tt.item, items[i].Name.TokenLiteral())
		}

		if items[i].Description != nil && items[i].Description.TokenLiteral() != tt.description {
			t.Fatalf("Expecting %s got %s", tt.description, items[i].Description.TokenLiteral())
		}

		if items[i].Category.TokenLiteral() != tt.category {
			t.Fatalf("Expecting %s got %s, for item %s", tt.category, items[i].Category.TokenLiteral(), items[i].Name.Value)
		}

		if len(items[i].Children) != tt.children {
			t.Fatalf("Expecting children count %d got %d, for item %s", tt.children, len(items[i].Children), items[i].Name.Value)
		}
	}
}
