package ast

import "strings"

type Stmt interface {
	Node
	CopyStmt() Stmt
	stmtNode()
}

type (

	// BlockStmt represents a brace statement list
	BlockStmt struct {
		List []Node
		Pos
	}

	// AssignStmt represents assigning of a value to a previously declared variable
	AssignStmt struct {
		Pos
		Ident Token
		Val   Expr
	}

	// LetStmt represents a variable assignment. Number only
	LetStmt struct {
		Pos
		Ident Token
		Val   Expr
	}

	// ReturnStmt represents return <expr>.
	ReturnStmt struct {
		Pos
		Val Expr
	}

	// CtrlStmt represents continue, break, and fallthrough
	CtrlStmt struct {
		Pos
		Typ CtrlTyp
	}

	// IfStmt represents an if block
	IfStmt struct {
		Pos
		Condition Expr       //TODO: this should be an expr?
		IfBlock   *BlockStmt // What to execute if true
		ElseBlock Stmt       // What to execute if false
	}

	// SwitchStmt represent a switch block
	SwitchStmt struct {
		Pos
		Condition Expr // the condition to switch on
		Cases     []*CaseStmt
		Default   *BlockStmt // default case
	}

	// CaseStmt represents a case in a switch block
	CaseStmt struct {
		Pos
		Condition Expr
		Body      *BlockStmt
	}

	// A FnStmt node represents a function.
	FnStmt struct {
		Pos
		FunVal Token
		Args   []*Ident
		Body   *BlockStmt
	}

	// WhileStmt represents a while block
	WhileStmt struct {
		Pos
		Condition  Expr       //TODO: this should be an expr?
		WhileBlock *BlockStmt // What to execute if true
	}

	// ForStmt represents a for block
	ForStmt struct {
		Pos
		Init Stmt // initialization statement; or nil
		Cond Expr // condition; or nil
		Post Stmt // post iteration statement; or nil
		Body *BlockStmt
	}
)

type CtrlTyp int

const (
	InvalidCtrl CtrlTyp = iota
	CtrlBreak
	CtrlContinue
	CtrlFallthrough
)

// stmtNode()
func (*BlockStmt) stmtNode()  {}
func (*AssignStmt) stmtNode() {}
func (*LetStmt) stmtNode()    {}
func (*CtrlStmt) stmtNode()   {}
func (*ReturnStmt) stmtNode() {}
func (*IfStmt) stmtNode()     {}
func (*SwitchStmt) stmtNode() {}
func (*CaseStmt) stmtNode()   {}
func (*FnStmt) stmtNode()     {}
func (*WhileStmt) stmtNode()  {}
func (*ForStmt) stmtNode()    {}

// BlockStmt.
func NewBlockStmt(pos Pos) *BlockStmt {
	return &BlockStmt{Pos: pos}
}

func (b *BlockStmt) Append(n Node) {
	b.List = append(b.List, n)
}

func (b *BlockStmt) String() string {
	var sb strings.Builder
	b.writeTo(&sb)
	return sb.String()
}

func (b *BlockStmt) writeTo(sb *strings.Builder) {
	for _, n := range b.List {
		n.writeTo(sb)
		sb.WriteString("\n")
	}
}

func (b *BlockStmt) CopyBlock() *BlockStmt {
	if b == nil {
		return b
	}
	n := NewBlockStmt(b.Pos)
	for _, elem := range b.List {
		n.Append(elem.Copy())
	}
	return n
}

func (b *BlockStmt) CopyStmt() Stmt {
	return b.CopyBlock()
}

func (b *BlockStmt) Copy() Node {
	return b.CopyBlock()
}

// AssignStmt.

func (a *AssignStmt) String() string {
	var sb strings.Builder
	a.writeTo(&sb)
	return sb.String()
}

func (a *AssignStmt) writeTo(sb *strings.Builder) {
	sb.WriteString(a.Ident.String())
	sb.WriteString(" = ")
	a.Val.writeTo(sb)
}

func (a *AssignStmt) CopyAssign() *AssignStmt {
	if a == nil {
		return a
	}
	n := &AssignStmt{
		Pos:   a.Pos,
		Ident: a.Ident,
	}
	n.Val = a.Val.CopyExpr()
	return n
}

func (a *AssignStmt) CopyStmt() Stmt {
	return a.CopyAssign()
}

func (a *AssignStmt) Copy() Node {
	return a.CopyAssign()
}

