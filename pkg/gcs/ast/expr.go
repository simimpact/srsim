package ast

import (
	"strconv"
	"strings"
)

type Expr interface {
	Node
	exprNode()
	CopyExpr() Expr
}

// An expression is represented by a tree consisting of one or
// more of the following concrete expression nodes
type (
	NumberLit struct {
		Pos
		IntVal   int64
		FloatVal float64
		IsFloat  bool
	}

	StringLit struct {
		Pos
		Value string
	}

	BoolLit struct {
		Pos
		Value bool
	}

	NullLit struct {
		Pos
	}

	// A FuncLit node represents a function literal.
	FuncLit struct {
		Pos
		Args []*Ident
		Body *BlockStmt
	}

	Ident struct {
		Pos
		Value string
	}

	// A CallExpr node represents an expression followed by an argument list.
	CallExpr struct {
		Pos
		Fun  Expr   // function expression
		Args []Expr // function arguments; or nil
	}

	// A UnaryExpr node represents a unary expression.
	UnaryExpr struct {
		Pos
		Op    Token
		Right Expr // operand
	}

	//A BinaryExpr node represents a binary expression i.e. a > b, 1 + 1, etc..
	BinaryExpr struct {
		Pos
		Left  Expr
		Right Expr  // need to evalute to same type as lhs
		Op    Token //should be > itemCompareOP and < itemDot
	}

	MapExpr struct {
		Pos
		Fields map[string]Expr
	}
)

// exprNode()
func (*NumberLit) exprNode()  {}
func (*StringLit) exprNode()  {}
func (*BoolLit) exprNode()    {}
func (*NullLit) exprNode()    {}
func (*FuncLit) exprNode()    {}
func (*Ident) exprNode()      {}
func (*CallExpr) exprNode()   {}
func (*UnaryExpr) exprNode()  {}
func (*BinaryExpr) exprNode() {}
func (*MapExpr) exprNode()    {}

// NumberLit.

func (n *NumberLit) CopyExpr() Expr {
	if n == nil {
		return nil
	}
	return &NumberLit{
		Pos:      n.Pos,
		IntVal:   n.IntVal,
		FloatVal: n.FloatVal,
		IsFloat:  n.IsFloat,
	}
}

func (n *NumberLit) Copy() Node {
	return n.CopyExpr()
}

func (n *NumberLit) String() string {
	var sb strings.Builder
	n.writeTo(&sb)
	return sb.String()
}

func (n *NumberLit) writeTo(sb *strings.Builder) {
	if n.IsFloat {
		sb.WriteString(strconv.FormatFloat(n.FloatVal, 'f', -1, 64))
	} else {
		sb.WriteString(strconv.FormatInt(n.IntVal, 10))
	}
}

// StringLit.

func (n *StringLit) CopyExpr() Expr {
	if n == nil {
		return nil
	}
	return &StringLit{
		Pos:   n.Pos,
		Value: n.Value,
	}
}

func (n *StringLit) Copy() Node {
	return n.CopyExpr()
}

func (n *StringLit) String() string {
	return n.Value
}

func (n *StringLit) writeTo(sb *strings.Builder) {
	sb.WriteString(n.Value)
}

// BoolLit.

func (b *BoolLit) CopyExpr() Expr {
	if b == nil {
		return nil
	}
	return &BoolLit{
		Pos:   b.Pos,
		Value: b.Value,
	}
}

func (b *BoolLit) Copy() Node {
	return b.CopyExpr()
}

func (b *BoolLit) String() string {
	var sb strings.Builder
	b.writeTo(&sb)
	return sb.String()
}

func (b *BoolLit) writeTo(sb *strings.Builder) {
	if b.Value {
		sb.WriteString("true")
	} else {
		sb.WriteString("false")
	}
}

// BoolLit.

func (n *NullLit) CopyExpr() Expr {
	if n == nil {
		return nil
	}
	return &NullLit{Pos: n.Pos}
}

func (n *NullLit) Copy() Node {
	return n.CopyExpr()
}

func (n *NullLit) String() string {
	var sb strings.Builder
	n.writeTo(&sb)
	return sb.String()
}

func (n *NullLit) writeTo(sb *strings.Builder) {
	sb.WriteString("null")
}

// FuncLit.

func (f *FuncLit) CopyExpr() Expr {
	if f == nil {
		return nil
	}
	n := &FuncLit{
		Pos:  f.Pos,
		Args: make([]*Ident, 0, len(f.Args)),
		Body: f.Body.CopyBlock(),
	}
	for i := range f.Args {
		n.Args = append(n.Args, f.Args[i].CopyIdent())
	}
	return n
}

func (f *FuncLit) Copy() Node {
	return f.CopyExpr()
}

