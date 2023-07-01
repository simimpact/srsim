package parse

import (
	"fmt"
	"strconv"

	"github.com/simimpact/srsim/pkg/gcs"
	"github.com/simimpact/srsim/pkg/gcs/ast"
)

type precedence int

const (
	_ precedence = iota
	Lowest
	LogicalOr
	LogicalAnd // TODO: or make one for && and ||?
	Equals
	LessOrGreater
	Sum
	Product
	Prefix
	Call
)

var precedences = map[ast.TokenType]precedence{
	ast.LogicOr:              LogicalOr,
	ast.LogicAnd:             LogicalAnd,
	ast.OpEqual:              Equals,
	ast.OpNotEqual:           Equals,
	ast.OpLessThan:           LessOrGreater,
	ast.OpGreaterThan:        LessOrGreater,
	ast.OpLessThanOrEqual:    LessOrGreater,
	ast.OpGreaterThanOrEqual: LessOrGreater,
	ast.ItemPlus:             Sum,
	ast.ItemMinus:            Sum,
	ast.ItemForwardSlash:     Product,
	ast.ItemAsterisk:         Product,
	ast.ItemLeftParen:        Call,
}

func tokenPrecendence(t ast.Token) precedence {
	if p, ok := precedences[t.Typ]; ok {
		return p
	}
	return Lowest
}

// Parse returns the ActionList and any error that prevents the ActionList from being parsed
func (p *Parser) Parse() (*gcs.ActionList, error) {
	var err error
	for state := parseRows; state != nil; {
		state, err = state(p)
		if err != nil {
			return nil, err
		}
	}

	// build the err msgs
	p.res.ErrorMsgs = make([]string, 0, len(p.res.Errors))
	for _, v := range p.res.Errors {
		p.res.ErrorMsgs = append(p.res.ErrorMsgs, v.Error())
	}

	return p.res, nil
}

func parseRows(p *Parser) (parseFn, error) {
	switch n := p.peek(); n.Typ {
	case ast.ItemEOF:
		return nil, nil
	default: // default should be look for gcsl
		node, err := p.parseStatement()
		p.res.Program.Append(node)
		if err != nil {
			return nil, err
		}
		return parseRows, nil
	}
}

func (p *Parser) parseStatement() (ast.Node, error) {
	// some statements end in semi, other don't
	hasSemi := true
	stmtType := ""
	var node ast.Node
	var err error
	switch n := p.peek(); n.Typ {
	case ast.KeywordBreak:
		fallthrough
	case ast.KeywordFallthrough:
		fallthrough
	case ast.KeywordContinue:
		stmtType = "continue"
		node, err = p.parseCtrl()
	case ast.KeywordLet:
		stmtType = "let"
		node, err = p.parseLet()
	case ast.KeywordReturn:
		stmtType = "return"
		node, err = p.parseReturn()
	case ast.KeywordIf:
		node, err = p.parseIf()
		hasSemi = false
	case ast.KeywordSwitch:
		node, err = p.parseSwitch()
		hasSemi = false
	case ast.KeywordFn:
		node, err = p.parseFn(true)
		hasSemi = false
	case ast.KeywordWhile:
		node, err = p.parseWhile()
		hasSemi = false
	case ast.KeywordFor:
		node, err = p.parseFor()
		hasSemi = false
	case ast.ItemLeftBrace:
		node, err = p.parseBlock()
		hasSemi = false
	case ast.ItemIdentifier:
		p.next()
		// check if = after
		if x := p.peek(); x.Typ == ast.ItemAssign {
			p.backup()
			node, err = p.parseAssign()
			break
		}
		// it's an expr if no assign
		p.backup()
		fallthrough
	default:
		node, err = p.parseExpr(Lowest)
	}
	// check if any of the parse error'd
	if err != nil {
		return node, err
	}
	// check for semi
	if hasSemi {
		n, err := p.consume(ast.ItemTerminateLine)
		if err != nil {
			return nil, fmt.Errorf("ln%v: expecting ; at end of %v statement, got %v", n.Line, stmtType, n.Val)
		}
	}
	return node, nil
}

// excepting let ident = expr;
func (p *Parser) parseLet() (ast.Stmt, error) {
	n := p.next()

	ident, err := p.consume(ast.ItemIdentifier)
	if err != nil {
		// next token not an identifier
		return nil, fmt.Errorf("ln%v: expecting identifier after let, got %v", ident.Line, ident.Val)
	}

	a, err := p.consume(ast.ItemAssign)
	if err != nil {
		// next token not and identifier
		return nil, fmt.Errorf("ln%v: expecting = after identifier in let statement, got %v", a.Line, a.Val)
	}

	expr, err := p.parseExpr(Lowest)

	stmt := &ast.LetStmt{
		Pos:   n.Pos,
		Ident: ident,
		Val:   expr,
	}

	return stmt, err
}