// LetStmt.

func (l *LetStmt) String() string {
	var sb strings.Builder
	l.writeTo(&sb)
	return sb.String()
}

func (l *LetStmt) writeTo(sb *strings.Builder) {
	sb.WriteString("let ")
	sb.WriteString(l.Ident.String())
	sb.WriteString(" = ")
	if l.Val != nil {
		l.Val.writeTo(sb)
	}
}

func (l *LetStmt) CopyLet() *LetStmt {
	if l == nil {
		return l
	}
	n := &LetStmt{
		Pos:   l.Pos,
		Ident: l.Ident,
	}
	n.Val = l.Val.CopyExpr()
	return n
}

func (l *LetStmt) CopyStmt() Stmt {
	return l.CopyLet()
}

func (l *LetStmt) Copy() Node {
	return l.CopyLet()
}

// ReturnStmt.

func (r *ReturnStmt) String() string {
	var sb strings.Builder
	r.writeTo(&sb)
	return sb.String()
}

func (r *ReturnStmt) writeTo(sb *strings.Builder) {
	sb.WriteString("return ")
	r.Val.writeTo(sb)
}

func (r *ReturnStmt) CopyReturn() *ReturnStmt {
	if r == nil {
		return r
	}
	n := &ReturnStmt{
		Pos: r.Pos,
	}
	n.Val = r.Val.CopyExpr()
	return n
}

func (r *ReturnStmt) CopyStmt() Stmt {
	return r.CopyReturn()
}

func (r *ReturnStmt) Copy() Node {
	return r.CopyReturn()
}

// CtrlStmt.

func (c *CtrlStmt) String() string {
	var sb strings.Builder
	c.writeTo(&sb)
	return sb.String()
}

func (c *CtrlStmt) writeTo(sb *strings.Builder) {
	switch c.Typ {
	case CtrlContinue:
		sb.WriteString("continue")
	case CtrlBreak:
		sb.WriteString("break")
	case CtrlFallthrough:
		sb.WriteString("fallthrough")
	}
}

func (c *CtrlStmt) CopyControl() *CtrlStmt {
	if c == nil {
		return c
	}
	n := &CtrlStmt{
		Pos: c.Pos,
		Typ: c.Typ,
	}
	return n
}

func (c *CtrlStmt) CopyStmt() Stmt {
	return c.CopyControl()
}

func (c *CtrlStmt) Copy() Node {
	return c.CopyControl()
}

// IfStmt.

func (i *IfStmt) String() string {
	var sb strings.Builder
	i.writeTo(&sb)
	return sb.String()
}

func (i *IfStmt) writeTo(sb *strings.Builder) {
	sb.WriteString("if ")
	i.Condition.writeTo(sb)
	sb.WriteString(" {\n")
	i.IfBlock.writeTo(sb)
	sb.WriteString("}")
	if i.ElseBlock != nil {
		sb.WriteString("else {\n")
		sb.WriteString(i.ElseBlock.String())
		sb.WriteString("}")
	}
}

func (i *IfStmt) CopyIfStmt() *IfStmt {
	if i == nil {
		return nil
	}
	n := &IfStmt{
		Pos:       i.Pos,
		Condition: i.Condition.CopyExpr(),
		IfBlock:   i.IfBlock.CopyBlock(),
	}
	if i.ElseBlock != nil {
		n.ElseBlock = i.ElseBlock.CopyStmt()
	}
	return n
}

func (i *IfStmt) CopyStmt() Stmt {
	return i.CopyIfStmt()
}

func (i *IfStmt) Copy() Node {
	return i.CopyIfStmt()
}

// SwitchStmt.

func (s *SwitchStmt) String() string {
	var sb strings.Builder
	s.writeTo(&sb)
	return sb.String()
}

func (s *SwitchStmt) writeTo(sb *strings.Builder) {
	sb.WriteString("switch ")
	s.Condition.writeTo(sb)
	sb.WriteString(" {\n")
	for _, v := range s.Cases {
		v.writeTo(sb)
	}
	if s.Default != nil {
		sb.WriteString("default: {\n")
		s.Default.writeTo(sb)
		sb.WriteString("}")
	}
	sb.WriteString("}")
}

