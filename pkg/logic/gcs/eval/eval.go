package eval

import (
	"context"
	"fmt"

	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/logic"
	"github.com/simimpact/srsim/pkg/logic/gcs/ast"
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
	engine engine.Engine

	targetNode     map[key.TargetID]TargetNode
	ultNodes       []TargetNode
	defaultActions map[key.TargetID]logic.Action
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

//nolint:gocritic // *Obj is a ptrToRefParam, should be refactored to use Obj instead
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

func (e *Env) setBuiltinFunc(name string, f func(c *ast.CallExpr, env *Env) (Obj, error)) {
	var obj Obj = &bfuncval{Body: f}
	e.setv(name, obj)
}

func (e *Env) setv(name string, value Obj) {
	e.varMap[name] = &value
}

func New(ctx context.Context, ast *ast.BlockStmt) *Eval {
	e := &Eval{AST: ast}
	e.ctx = ctx
	return e
}

// Run will execute the provided AST.
func (e *Eval) Init(eng engine.Engine) error {
	e.engine = eng
	e.global = NewEnv(nil)
	e.targetNode = make(map[key.TargetID]TargetNode)
	e.ultNodes = make([]TargetNode, 0)
	e.defaultActions = make(map[key.TargetID]logic.Action)

	// standart functions
	e.initSysFuncs(e.global)
	e.initActions(e.global)

	// conditionals for hsr
	if e.engine != nil {
		e.initConditionalFuncs(e.global)
		e.initEnums(e.global)
		e.initCharNames(e.global)
	}

	_, err := e.evalNode(e.AST, e.global)
	if err != nil {
		return err
	}
	return nil
}

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
