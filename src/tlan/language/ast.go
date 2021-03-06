package language

type Node interface {
	TokenLiteral() string
}

type
Item struct {
	Token       Token
	Type        string
	Name        *Name
	Description *Description
	Category    *Category
	Children    []*Item
	Marked      bool
	Annotations []Annotation
	Target      string
}

func (s *Item) TokenLiteral() string { return s.Token.Literal }

const (
	Task    = "Task"
	Project = "Project"
	Pointer = "Pointer"
)

type Name struct {
	Token Token
	Value string
}

func (s *Name) TokenLiteral() string { return s.Token.Literal }

type Description struct {
	Token Token
	Value string
}

func (s *Description) TokenLiteral() string { return s.Token.Literal }

type Category struct {
	Token       Token
	Value       string
	Annotations []Annotation
}

func (i *Category) TokenLiteral() string { return i.Token.Literal }

type Annotation interface {
	Type() string
	ToString() string
}

const (
	UNARY  = "UNARY"
	BINARY = "BINARY"
)

type UnaryAnnotation struct {
	Token Token
	Name  Name
}

func (i *UnaryAnnotation) TokenLiteral() string { return i.Token.Literal }
func (i *UnaryAnnotation) Type() string         { return UNARY }
func (i *UnaryAnnotation) ToString() string     { return i.Name.Value }

type BinaryAnnotation struct {
	Token    Token
	Left     *Name
	Right    *Name
	Operator *Operator
}

func (i *BinaryAnnotation) TokenLiteral() string { return i.Token.Literal }
func (i *BinaryAnnotation) Type() string         { return BINARY }
func (i *BinaryAnnotation) ToString() string {
	return i.Left.Value + i.Operator.TokenLiteral() + i.Right.Value
}

type Operator struct {
	Token Token
}

func (i *Operator) TokenLiteral() string { return i.Token.Literal }