// expecting ident = expr
func (p *Parser) parseAssign() (ast.Stmt, error) {
	ident, err := p.consume(ast.ItemIdentifier)
	if err != nil {
		// next token not and identifier
		return nil, fmt.Errorf("ln%v: expecting identifier in assign statement, got %v", ident.Line, ident.Val)
	}

	a, err := p.consume(ast.ItemAssign)
	if err != nil {
		// next token not and identifier
		return nil, fmt.Errorf("ln%v: expecting = after identifier in assign statement, got %v", a.Line, a.Val)
	}

	expr, err := p.parseExpr(Lowest)

	if err != nil {
		return nil, err
	}

	stmt := &ast.AssignStmt{
		Pos:   ident.Pos,
		Ident: ident,
		Val:   expr,
	}

	return stmt, nil
}

func (p *Parser) parseIf() (ast.Stmt, error) {
	n := p.next()

	stmt := &ast.IfStmt{
		Pos: n.Pos,
	}

	var err error

	stmt.Condition, err = p.parseExpr(Lowest)
	if err != nil {
		return nil, err
	}

	// expecting a { next
	if n := p.peek(); n.Typ != ast.ItemLeftBrace {
		return nil, fmt.Errorf("ln%v: expecting { after if, got %v", n.Line, n.Val)
	}

	stmt.IfBlock, err = p.parseBlock() // parse block here
	if err != nil {
		return nil, err
	}

	// stop if no else
	if n := p.peek(); n.Typ != ast.KeywordElse {
		return stmt, nil
	}

	// skip the else keyword
	p.next()

	// expecting another stmt (should be either if or block)
	block, err := p.parseStatement()
	switch block.(type) {
	case *ast.IfStmt, *ast.BlockStmt:
	default:
		stmt.ElseBlock = nil
		return stmt, fmt.Errorf("ln%v: expecting either if or normal block after else", n.Line)
	}

	stmt.ElseBlock = block.(ast.Stmt)

	return stmt, err
}

func (p *Parser) parseSwitch() (ast.Stmt, error) {
	// switch expr { }
	n, err := p.consume(ast.KeywordSwitch)
	if err != nil {
		panic("unreachable")
	}

	stmt := &ast.SwitchStmt{
		Pos: n.Pos,
	}

	// condition can be optional; if next item is itemLeftBrace then simply set condition to 1
	if n := p.peek(); n.Typ != ast.ItemLeftBrace {
		stmt.Condition, err = p.parseExpr(Lowest)
		if err != nil {
			return nil, err
		}
	} else {
		stmt.Condition = nil
	}

	if n := p.next(); n.Typ != ast.ItemLeftBrace {
		return nil, fmt.Errorf("ln%v: expecting { after switch, got %v", n.Line, n.Val)
	}

	// look for cases while not }
	for n := p.next(); n.Typ != ast.ItemRightBrace; n = p.next() {
		var err error
		// expecting case expr: block
		switch n.Typ {
		case ast.KeywordCase:
			cs := &ast.CaseStmt{
				Pos: n.Pos,
			}
			cs.Condition, err = p.parseExpr(Lowest)
			if err != nil {
				return nil, err
			}
			// colon, then read until we hit next case
			if n := p.peek(); n.Typ != ast.ItemColon {
				return nil, fmt.Errorf("ln%v: expecting : after case, got %v", n.Line, n.Val)
			}
			cs.Body, err = p.parseCaseBody()
			if err != nil {
				return nil, err
			}
			stmt.Cases = append(stmt.Cases, cs)
		case ast.KeywordDefault:
			// colon, then read until we hit next case
			if p.peek().Typ != ast.ItemColon {
				return nil, fmt.Errorf("ln%v: expecting : after default, got %v", n.Line, n.Val)
			}
			stmt.Default, err = p.parseCaseBody()
			if err != nil {
				return nil, err
			}
		default:
			return nil, fmt.Errorf("ln%v: expecting case or default token, got %v", n.Line, n.Val)
		}
	}

	return stmt, nil
}

