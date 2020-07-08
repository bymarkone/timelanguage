package parser

import "testing"

func TestTreeCreation(t *testing.T) {
  input := `
- change the world: "Because the world needs to be changed"
  - what to build: "I cannot understand what I cannot build"
  - what to say: "Words will always retain their power" 
`

  cases := []struct {
    project       string
    description   string
    level         int
  }{
    { "change the world", "Because the world needs to be changed", 1 },
    { "what to build", "I cannot understand what I cannot build", 2 },
    { "what to say", "Words will always retain their power", 2 },
  }

  lexer := NewLexer(input)
  parser := NewParser(lexer)
  valuable := parser.Parse()
  valuables := append([]*Valuable{valuable}, valuable.Children...)

  for i, tt := range cases {
    if valuables[i].Name.TokenLiteral() != tt.project {
      t.Fatalf("Expecting %s got %s", tt.project, valuables[i].Name.TokenLiteral())
    }

    if valuables[i].Description.TokenLiteral() != tt.description {
      t.Fatalf("Expecting %s got %s", tt.description, valuables[i].Description.TokenLiteral())
    }
  }
}
