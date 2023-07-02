package testcfg

import (
	"context"

	"github.com/simimpact/srsim/pkg/logic/gcs/eval"
	"github.com/simimpact/srsim/pkg/logic/gcs/parse"
)

const actions = `
set_default_action(danheng, attack(LowestHP));
register_skill_cb(danheng, fn () { return skill(LowestHP); });
register_ult_cb(danheng, fn () { return ult(LowestHP); });
`

func StandardTestEval() *eval.Eval {
	p := parse.New(actions)
	res, err := p.Parse()
	if err != nil {
		panic(err)
	}
	return eval.New(context.Background(), res.Program)
}