func (p *Parser) parseCaseBody() (*ast.BlockStmt, error) {
	n := p.next() // start with :
	block := ast.NewBlockStmt(n.Pos)
	var node ast.Node
	var err error
	// parse line by line until we hit }
	for {
		// make sure we don't get any illegal lines
		switch n := p.peek(); n.Typ {
		case ast.KeywordDefault:
			fallthrough
		case ast.KeywordCase:
			fallthrough
		case ast.ItemRightBrace:
			return block, nil
		case ast.ItemEOF:
			return nil, fmt.Errorf("reached end of file without closing }")
		}
		// parse statement here
		node, err = p.parseStatement()
		if err != nil {
			return nil, err
		}
		block.Append(node)
	}
}

// while { }
func (p *Parser) parseWhile() (ast.Stmt, error) {
	n := p.next()

	stmt := &ast.WhileStmt{
		Pos: n.Pos,
	}

	var err error

	stmt.Condition, err = p.parseExpr(Lowest)
	if err != nil {
		return nil, err
	}

	// expecting a { next
	if n := p.peek(); n.Typ != ast.ItemLeftBrace {
		return nil, fmt.Errorf("ln%v: expecting { after while, got %v", n.Line, n.Val)
	}

	stmt.WhileBlock, err = p.parseBlock() // parse block here

	return stmt, err
}

// for <init ;> <cond> <; post> { <body> }
// for { <body > }
func (p *Parser) existVarDecl() bool {
	switch n := p.peek(); n.Typ {
	case ast.KeywordLet:
		return true
	case ast.ItemIdentifier:
		p.next()
		b := p.peek().Typ == ast.ItemAssign
		p.backup()
		return b
	}
	return false
}

func (p *Parser) parseFor() (ast.Stmt, error) {
	n := p.next()

	stmt := &ast.ForStmt{
		Pos: n.Pos,
	}

	var err error

	if n := p.peek(); n.Typ == ast.ItemLeftBrace {
		stmt.Body, err = p.parseBlock() // parse block here
		return stmt, err
	}

	// init
	if p.existVarDecl() {
		if n := p.peek(); n.Typ == ast.KeywordLet {
			stmt.Init, err = p.parseLet()
		} else {
			stmt.Init, err = p.parseAssign()
		}
		if err != nil {
			return nil, err
		}

		if n := p.peek(); n.Typ != ast.ItemTerminateLine {
			return nil, fmt.Errorf("ln%v: expecting ; after statement, got %v", n.Line, n.Val)
		}
		p.next() // skip ;
	}

	// cond
	stmt.Cond, err = p.parseExpr(Lowest)
	if err != nil {
		return nil, err
	}

	// post
	if n := p.peek(); n.Typ == ast.ItemTerminateLine {
		p.next() // skip ;
		if n := p.peek(); n.Typ != ast.ItemLeftBrace {
			stmt.Post, err = p.parseAssign()
			if err != nil {
				return nil, err
			}
		}
	}

	// expecting a { next
	if n := p.peek(); n.Typ != ast.ItemLeftBrace {
		return nil, fmt.Errorf("ln%v: expecting { after for, got %v", n.Line, n.Val)
	}

	stmt.Body, err = p.parseBlock() // parse block here

	return stmt, err
}

func (p *Parser) parseFn(ident bool) (ast.Stmt, error) {
	// fn ident(...ident){ block }
	// consume fn
	n := p.next()
	stmt := &ast.FnStmt{
		Pos: n.Pos,
	}

	var err error
	if ident {
		// ident next
		n, err := p.consume(ast.ItemIdentifier)
		if err != nil {
			return nil, fmt.Errorf("ln%v: expecting identifier after fn, got %v", n.Line, n.Val)
		}
		stmt.FunVal = n
	}

	if l := p.peek(); l.Typ != ast.ItemLeftParen {
		return nil, fmt.Errorf("ln%v: expecting ( after identifier, got %v", l.Line, l.Val)
	}

	stmt.Args, err = p.parseFnArgs()
	if err != nil {
		return nil, err
	}
	stmt.Body, err = p.parseBlock()
	if err != nil {
		return nil, err
	}

	// check that args are not duplicates
	chk := make(map[string]bool)
	for _, v := range stmt.Args {
		if _, ok := chk[v.Value]; ok {
			return nil, fmt.Errorf("fn %v contains duplicated param name %v", stmt.FunVal.Val, v.Value)
		}
		chk[v.Value] = true
	}

	return stmt, nil
}

