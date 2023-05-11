package eval

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/simimpact/srsim/pkg/gcs/ast"
)

func (e *Eval) initSysFuncs(env *Env) {
	// std funcs
	e.addSysFunc("rand", e.rand, env)
	e.addSysFunc("randnorm", e.randnorm, env)
	e.addSysFunc("print", e.print, env)
	e.addSysFunc("type", e.typeval, env)
}

func (e *Eval) addSysFunc(name string, f func(c *ast.CallExpr, env *Env) (Obj, error), env *Env) {
	var obj Obj = &bfuncval{Body: f}
	env.varMap[name] = &obj
}

func (e *Eval) print(c *ast.CallExpr, env *Env) (Obj, error) {
	//concat all args
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
	x := rand.Float64() // TODO: rand with a specific seed
	return &number{
		fval:    x,
		isFloat: true,
	}, nil
}

func (e *Eval) randnorm(c *ast.CallExpr, env *Env) (Obj, error) {
	x := rand.NormFloat64() // TODO: rand with a specific seed
	return &number{
		fval:    x,
		isFloat: true,
	}, nil
}

func (e *Eval) typeval(c *ast.CallExpr, env *Env) (Obj, error) {
	//type(var)
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
	case typFun:
		fallthrough
	case typBif:
		str = t.Inspect()
	}

	return &strval{str}, nil
}