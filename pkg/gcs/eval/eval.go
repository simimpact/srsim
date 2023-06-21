package eval

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/action"
	"github.com/simimpact/srsim/pkg/engine/target/evaltarget"
	"github.com/simimpact/srsim/pkg/gcs/ast"
	"github.com/simimpact/srsim/pkg/key"
)

type TargetNode struct {
	target key.TargetID
	env    *Env
	node   ast.Node
}

type Eval struct {
	AST    ast.Node
	global *Env
	ctx    context.Context
	Engine engine.Engine

	targetNode     map[key.TargetID]TargetNode
	ultNodes       []TargetNode
	defaultActions map[key.TargetID]*action.Action
}

type Env struct {
	parent *Env
	varMap map[string]*Obj
}

func NewEnv(parent *Env) *Env {
	return &Env{
		parent: parent,
		varMap: make(map[string]*Obj),
	}
}

func (e *Env) v(s string) (*Obj, error) {
	v, ok := e.varMap[s]
	if ok {
		return v, nil
	}
	if e.parent != nil {
		return e.parent.v(s)
	}
	return nil, fmt.Errorf("variable %v does not exist", s)
}

func New(ast *ast.BlockStmt, ctx context.Context) *Eval {
	e := &Eval{AST: ast}
	e.Init(ctx)
	return e
}

// Run will execute the provided AST.
func (e *Eval) Init(ctx context.Context) error {
	e.ctx = ctx
	e.global = NewEnv(nil)
	e.targetNode = make(map[key.TargetID]TargetNode)
	e.ultNodes = make([]TargetNode, 0)
	e.defaultActions = make(map[key.TargetID]*action.Action)
	e.initSysFuncs(e.global)

	_, err := e.evalNode(e.AST, e.global)
	if err != nil {
		return err
	}
	return nil
}

var ErrTerminated = errors.New("eval terminated")

type Obj interface {
	Inspect() string
	Typ() ObjTyp
}

type ObjTyp int

const (
	typNull ObjTyp = iota
	typNum
	typStr
	typFun
	typBif // built-in function
	typAct
	typMap
	typRet
	typCtr
	// typTerminate
)

// various Obj types
type (
	null   struct{}
	number struct {
		ival    int64
		fval    float64
		isFloat bool
	}

	strval struct {
		str string
	}

	funcval struct {
		Args []*ast.Ident
		Body *ast.BlockStmt
	}

	bfuncval struct {
		Body func(c *ast.CallExpr, env *Env) (Obj, error)
	}

	actionval struct {
		val action.Action
	}

	mapval struct {
		fields map[string]Obj
	}

	retval struct {
		res Obj
	}

	ctrl struct {
		typ ast.CtrlTyp
	}
)

// null.
func (n *null) Inspect() string { return "null" }
func (n *null) Typ() ObjTyp     { return typNull }

// terminate.
// func (n *terminate) Inspect() string { return "terminate" }
// func (n *terminate) Typ() ObjTyp     { return typTerminate }

// number.
func (n *number) Inspect() string {
	if n.isFloat {
		return strconv.FormatFloat(n.fval, 'f', -1, 64)
	} else {
		return strconv.FormatInt(n.ival, 10)
	}
}
func (n *number) Typ() ObjTyp { return typNum }

// strval.
func (s *strval) Inspect() string { return s.str }
func (s *strval) Typ() ObjTyp     { return typStr }

// funcval.
func (f *funcval) Inspect() string { return "function" }
func (f *funcval) Typ() ObjTyp     { return typFun }

// bfuncval.
func (b *bfuncval) Inspect() string { return "built-in function" }
func (b *bfuncval) Typ() ObjTyp     { return typBif }

// retval.
func (r *retval) Inspect() string {
	return r.res.Inspect()
}
func (n *retval) Typ() ObjTyp { return typRet }

// actionval.
func (a *actionval) Inspect() string {
	targeteval := ""
	switch a.val.TargetEvaluator {
	case evaltarget.First:
		targeteval = "First"
	case evaltarget.LowestHP:
		targeteval = "LowestHP"
	case evaltarget.LowestHPRatio:
		targeteval = "LowestHPRatio"
	default:
		targeteval = strconv.Itoa(int(a.val.TargetEvaluator))
	}
	return string(a.val.Type) + "(" + targeteval + ")"
}
func (a *actionval) Typ() ObjTyp { return typAct }

// mapval.
func (m *mapval) Inspect() string {
	str := "["
	done := false
	for k, v := range m.fields {
		if done {
			str += ", "
		}
		done = true

		str += k + " = " + v.Inspect()
	}
	str += "]"
	return str
}
func (m *mapval) Typ() ObjTyp { return typMap }

// ctrl.
func (c *ctrl) Inspect() string {
	switch c.typ {
	case ast.CtrlContinue:
		return "continue"
	case ast.CtrlBreak:
		return "break"
	case ast.CtrlFallthrough:
		return "fallthrough"
	}
	return "invalid"
}
func (c *ctrl) Typ() ObjTyp { return typCtr }

func (e *Eval) evalNode(n ast.Node, env *Env) (Obj, error) {
	switch v := n.(type) {
	case ast.Expr:
		return e.evalExpr(v, env)
	case ast.Stmt:
		return e.evalStmt(v, env)
	default:
		return &null{}, nil
	}
}