func (p *Parser) parseFnArgs() ([]*ast.Ident, error) {
	// consume (
	var args []*ast.Ident
	p.next()
	for n := p.next(); n.Typ != ast.ItemRightParen; n = p.next() {
		a := &ast.Ident{}
		// expecting ident, comma
		if n.Typ != ast.ItemIdentifier {
			return nil, fmt.Errorf("ln%v: expecting identifier in param list, got %v", n.Line, n.Val)
		}
		a.Pos = n.Pos
		a.Value = n.Val

		args = append(args, a)

		// if next token is a comma, then there should be another ident after that
		// otherwise we have a problem
		if l := p.peek(); l.Typ == ast.ItemComma {
			p.next() // consume the comma
			if l = p.peek(); l.Typ != ast.ItemIdentifier {
				return nil, fmt.Errorf("ln%v: expecting another identifier after comma in param list, got %v", n.Line, n.Val)
			}
		}
	}
	return args, nil
}

func (p *Parser) parseReturn() (ast.Stmt, error) {
	n := p.next() // return
	stmt := &ast.ReturnStmt{
		Pos: n.Pos,
	}
	var err error
	stmt.Val, err = p.parseExpr(Lowest)
	return stmt, err
}

func (p *Parser) parseCtrl() (ast.Stmt, error) {
	n := p.next()
	stmt := &ast.CtrlStmt{
		Pos: n.Pos,
	}
	switch n.Typ {
	case ast.KeywordBreak:
		stmt.Typ = ast.CtrlBreak
	case ast.KeywordContinue:
		stmt.Typ = ast.CtrlContinue
	case ast.KeywordFallthrough:
		stmt.Typ = ast.CtrlFallthrough
	default:
		return nil, fmt.Errorf("ln%v: expecting ctrl token, got %v", n.Line, n.Val)
	}
	return stmt, nil
}

func (p *Parser) parseCall(fun ast.Expr) (ast.Expr, error) {
	// expecting (params)
	n, err := p.consume(ast.ItemLeftParen)
	if err != nil {
		return nil, fmt.Errorf("expecting ( after ident, got %v", fun.String())
	}
	expr := &ast.CallExpr{
		Pos: n.Pos,
		Fun: fun,
	}
	expr.Args, err = p.parseCallArgs()

	return expr, err
}

func (p *Parser) parseCallArgs() ([]ast.Expr, error) {
	var args []ast.Expr

	if p.peek().Typ == ast.ItemRightParen {
		// consume the right paren
		p.next()
		return args, nil
	}

	// next should be an expression
	exp, err := p.parseExpr(Lowest)
	if err != nil {
		return args, err
	}
	args = append(args, exp)

	for p.peek().Typ == ast.ItemComma {
		p.next() // skip the comma
		exp, err = p.parseExpr(Lowest)
		if err != nil {
			return args, err
		}
		args = append(args, exp)
	}

	if n := p.next(); n.Typ != ast.ItemRightParen {
		p.backup()
		return nil, fmt.Errorf("ln%v: expecting ) at end of function call, got: %v", n.Line, n.Val)
	}

	return args, nil
}

// parseBlock return a node contain and BlockStmt
func (p *Parser) parseBlock() (*ast.BlockStmt, error) {
	// should be surronded by {}
	n, err := p.consume(ast.ItemLeftBrace)
	if err != nil {
		return nil, fmt.Errorf("ln%v: expecting {, got %v", n.Line, n.Val)
	}
	block := ast.NewBlockStmt(n.Pos)
	var node ast.Node
	// parse line by line until we hit }
	for {
		// make sure we don't get any illegal lines
		switch n := p.peek(); n.Typ {
		case ast.ItemRightBrace:
			p.next() // consume the braces
			return block, nil
		case ast.ItemEOF:
			return nil, fmt.Errorf("reached end of file without closing }")
		}
		// parse statement here
		node, err = p.parseStatement()
		if err != nil {
			return nil, err
		}
		block.Append(node)
	}
}

func (p *Parser) parseExpr(pre precedence) (ast.Expr, error) {
	t := p.next()
	prefix := p.prefixParseFns[t.Typ]
	if prefix == nil {
		return nil, nil
	}
	p.backup()
	leftExp, err := prefix()
	if err != nil {
		return nil, err
	}

	for n := p.peek(); n.Typ != ast.ItemTerminateLine && pre < tokenPrecendence(n); n = p.peek() {
		infix := p.infixParseFns[n.Typ]
		if infix == nil {
			return leftExp, nil
		}

		leftExp, err = infix(leftExp)
		if err != nil {
			return nil, err
		}
	}

	return leftExp, nil
}

