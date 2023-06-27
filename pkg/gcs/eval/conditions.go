package eval

import (
	"fmt"

	"github.com/simimpact/srsim/pkg/gcs/ast"
	"github.com/simimpact/srsim/pkg/key"
)

func (e *Eval) initConditionalFuncs(env *Env) {
	e.addFunction("skill_ready", e.skillReady, env)
}

func (e *Eval) skillReady(c *ast.CallExpr, env *Env) (Obj, error) {
	// skill_ready(char)
	if len(c.Args) != 1 {
		return nil, fmt.Errorf("invalid number of params for skill_ready, expected 1 got %v", len(c.Args))
	}

	// should eval to a number
	tarobj, err := e.evalExpr(c.Args[0], env)
	if err != nil {
		return nil, err
	}
	if tarobj.Typ() != typNum {
		return nil, fmt.Errorf("register_skill_cb argument char should evaluate to a number, got %v", tarobj.Inspect())
	}
	target := key.TargetID(tarobj.(*number).ival)

	result, err := e.Engine.CanUseSkill(target)
	if err != nil {
		return nil, err
	}
	return bton(result), nil
}
