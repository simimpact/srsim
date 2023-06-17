package testcfg

import (
	"context"
	"github.com/simimpact/srsim/pkg/gcs/eval"
	"github.com/simimpact/srsim/pkg/gcs/parse"
)

const actions = `
register_skill_cb(0, fn () { return true; });

let value = true;
register_skill_cb(1, fn () {
	value = !value;
	return value;
});

let use = false;
register_burst_cb(0, fn () {
	if use {
		use = false;
		return true;
	}
	return false;
});
register_burst_cb(1, fn () {
	use = true;
	return true;
});
`

func StandardTestEval() *eval.Eval {
	p := parse.New(actions)
	res, err := p.Parse()
	if err != nil {
		panic(err)
	}

	e := &eval.Eval{AST: res.Program}
	ctx := context.Background()
	ok := e.Init(ctx)
	if !ok {
		err := <-e.Err
		panic(err)
	}
	return e
}
