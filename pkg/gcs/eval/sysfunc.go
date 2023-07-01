package eval

import (
	"fmt"
	"strings"

	"github.com/simimpact/srsim/pkg/engine/action"
	"github.com/simimpact/srsim/pkg/engine/target/evaltarget"
	"github.com/simimpact/srsim/pkg/gcs/ast"
	"github.com/simimpact/srsim/pkg/key"
)

func (e *Eval) initSysFuncs(env *Env) {
	// std funcs
	e.addFunction("rand", e.rand, env)
	e.addFunction("randnorm", e.randnorm, env)
	e.addFunction("print", e.print, env)
	e.addFunction("type", e.typeval, env)
	e.addFunction("register_skill_cb", e.registerSkillCB, env)
	e.addFunction("register_ult_cb", e.registerUltCB, env)
	e.addFunction("set_default_action", e.setDefaultAction, env)

	// actions
	e.addAction(key.ActionAttack, env)
	e.addAction(key.ActionSkill, env)
	e.addAction(key.ActionUlt, env)
	e.addAction(key.ActionUltAttack, env)
	e.addAction(key.ActionUltSkill, env)

	// target evaluators
	e.addConstant("First", &number{ival: int64(evaltarget.First)}, env)
	e.addConstant("LowestHP", &number{ival: int64(evaltarget.LowestHP)}, env)
	e.addConstant("LowestHPRatio", &number{ival: int64(evaltarget.LowestHPRatio)}, env)

	// chars
	if e.Engine != nil {
		for _, k := range e.Engine.Characters() {
			char, err := e.Engine.CharacterInfo(k)
			if err != nil { // ???
				return
			}
			e.addConstant(string(char.Key), &number{ival: int64(k)}, env)
		}
	}
}

func (e *Eval) print(c *ast.CallExpr, env *Env) (Obj, error) {
	// concat all args
	var sb strings.Builder
	for _, arg := range c.Args {
		val, err := e.evalExpr(arg, env)
		if err != nil {
			return nil, err
		}
		sb.WriteString(val.Inspect())
	}
	fmt.Println(sb.String())
	return &null{}, nil
}

func (e *Eval) rand(c *ast.CallExpr, env *Env) (Obj, error) {
	x := e.Engine.Rand().Float64()
	return &number{
		fval:    x,
		isFloat: true,
	}, nil
}

func (e *Eval) randnorm(c *ast.CallExpr, env *Env) (Obj, error) {
	x := e.Engine.Rand().NormFloat64()
	return &number{
		fval:    x,
		isFloat: true,
	}, nil
}

func (e *Eval) typeval(c *ast.CallExpr, env *Env) (Obj, error) {
	// type(var)
	if len(c.Args) != 1 {
		return nil, fmt.Errorf("invalid number of params for type, expected 1 got %v", len(c.Args))
	}

	t, err := e.evalExpr(c.Args[0], env)
	if err != nil {
		return nil, err
	}

	str := "unknown"
	switch t.Typ() {
	case typNull:
		str = "null"
	case typNum:
		str = "number"
	case typStr:
		str = "string"
	case typMap:
		str = "map"
	case typAct:
		str = "action"
	case typFun:
		fallthrough
	case typBif:
		str = t.Inspect()
	}

	return &strval{str}, nil
}

func (e *Eval) registerSkillCB(c *ast.CallExpr, env *Env) (Obj, error) {
	// register_skill_cb(char, func)
	if len(c.Args) != 2 {
		return nil, fmt.Errorf("invalid number of params for register_skill_cb, expected 2 got %v", len(c.Args))
	}

	// should eval to a number
	tarobj, err := e.evalExpr(c.Args[0], env)
	if err != nil {
		return nil, err
	}
	if tarobj.Typ() != typNum {
		return nil, fmt.Errorf("register_skill_cb argument char should evaluate to a number, got %v", tarobj.Inspect())
	}
	target := tarobj.(*number).ival

	// should eval to a function
	funcobj, err := e.evalExpr(c.Args[1], env)
	if err != nil {
		return nil, err
	}
	if funcobj.Typ() != typFun {
		return nil, fmt.Errorf("register_skill_cb argument func should evaluate to a function, got %v", funcobj.Inspect())
	}
	fn := funcobj.(*funcval)

	node := TargetNode{
		target: key.TargetID(target),
		env:    NewEnv(env),
		node:   fn.Body,
	}
	for i, v := range fn.Args {
		param, err := e.evalExpr(c.Args[i], env)
		if err != nil {
			return nil, err
		}
		node.env.varMap[v.Value] = &param
	}
	e.targetNode[key.TargetID(target)] = node
	return &null{}, nil
}

