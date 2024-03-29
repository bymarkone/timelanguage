package language

import "testing"

func TestTreeCreation(t *testing.T) {
	input := `
First Category
AI
- Math >> Mathematician
  - Bachelors Degree 
- Foundations [Unary. 15/01/21-06/12/22]
* Books [Unary. 05:00-07:00]
- Parent :: Child
- (Research)
  + Follow another list, but not too "eagerly"
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
		target      string
		preName     string
	}{
		{"Normal", "Math", "", 1, "AI", 1, false, []string{}, "Mathematician", ""},
		{"Normal", "Foundations", "", 1, "AI", 0, false, []string{"Unary", "15/01/21-06/12/22"}, "", ""},
		{"Star", "Books", "", 1, "AI", 0, false, []string{"Unary", "05:00-07:00"}, "", ""},
		{"Normal", "Child", "", 1, "AI", 0, false, []string{}, "", "Parent"},
		{"Normal", "Research", "", 1, "AI", 1, true, []string{}, "", ""},
	}

	lexer := NewLexer(input)
	parser := NewParser("test", lexer)
	categories, items := parser.Parse()

	if categories[0].Value != "First Category" {
		t.Fatalf("Expecting %s got %s", "First Category", categories[0].Value)
	}

	for i, tt := range cases {
		if items[i].Type != tt.itemType {
			t.Fatalf("Expecting %s got %s", tt.itemType, items[i].Type)
		}

		if items[i].Name.TokenLiteral() != tt.name {
			t.Fatalf("Expecting %s got %s", tt.name, items[i].Name.TokenLiteral())
		}

		if items[i].PreName != nil && items[i].PreName.TokenLiteral() != tt.preName {
			t.Fatalf("Expecting %s got %s", tt.preName, items[i].PreName.TokenLiteral())
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

		for j, _ := range items[i].Annotations {
			if items[i].Annotations[j].ToString() != tt.annotations[j] {
				t.Fatalf("Expecting annotations to be %s got %s, for item %s", tt.annotations[j], items[i].Annotations[j].ToString(), name)
			}
		}

		if items[i].Target != tt.target {
			t.Fatalf("Expecting %s got %s, for item target %s", tt.target, items[i].Target, name)
		}

	}
}