func (f *FuncLit) String() string {
	var sb strings.Builder
	f.writeTo(&sb)
	return sb.String()
}

func (f *FuncLit) writeTo(sb *strings.Builder) {
	sb.WriteString("fn(")
	for i, v := range f.Args {
		if i > 0 {
			sb.WriteString(", ")
		}
		v.writeTo(sb)
	}
	sb.WriteString(") {\n")
	f.Body.writeTo(sb)
	sb.WriteString("}")
}

// Ident.

func (i *Ident) CopyIdent() *Ident {
	if i == nil {
		return nil
	}
	return &Ident{Pos: i.Pos, Value: i.Value}
}

func (i *Ident) CopyExpr() Expr {
	return i.CopyIdent()
}

func (i *Ident) Copy() Node {
	return i.CopyIdent()
}

func (b *Ident) String() string {
	var sb strings.Builder
	b.writeTo(&sb)
	return sb.String()
}

func (b *Ident) writeTo(sb *strings.Builder) {
	sb.WriteString(b.Value)
}

// CallExpr.

func (c *CallExpr) CopyFn() Expr {
	if c == nil {
		return nil
	}
	n := &CallExpr{
		Pos:  c.Pos,
		Fun:  c.Fun.CopyExpr(),
		Args: make([]Expr, 0, len(c.Args)),
	}
	for i := range c.Args {
		n.Args = append(n.Args, c.Args[i].CopyExpr())
	}

	return n
}

func (f *CallExpr) CopyExpr() Expr {
	return f.CopyFn()
}

func (f *CallExpr) Copy() Node {
	return f.CopyExpr()
}

func (f *CallExpr) String() string {
	var sb strings.Builder
	f.writeTo(&sb)
	return sb.String()
}

func (b *CallExpr) writeTo(sb *strings.Builder) {
	b.Fun.writeTo(sb)
	sb.WriteString("(")
	for i, v := range b.Args {
		if i > 0 {
			sb.WriteString(", ")
		}
		v.writeTo(sb)
	}
	sb.WriteString(")")
}

// UnaryExpr.

func (u *UnaryExpr) CopyUnaryExpr() *UnaryExpr {
	if u == nil {
		return u
	}
	n := &UnaryExpr{Pos: u.Pos}
	n.Right = u.Right.CopyExpr()
	n.Op = u.Op
	return n
}

func (u *UnaryExpr) CopyExpr() Expr {
	return u.CopyUnaryExpr()
}

func (u *UnaryExpr) Copy() Node {
	return u.CopyUnaryExpr()
}

func (u *UnaryExpr) String() string {
	var sb strings.Builder
	u.writeTo(&sb)
	return sb.String()
}

func (u *UnaryExpr) writeTo(sb *strings.Builder) {
	sb.WriteString("(")
	sb.WriteString(u.Op.String())
	u.Right.writeTo(sb)
	sb.WriteString(")")
}

// BinaryExpr.

func (b *BinaryExpr) CopyBinaryExpr() *BinaryExpr {
	if b == nil {
		return b
	}
	n := &BinaryExpr{Pos: b.Pos}
	n.Left = b.Left.CopyExpr()
	n.Right = b.Right.CopyExpr()
	n.Op = b.Op
	return n
}

func (b *BinaryExpr) CopyExpr() Expr {
	return b.CopyBinaryExpr()
}

func (b *BinaryExpr) Copy() Node {
	return b.CopyBinaryExpr()
}

func (b *BinaryExpr) String() string {
	var sb strings.Builder
	b.writeTo(&sb)
	return sb.String()
}

func (b *BinaryExpr) writeTo(sb *strings.Builder) {
	sb.WriteString("(")
	b.Left.writeTo(sb)
	sb.WriteString(" ")
	sb.WriteString(b.Op.String())
	sb.WriteString(" ")
	b.Right.writeTo(sb)
	sb.WriteString(")")
}

// MapExpr.

func (m *MapExpr) CopyExpr() Expr {
	if m == nil {
		return m
	}
	n := &MapExpr{
		Pos:    m.Pos,
		Fields: make(map[string]Expr),
	}
	for k, v := range m.Fields {
		n.Fields[k] = v.CopyExpr()
	}
	return n
}

func (m *MapExpr) Copy() Node {
	return m.CopyExpr()
}

func (m *MapExpr) String() string {
	var sb strings.Builder
	m.writeTo(&sb)
	return sb.String()
}

func (m *MapExpr) writeTo(sb *strings.Builder) {
	sb.WriteString("[")
	done := false
	for k, v := range m.Fields {
		if done {
			sb.WriteString(", ")
		}
		done = true

		sb.WriteString(k)
		sb.WriteString(" = ")
		sb.WriteString(v.String())
	}
	sb.WriteString("]")
}
