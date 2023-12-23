package eval

import (
	"strconv"

	"github.com/simimpact/srsim/pkg/engine/target/evaltarget"
	"github.com/simimpact/srsim/pkg/logic"
	"github.com/simimpact/srsim/pkg/logic/gcs/ast"
)

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

func (o ObjTyp) String() string {
	typeMap := map[ObjTyp]string{
		typNull: "null",
		typNum:  "number",
		typStr:  "string",
		typMap:  "map",
		typAct:  "action",
		typFun:  "function",
		typBif:  "built-in function",
	}
	if name, ok := typeMap[o]; ok {
		return name
	}
	return "unknown"
}

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
		val logic.Action
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
	}
	return strconv.FormatInt(n.ival, 10)
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
func (r *retval) Typ() ObjTyp { return typRet }

// actionval.
func (a *actionval) Inspect() string {
	var targeteval string
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
