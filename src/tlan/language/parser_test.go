package language

import "testing"

func TestTreeCreation(t *testing.T) {
	input := `
AI
- Math
  - Bachelors Degree 
- Foundations [Unary, 1501-0612]
* Books
- (Research)
`
	cases := []struct {
		itemType    string
		name        string
		description string
		level       int
		category    string
		children    int
		parenthesis bool
		annotations []string
	}{
		{"-", "Math", "", 1, "AI", 1, false, []string{}},
		{"-", "Foundations", "", 1, "AI", 0, false, []string{"Unary", "1501-0612"}},
		{"*", "Books", "", 1, "AI", 0, false, []string{}},
		{"-", "Research", "", 1, "AI", 0, true, []string{}},
	}

	lexer := NewLexer(input)
	parser := NewParser(lexer)
	items := parser.Parse()

	for i, tt := range cases {
		if items[i].Type.TokenLiteral() != tt.itemType {
			t.Fatalf("Expecting %s got %s", tt.itemType, items[i].Type.TokenLiteral())
		}

		if items[i].Name.TokenLiteral() != tt.name {
			t.Fatalf("Expecting %s got %s", tt.name, items[i].Name.TokenLiteral())
		}

		if items[i].Description != nil && items[i].Description.TokenLiteral() != tt.description {
			t.Fatalf("Expecting %s got %s", tt.description, items[i].Description.TokenLiteral())
		}

		name := items[i].Name.Value
		if items[i].Category.TokenLiteral() != tt.category {
			t.Fatalf("Expecting %s got %s, for item %s", tt.category, items[i].Category.TokenLiteral(), name)
		}

		if len(items[i].Children) != tt.children {
			t.Fatalf("Expecting children count %d got %d, for item %s", tt.children, len(items[i].Children), name)
		}

		if items[i].Marked != tt.parenthesis {
			t.Fatalf("Expecting %v got %v, for item %s", tt.parenthesis, items[i].Marked, name)
		}

		if len(items[i].Annotations) != len(tt.annotations) {
			t.Fatalf("Expecting annotations count to be %d got %d, for item %s", len(tt.annotations), len(items[i].Annotations), name)
		}
	}
}