// next is an identifier
func (p *Parser) parseIdent() (ast.Expr, error) {
	n := p.next()
	return &ast.Ident{Pos: n.Pos, Value: n.Val}, nil
}

func (p *Parser) parseString() (ast.Expr, error) {
	n := p.next()
	return &ast.StringLit{Pos: n.Pos, Value: n.Val}, nil
}

func (p *Parser) parseNull() (ast.Expr, error) {
	n := p.next()
	return &ast.NullLit{Pos: n.Pos}, nil
}

func (p *Parser) parseFnLit() (ast.Expr, error) {
	n := p.peek()
	stmt, err := p.parseFn(false)
	if err != nil {
		return nil, err
	}

	f := stmt.(*ast.FnStmt)
	return &ast.FuncLit{
		Pos:  n.Pos,
		Args: f.Args,
		Body: f.Body,
	}, nil
}

func (p *Parser) parseNumber() (ast.Expr, error) {
	// string, int, float, or bool
	n := p.next()
	num := &ast.NumberLit{Pos: n.Pos}
	// try parse int, if not ok then try parse float
	iv, err := strconv.ParseInt(n.Val, 10, 64)
	if err == nil {
		num.IntVal = iv
		num.FloatVal = float64(iv)
	} else {
		fv, err := strconv.ParseFloat(n.Val, 64)
		if err != nil {
			return nil, fmt.Errorf("ln%v: cannot parse %v to number", n.Line, n.Val)
		}
		num.IsFloat = true
		num.FloatVal = fv
	}
	return num, nil
}

func (p *Parser) parseBool() (ast.Expr, error) {
	// bool is a number (true = 1, false = 0)
	n := p.next()
	num := &ast.NumberLit{Pos: n.Pos}
	switch n.Val {
	case "true":
		num.IntVal = 1
		num.FloatVal = 1
	case "false":
		num.IntVal = 0
		num.FloatVal = 0
	default:
		return nil, fmt.Errorf("ln%v: expecting boolean, got %v", n.Line, n.Val)
	}
	return num, nil
}

func (p *Parser) parseUnaryExpr() (ast.Expr, error) {
	n := p.next()
	switch n.Typ {
	case ast.LogicNot:
	case ast.ItemMinus:
	default:
		return nil, fmt.Errorf("ln%v: unrecognized unary operator %v", n.Line, n.Val)
	}
	var err error
	expr := &ast.UnaryExpr{
		Pos: n.Pos,
		Op:  n,
	}
	expr.Right, err = p.parseExpr(Prefix)
	return expr, err
}

func (p *Parser) parseBinaryExpr(left ast.Expr) (ast.Expr, error) {
	n := p.next()
	expr := &ast.BinaryExpr{
		Pos:  n.Pos,
		Op:   n,
		Left: left,
	}
	pr := tokenPrecendence(n)
	var err error
	expr.Right, err = p.parseExpr(pr)
	return expr, err
}

func (p *Parser) parseParen() (ast.Expr, error) {
	// skip the paren
	p.next()

	exp, err := p.parseExpr(Lowest)
	if err != nil {
		return nil, err
	}

	if n := p.peek(); n.Typ != ast.ItemRightParen {
		return nil, nil
	}
	p.next() // consume the right paren

	return exp, nil
}

func (p *Parser) parseMap() (ast.Expr, error) {
	// skip the paren
	n := p.next()
	expr := &ast.MapExpr{
		Pos:    n.Pos,
		Fields: make(map[string]ast.Expr),
	}

	if p.peek().Typ == ast.ItemRightSquareParen { // empty map
		p.next()
		return expr, nil
	}

	// loop until we hit square paren
	for {
		// we're expecting ident = int
		i, err := p.consume(ast.ItemIdentifier)
		if err != nil {
			return nil, fmt.Errorf("ln%v: expecting identifier in map expression, got %v", i.Line, i.Val)
		}

		a, err := p.consume(ast.ItemAssign)
		if err != nil {
			return nil, fmt.Errorf("ln%v: expecting = after identifier in map expression, got %v", a.Line, a.Val)
		}

		e, err := p.parseExpr(Lowest)
		if err != nil {
			return nil, err
		}
		expr.Fields[i.Val] = e

		// if we hit ], return; if we hit , keep going, other wise error
		n := p.next()
		switch n.Typ {
		case ast.ItemRightSquareParen:
			return expr, nil
		case ast.ItemComma:
			// do nothing, keep going
		default:
			return nil, fmt.Errorf("ln%v: <action param> bad token %v", n.Line, n)
		}
	}
}