func (e *Eval) registerUltCB(c *ast.CallExpr, env *Env) (Obj, error) {
	// register_ult_cb(char, func)
	if len(c.Args) != 2 {
		return nil, fmt.Errorf("invalid number of params for register_ult_cb, expected 2 got %v", len(c.Args))
	}

	// should eval to a function
	tarobj, err := e.evalExpr(c.Args[0], env)
	if err != nil {
		return nil, err
	}
	if tarobj.Typ() != typNum {
		return nil, fmt.Errorf("register_ult_cb argument char should evaluate to a number, got %v", tarobj.Inspect())
	}
	target := tarobj.(*number).ival

	// should eval to a function
	funcobj, err := e.evalExpr(c.Args[1], env)
	if err != nil {
		return nil, err
	}
	if funcobj.Typ() != typFun {
		return nil, fmt.Errorf("register_ult_cb argument func should evaluate to a function, got %v", funcobj.Inspect())
	}
	fn := funcobj.(*funcval)

	node := TargetNode{
		target: key.TargetID(target),
		env:    NewEnv(env),
		node:   fn.Body,
	}
	for i, v := range fn.Args {
		param, err := e.evalExpr(c.Args[i], env)
		if err != nil {
			return nil, err
		}
		node.env.varMap[v.Value] = &param
	}
	e.ultNodes = append(e.ultNodes, node)
	return &null{}, nil
}

func (e *Eval) setDefaultAction(c *ast.CallExpr, env *Env) (Obj, error) {
	// set_default_action(char, action)
	if len(c.Args) != 2 {
		return nil, fmt.Errorf("invalid number of params for set_default_action, expected 2 got %v", len(c.Args))
	}

	// should eval to a function
	tarobj, err := e.evalExpr(c.Args[0], env)
	if err != nil {
		return nil, err
	}
	if tarobj.Typ() != typNum {
		return nil, fmt.Errorf("set_default_action argument char should evaluate to a number, got %v", tarobj.Inspect())
	}
	target := tarobj.(*number).ival

	// should eval to an action
	actobj, err := e.evalExpr(c.Args[1], env)
	if err != nil {
		return nil, err
	}
	if actobj.Typ() != typAct {
		return nil, fmt.Errorf("set_default_action argument func should evaluate to an action, got %v", actobj.Inspect())
	}

	act := *actobj.(*actionval)
	if act.val.Type != key.ActionAttack {
		return nil, fmt.Errorf("action should be an attack, got %v", actobj.Inspect())
	}
	act.val.Target = key.TargetID(target)
	e.defaultActions[act.val.Target] = act.val
	return &null{}, nil
}

func (e *Eval) addAction(at key.ActionType, env *Env) {
	f := func(c *ast.CallExpr, env *Env) (Obj, error) {
		// attack/skill/ult(evaltarget)
		if len(c.Args) != 1 {
			return nil, fmt.Errorf("invalid number of params for action, expected 1 got %v", len(c.Args))
		}

		etval, err := e.evalExpr(c.Args[0], env)
		if err != nil {
			return nil, err
		}
		if etval.Typ() != typNum {
			return nil, fmt.Errorf("action argument char should evaluate to a number, got %v", etval.Inspect())
		}
		evaltarget := etval.(*number).ival

		return &actionval{
			val: action.Action{
				Type:            at,
				TargetEvaluator: key.TargetEvaluator(evaltarget), // TODO: check is valid
			},
		}, nil
	}

	e.addFunction(string(at), f, env)
}
