package eval

import (
	"fmt"

	"github.com/simimpact/srsim/pkg/gcs/ast"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

// Functions for writing more flexible scripts.
func (e *Eval) initConditionalFuncs(env *Env) {
	// modifier
	e.addFunction("has_modifier", e.hasModifier, env)
	e.addFunction("modifier_count", e.modifierCount, env)

	// attribute
	e.addFunction("ult_ready", e.ultReady, env)
	e.addFunction("skill_points", e.skillPoints, env)
	e.addFunction("energy", e.energy, env)
	e.addFunction("max_energy", e.maxEnergy, env)
	e.addFunction("hp_ratio", e.hpRatio, env)

	// turn
	// TODO: whos_next()?

	// info
	e.addFunction("skill_ready", e.skillReady, env)

	// StatusType
	e.addConstant("StatusBuff", &number{ival: int64(model.StatusType_STATUS_BUFF)}, env)
	e.addConstant("StatusDebuff", &number{ival: int64(model.StatusType_STATUS_DEBUFF)}, env)
}

// modifier

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
		return nil, fmt.Errorf("has_modifier argument mod should evaluate to a string, got %v", modobj.Inspect())
	}
	modifier := key.Modifier(tarobj.(*strval).str)

	if !e.Engine.IsValid(target) {
		return nil, fmt.Errorf("target %d is invalid", target)
	}
	return bton(e.Engine.HasModifier(target, modifier)), nil
}

func (e *Eval) modifierCount(c *ast.CallExpr, env *Env) (Obj, error) {
	// modifier_count(char, type)
	if len(c.Args) != 1 {
		return nil, fmt.Errorf("invalid number of params for modifier_count, expected 2 got %v", len(c.Args))
	}

	// should eval to a number
	tarobj, err := e.evalExpr(c.Args[0], env)
	if err != nil {
		return nil, err
	}
	if tarobj.Typ() != typNum {
		return nil, fmt.Errorf("modifier_count argument char should evaluate to a number, got %v", tarobj.Inspect())
	}
	target := key.TargetID(tarobj.(*number).ival)

	// should eval to a number
	typobj, err := e.evalExpr(c.Args[1], env)
	if err != nil {
		return nil, err
	}
	if typobj.Typ() != typNum {
		return nil, fmt.Errorf("modifier_count argument type should evaluate to a number, got %v", typobj.Inspect())
	}
	status := model.StatusType(typobj.(*number).ival)

	if !e.Engine.IsValid(target) {
		return nil, fmt.Errorf("target %d is invalid", target)
	}
	return &number{ival: int64(e.Engine.ModifierCount(target, status))}, nil
}

// attribute

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

func (e *Eval) energy(c *ast.CallExpr, env *Env) (Obj, error) {
	// energy(char)
	if len(c.Args) != 1 {
		return nil, fmt.Errorf("invalid number of params for energy, expected 1 got %v", len(c.Args))
	}

	// should eval to a number
	tarobj, err := e.evalExpr(c.Args[0], env)
	if err != nil {
		return nil, err
	}
	if tarobj.Typ() != typNum {
		return nil, fmt.Errorf("energy argument char should evaluate to a number, got %v", tarobj.Inspect())
	}
	target := key.TargetID(tarobj.(*number).ival)

	if !e.Engine.IsValid(target) {
		return nil, fmt.Errorf("target %d is invalid", target)
	}
	return &number{ival: int64(e.Engine.Energy(target))}, nil
}

func (e *Eval) maxEnergy(c *ast.CallExpr, env *Env) (Obj, error) {
	// max_energy(char)
	if len(c.Args) != 1 {
		return nil, fmt.Errorf("invalid number of params for max_energy, expected 1 got %v", len(c.Args))
	}

	// should eval to a number
	tarobj, err := e.evalExpr(c.Args[0], env)
	if err != nil {
		return nil, err
	}
	if tarobj.Typ() != typNum {
		return nil, fmt.Errorf("max_energy argument char should evaluate to a number, got %v", tarobj.Inspect())
	}
	target := key.TargetID(tarobj.(*number).ival)

	if !e.Engine.IsValid(target) {
		return nil, fmt.Errorf("target %d is invalid", target)
	}
	return &number{ival: int64(e.Engine.MaxEnergy(target))}, nil
}

func (e *Eval) hpRatio(c *ast.CallExpr, env *Env) (Obj, error) {
	// hp_ratio(char)
	if len(c.Args) != 1 {
		return nil, fmt.Errorf("invalid number of params for hp_ratio, expected 1 got %v", len(c.Args))
	}

	// should eval to a number
	tarobj, err := e.evalExpr(c.Args[0], env)
	if err != nil {
		return nil, err
	}
	if tarobj.Typ() != typNum {
		return nil, fmt.Errorf("hp_ratio argument char should evaluate to a number, got %v", tarobj.Inspect())
	}
	target := key.TargetID(tarobj.(*number).ival)

	if !e.Engine.IsValid(target) {
		return nil, fmt.Errorf("target %d is invalid", target)
	}
	return &number{
		fval:    e.Engine.HPRatio(target),
		isFloat: true,
	}, nil
}

// info

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
