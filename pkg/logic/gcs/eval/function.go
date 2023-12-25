package eval

import (
	"fmt"

	"github.com/simimpact/srsim/pkg/logic/gcs/ast"
)

func (e *Eval) validateArguments(args []ast.Expr, env *Env, objTypes ...ObjTyp) ([]Obj, error) {
	if len(args) != len(objTypes) {
		return nil, fmt.Errorf("invalid number of params, expected %v got %v", len(objTypes), len(args))
	}
	if len(objTypes) == 0 {
		return nil, nil
	}

	objs := make([]Obj, 0, len(objTypes))
	for i, objType := range objTypes {
		obj, err := e.evalExpr(args[i], env)
		if err != nil {
			return nil, err
		}
		if obj.Typ() != objType {
			return nil, fmt.Errorf("the function argument #%v should evaluate to %v, got %v", i+1, objType.String(), obj.Typ().String())
		}
		objs = append(objs, obj)
	}
	return objs, nil
}
