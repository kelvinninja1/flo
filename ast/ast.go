package ast

type Noder interface {
	String() string
}

// CURRENT AST NODES
// ast.Node
// ast.List
// ast.Statement
// ast.UnaryOp
// ast.BinaryOp
// ast.Function
// ast.Block

// ast.Node
type Node struct {
	Type  string
	Value string
}

func (n *Node) String() string {
	return n.Value
}

// ast.List
type List struct {
	Values []Noder
}

func (l *List) String() string {
	values := "["
	if l.Values != nil {
		for _, value := range l.Values {
			values = values + value.String() + ", "
		}
		values = values[:len(values)-2]
	}
	values += "]"
	return values
}

// ast.Statement
type Statement struct {
	Value Noder
	Next  Noder
}

func (s *Statement) String() string {
	value := "<nil>"
	next := "<nil>"
	if s.Value != nil {
		value = s.Value.String()
	}
	if s.Next != nil {
		next = s.Next.String()
	}
	return value + " -> " + next
}

// ast.UnaryOp
type UnaryOp struct {
	Operator string
	Value    Noder
}

func (u *UnaryOp) String() string {

	value := "<nil>"

	if u.Value != nil {
		value = u.Value.String()
	}

	return "( " + u.Operator + " " + value + " )"
}

// ast.BinaryOp
type BinaryOp struct {
	Operator string
	Left     Noder
	Right    Noder
}

func (b *BinaryOp) String() string {

	left := "<nil>"
	right := "<nil>"

	if b.Left != nil {
		left = b.Left.String()
	}
	if b.Right != nil {
		right = b.Right.String()
	}

	return "( " + b.Operator + " " + left + " " + right + " )"
}

// ast.Function
type Function struct {
	Name       string
	Parameters List
	Body       Noder
}

func (f *Function) String() string {

	body := "<nil>"

	if f.Body != nil {
		body = f.Body.String()
	}

	return "( " + f.Name + " " + f.Parameters.String() + " " + body + " )"
}

// ast.Block
type Block struct {
	Value Noder
}

func (b *Block) String() string {
	v := ""
	if b.Value != nil {
		v = b.Value.String()
	}
	return "{ " + v + " }"
}
