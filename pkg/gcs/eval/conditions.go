package eval

import (
	"fmt"

	"github.com/simimpact/srsim/pkg/gcs/ast"
	"github.com/simimpact/srsim/pkg/key"
)

// Functions for writing more flexible scripts.
func (e *Eval) initConditionalFuncs(env *Env) {
	e.addFunction("skill_ready", e.skillReady, env)
	e.addFunction("ult_ready", e.ultReady, env)
	e.addFunction("skill_points", e.skillPoints, env)
	// TODO: whos_next()?
	e.addFunction("has_modifier", e.hasModifier, env)
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
		return nil, fmt.Errorf("skill_ready argument char should evaluate to a number, got %v", tarobj.Inspect())
	}
	target := key.TargetID(tarobj.(*number).ival)

	result, err := e.Engine.CanUseSkill(target)
	if err != nil {
		return nil, err
	}
	return bton(result), nil
}

func (e *Eval) ultReady(c *ast.CallExpr, env *Env) (Obj, error) {
	// ult_ready(char)
	if len(c.Args) != 1 {
		return nil, fmt.Errorf("invalid number of params for ult_ready, expected 1 got %v", len(c.Args))
	}

	// should eval to a number
	tarobj, err := e.evalExpr(c.Args[0], env)
	if err != nil {
		return nil, err
	}
	if tarobj.Typ() != typNum {
		return nil, fmt.Errorf("ult_ready argument char should evaluate to a number, got %v", tarobj.Inspect())
	}
	target := key.TargetID(tarobj.(*number).ival)

	if !e.Engine.IsCharacter(target) {
		return nil, fmt.Errorf("target %d is not a character", target)
	}
	return bton(e.Engine.EnergyRatio(target) >= 1), nil
}

func (e *Eval) skillPoints(c *ast.CallExpr, env *Env) (Obj, error) {
	// skill_points()
	if len(c.Args) != 0 {
		return nil, fmt.Errorf("invalid number of params for skill_points, expected 0 got %v", len(c.Args))
	}

	return &number{ival: int64(e.Engine.SP())}, nil
}

func (e *Eval) hasModifier(c *ast.CallExpr, env *Env) (Obj, error) {
	// has_modifier(char, mod)
	if len(c.Args) != 1 {
		return nil, fmt.Errorf("invalid number of params for has_modifier, expected 2 got %v", len(c.Args))
	}

	// should eval to a number
	tarobj, err := e.evalExpr(c.Args[0], env)
	if err != nil {
		return nil, err
	}
	if tarobj.Typ() != typNum {
		return nil, fmt.Errorf("has_modifier argument char should evaluate to a number, got %v", tarobj.Inspect())
	}
	target := key.TargetID(tarobj.(*number).ival)

	// should eval to a string
	modobj, err := e.evalExpr(c.Args[1], env)
	if err != nil {
		return nil, err
	}
	if modobj.Typ() != typStr {
		return nil, fmt.Errorf("has_modifier argument mod should evaluate to a string, got %v", tarobj.Inspect())
	}
	modifier := key.Modifier(tarobj.(*strval).str)

	if !e.Engine.IsValid(target) {
		return nil, fmt.Errorf("target %d is invalid", target)
	}
	return bton(e.Engine.HasModifier(target, modifier)), nil
}
