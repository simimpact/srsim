package eval

import (
	"fmt"
	"sort"
	"strings"

	"github.com/simimpact/srsim/pkg/engine/target/evaltarget"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/logic"
	"github.com/simimpact/srsim/pkg/logic/gcs/ast"
)

func (e *Eval) initSysFuncs(env *Env) {
	funcs := map[string]func(*ast.CallExpr, *Env) (Obj, error){
		// basic
		"print": e.print,
		"type":  e.typeval,
		// math
		"rand":     e.rand,
		"randnorm": e.randnorm,
		// map
		"sort":  e.sort,
		"first": e.first,
		"any":   e.any,
		// action
		"register_skill_cb":  e.registerSkillCB,
		"register_ult_cb":    e.registerUltCB,
		"set_default_action": e.setDefaultAction,
	}
	for name, function := range funcs {
		env.setBuiltinFunc(name, function)
	}
}

func (e *Eval) initActions(env *Env) {
	actions := []logic.ActionType{
		logic.ActionAttack,
		logic.ActionSkill,
		logic.ActionUlt,
		logic.ActionUltAttack,
		logic.ActionUltSkill,
	}
	for _, action := range actions {
		e.addAction(action, env)
	}

	targetEvaluators := map[string]Obj{
		"First":         &number{ival: int64(evaltarget.First)},
		"LowestHP":      &number{ival: int64(evaltarget.LowestHP)},
		"LowestHPRatio": &number{ival: int64(evaltarget.LowestHPRatio)},
	}
	for name, value := range targetEvaluators {
		env.setv(name, value)
	}
}

// print(...)
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

// rand()
func (e *Eval) rand(c *ast.CallExpr, env *Env) (Obj, error) {
	if _, err := e.validateArguments(c.Args, env); err != nil {
		return nil, err
	}

	x := e.engine.Rand().Float64()
	return &number{
		fval:    x,
		isFloat: true,
	}, nil
}

// randnorm()
func (e *Eval) randnorm(c *ast.CallExpr, env *Env) (Obj, error) {
	if _, err := e.validateArguments(c.Args, env); err != nil {
		return nil, err
	}

	x := e.engine.Rand().NormFloat64()
	return &number{
		fval:    x,
		isFloat: true,
	}, nil
}

// type(var)
func (e *Eval) typeval(c *ast.CallExpr, env *Env) (Obj, error) {
	if len(c.Args) != 1 {
		return nil, fmt.Errorf("invalid number of params for type, expected 1 got %v", len(c.Args))
	}

	t, err := e.evalExpr(c.Args[0], env)
	if err != nil {
		return nil, err
	}
	return &strval{t.Typ().String()}, nil
}

// register_skill_cb(char, func)
func (e *Eval) registerSkillCB(c *ast.CallExpr, env *Env) (Obj, error) {
	objs, err := e.validateArguments(c.Args, env, typNum, typFun)
	if err != nil {
		return nil, err
	}
	target := objs[0].(*number).ival
	fn := objs[1].(*funcval)

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

// register_ult_cb(char, func)
func (e *Eval) registerUltCB(c *ast.CallExpr, env *Env) (Obj, error) {
	objs, err := e.validateArguments(c.Args, env, typNum, typFun)
	if err != nil {
		return nil, err
	}
	target := objs[0].(*number).ival
	fn := objs[1].(*funcval)

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

// set_default_action(char, action)
func (e *Eval) setDefaultAction(c *ast.CallExpr, env *Env) (Obj, error) {
	objs, err := e.validateArguments(c.Args, env, typNum, typAct)
	if err != nil {
		return nil, err
	}
	target := objs[0].(*number).ival
	act := objs[1].(*actionval)
	if act.val.Type != logic.ActionAttack {
		return nil, fmt.Errorf("action should be an attack, got %v", act.val.Type)
	}

	act.val.Target = key.TargetID(target)
	e.defaultActions[act.val.Target] = act.val
	return &null{}, nil
}

// attack/skill/ult(evaltarget)
func (e *Eval) addAction(at logic.ActionType, env *Env) {
	f := func(c *ast.CallExpr, env *Env) (Obj, error) {
		objs, err := e.validateArguments(c.Args, env, typNum)
		if err != nil {
			return nil, err
		}
		evaltarget := objs[0].(*number).ival

		return &actionval{
			val: logic.Action{
				Type:            at,
				TargetEvaluator: key.TargetEvaluator(evaltarget), // TODO: check is valid
			},
		}, nil
	}

	env.setBuiltinFunc(string(at), f)
}

// sort(map, callback)
func (e *Eval) sort(c *ast.CallExpr, env *Env) (Obj, error) {
	objs, err := e.validateArguments(c.Args, env, typMap, typFun)
	if err != nil {
		return nil, err
	}
	m := objs[0].(*mapval)
	callback := objs[1].(*funcval)

	if len(callback.Args) != 2 {
		return nil, fmt.Errorf("invalid number of params for callback, expected 2 got %v", len(callback.Args))
	}

	local := NewEnv(env)
	sort.SliceStable(m.array, func(i, j int) bool {
		// cancel sorting if an error occurs
		if err != nil {
			return false
		}

		var result Obj
		local.varMap[callback.Args[0].Value] = &m.array[i]
		local.varMap[callback.Args[1].Value] = &m.array[j]
		result, err = e.evalNode(callback.Body, local)
		if err != nil {
			return false
		}

		if result.Typ() != typRet {
			err = fmt.Errorf("the callback must return a value")
			return false
		}
		return otob(result.(*retval).res)
	})
	if err != nil {
		return nil, err
	}
	return m, nil
}

// first(map)
func (e *Eval) first(c *ast.CallExpr, env *Env) (Obj, error) {
	objs, err := e.validateArguments(c.Args, env, typMap)
	if err != nil {
		return nil, err
	}
	m := objs[0].(*mapval)
	if len(m.array) == 0 {
		return &null{}, nil
	}
	return m.array[0], nil
}

// any(map, callback)
func (e *Eval) any(c *ast.CallExpr, env *Env) (Obj, error) {
	objs, err := e.validateArguments(c.Args, env, typMap, typFun)
	if err != nil {
		return nil, err
	}
	m := objs[0].(*mapval)
	callback := objs[1].(*funcval)

	if len(callback.Args) != 1 {
		return nil, fmt.Errorf("invalid number of params for callback, expected 1 got %v", len(callback.Args))
	}

	local := NewEnv(env)
	for _, value := range m.array {
		local.varMap[callback.Args[0].Value] = &value
		result, err := e.evalNode(callback.Body, local)
		if err != nil {
			return nil, err
		}

		if result.Typ() != typRet {
			err = fmt.Errorf("the callback must return a value")
			return nil, err
		}
		if otob(result.(*retval).res) {
			return bton(true), nil
		}
	}
	return bton(false), nil
}
