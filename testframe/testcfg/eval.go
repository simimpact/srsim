package testcfg

import (
	"context"
	"fmt"
	"strings"

	"github.com/simimpact/srsim/pkg/model"

	"github.com/simimpact/srsim/pkg/logic/gcs/eval"
	"github.com/simimpact/srsim/pkg/logic/gcs/parse"
)

const actions = `
set_default_action(dummy_character, attack(LowestHP));
register_skill_cb(dummy_character, fn () { return skill(LowestHP); });
register_ult_cb(dummy_character, fn () { return ult(LowestHP); });
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

func GenerateStandardTestEval(cfg *model.SimConfig) *eval.Eval {
	var action strings.Builder
	for _, char := range cfg.Characters {
		action.WriteString(fmt.Sprintf(`
set_default_action(%s, attack(LowestHP));
register_skill_cb(%s, fn () { return skill(LowestHP); });
register_ult_cb(%s, fn () { return ult(LowestHP); });
`, char.Key, char.Key, char.Key))
	}
	p := parse.New(action.String())
	res, err := p.Parse()
	if err != nil {
		panic(err)
	}
	return eval.New(context.Background(), res.Program)
}