func (s *SwitchStmt) CopySwitch() *SwitchStmt {
	if s == nil {
		return nil
	}
	n := &SwitchStmt{
		Pos:     s.Pos,
		Cases:   make([]*CaseStmt, 0, len(s.Cases)),
		Default: s.Default.CopyBlock(),
	}
	if s.Condition != nil {
		n.Condition = s.Condition.CopyExpr()
	}
	for i := range s.Cases {
		n.Cases = append(n.Cases, s.Cases[i].CopyCase())
	}
	return n
}

func (s *SwitchStmt) CopyStmt() Stmt {
	return s.CopySwitch()
}

func (s *SwitchStmt) Copy() Node {
	return s.CopySwitch()
}

// CaseStmt.

func (c *CaseStmt) String() string {
	var sb strings.Builder
	c.writeTo(&sb)
	return sb.String()
}

func (c *CaseStmt) writeTo(sb *strings.Builder) {
	sb.WriteString("case ")
	c.Condition.writeTo(sb)
	sb.WriteString(" {\n")
	c.Body.writeTo(sb)
	sb.WriteString("}")
}

func (c *CaseStmt) CopyCase() *CaseStmt {
	if c == nil {
		return nil
	}
	return &CaseStmt{
		Pos:       c.Pos,
		Condition: c.Condition.CopyExpr(),
		Body:      c.Body.CopyBlock(),
	}
}

func (c *CaseStmt) CopyStmt() Stmt {
	return c.CopyCase()
}

func (c *CaseStmt) Copy() Node {
	return c.CopyCase()
}

// FnStmt.

func (f *FnStmt) CopyFn() Stmt {
	if f == nil {
		return nil
	}
	n := &FnStmt{
		Pos:    f.Pos,
		FunVal: f.FunVal,
		Body:   f.Body.CopyBlock(),
		Args:   make([]*Ident, 0, len(f.Args)),
	}
	for i := range f.Args {
		n.Args = append(n.Args, f.Args[i].CopyIdent())
	}

	return n
}

func (f *FnStmt) CopyStmt() Stmt {
	return f.CopyFn()
}

func (f *FnStmt) Copy() Node {
	return f.CopyStmt()
}

func (f *FnStmt) String() string {
	var sb strings.Builder
	f.writeTo(&sb)
	return sb.String()
}

func (f *FnStmt) writeTo(sb *strings.Builder) {
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

// WhileStmt.

func (w *WhileStmt) String() string {
	var sb strings.Builder
	w.writeTo(&sb)
	return sb.String()
}

func (w *WhileStmt) writeTo(sb *strings.Builder) {
	sb.WriteString("while ")
	w.Condition.writeTo(sb)
	sb.WriteString(" {\n")
	w.WhileBlock.writeTo(sb)
	sb.WriteString("}")
}

func (w *WhileStmt) CopyWhileStmt() *WhileStmt {
	if w == nil {
		return nil
	}
	return &WhileStmt{
		Pos:        w.Pos,
		Condition:  w.Condition.CopyExpr(),
		WhileBlock: w.WhileBlock.CopyBlock(),
	}
}

func (w *WhileStmt) CopyStmt() Stmt {
	return w.CopyWhileStmt()
}

func (w *WhileStmt) Copy() Node {
	return w.CopyWhileStmt()
}

// ForStmt.

func (f *ForStmt) String() string {
	var sb strings.Builder
	f.writeTo(&sb)
	return sb.String()
}

func (f *ForStmt) writeTo(sb *strings.Builder) {
	sb.WriteString("for ")
	if f.Init != nil {
		f.Init.writeTo(sb)
	}
	sb.WriteString("; ")
	if f.Cond != nil {
		f.Cond.writeTo(sb)
	}
	sb.WriteString("; ")
	if f.Post != nil {
		f.Post.writeTo(sb)
	}
	sb.WriteString(" {\n")
	f.Body.writeTo(sb)
	sb.WriteString("}")
}

func (f *ForStmt) CopyForStmt() *ForStmt {
	if f == nil {
		return nil
	}
	n := &ForStmt{
		Pos:  f.Pos,
		Body: f.Body.CopyBlock(),
	}
	if f.Init != nil {
		n.Init = f.Init.CopyStmt()
	}
	if f.Cond != nil {
		n.Cond = f.Cond.CopyExpr()
	}
	if f.Post != nil {
		n.Post = f.Post.CopyStmt()
	}
	return n
}

func (f *ForStmt) CopyStmt() Stmt {
	return f.CopyForStmt()
}

func (f *ForStmt) Copy() Node {
	return f.CopyForStmt()
}
