package parser

import "testing"

func TestTreeCreation(t *testing.T) {
	input := `
AI
- Math
- Foundations
- Books
- Research
`
	cases := []struct {
		item        string
		description string
		level       int
	}{
		{"AI", "", 0},
		{"Math", "", 1},
		{"Foundations", "", 1},
		{"Books", "", 1},
		{"Research", "", 1},
	}

	lexer := NewLexer(input)
	parser := NewParser(lexer)
	item := parser.Parse()
	items := append([]*Item{item}, item.Children...)

	for i, tt := range cases {
		if items[i].Name.TokenLiteral() != tt.item {
			t.Fatalf("Expecting %s got %s", tt.item, items[i].Name.TokenLiteral())
		}

		if items[i].Description != nil && items[i].Description.TokenLiteral() != tt.description {
			t.Fatalf("Expecting %s got %s", tt.description, items[i].Description.TokenLiteral())
		}

		if items[i].Level != tt.level {
			t.Fatalf("Expecting %d got %d, for item %s", tt.level, items[i].Level, items[i].Name.Value)
		}
	}
}
