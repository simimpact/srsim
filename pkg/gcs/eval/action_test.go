package eval

import (
	"context"
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"
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

func TestCharAdd(t *testing.T) {
	p := parse.New(actions)
	res, err := p.Parse()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	eval := Eval{AST: res.Program}
	ctx := context.Background()
	ok := eval.Init(ctx)
	if !ok {
		err := <- eval.Err
		panic(err)
	}

	// skill
	fmt.Println("skill")
	spew.Config.Dump(eval.NextAction(0))
	spew.Config.Dump(eval.NextAction(1))
	spew.Config.Dump(eval.NextAction(1))
	
	// burst
	fmt.Println("burst")
	spew.Config.Dump(eval.BurstCheck())
	spew.Config.Dump(eval.BurstCheck())
	spew.Config.Dump(eval.BurstCheck())
}