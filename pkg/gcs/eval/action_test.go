package eval

import (
	"context"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/simimpact/srsim/pkg/engine/action"
	"github.com/simimpact/srsim/pkg/engine/target/evaltarget"
	"github.com/simimpact/srsim/pkg/gcs/parse"
	"github.com/simimpact/srsim/pkg/key"
)

const actions = `
register_skill_cb(0, fn () { return skill(LowestHP); });

let skill_pressed = true;
set_default_action(1, attack(LowestHP));
register_skill_cb(1, fn () {
    skill_pressed = !skill_pressed;
    if skill_pressed {
        return skill(First);
    }
	return null;
});

let use = false;
register_ult_cb(0, fn () {
	if use {
		use = false;
		return ult(First);
	}
	use = true;
	return null;
});

// use after skill
register_ult_cb(1, fn () {
    if skill_pressed {
		skill_pressed = false;
        return ult(LowestHP);
    }
	return null;
});
`

func TestCharAdd(t *testing.T) {
	p := parse.New(actions)
	res, err := p.Parse()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	eval := New(context.Background(), res.Program)
	err = eval.Init()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	// skill
	act, err := eval.NextAction(0)
	assertValidSkill(t, act, err, action.Action{
		Type:            key.ActionSkill,
		Target:          0,
		TargetEvaluator: evaltarget.LowestHP,
	})
	act, err = eval.NextAction(1)
	assertValidSkill(t, act, err, action.Action{
		Type:            key.ActionAttack,
		Target:          1,
		TargetEvaluator: evaltarget.LowestHP,
	})
	act, err = eval.NextAction(1)
	assertValidSkill(t, act, err, action.Action{
		Type:            key.ActionSkill,
		Target:          1,
		TargetEvaluator: evaltarget.First,
	})

	// ult
	acts, err := eval.UltCheck()
	assertValidUlt(t, acts, err, []action.Action{
		{
			Type:            key.ActionUlt,
			Target:          1,
			TargetEvaluator: evaltarget.LowestHP,
		},
	})
	acts, err = eval.UltCheck()
	assertValidUlt(t, acts, err, []action.Action{
		{
			Type:            key.ActionUlt,
			Target:          0,
			TargetEvaluator: evaltarget.First,
		},
	})
	acts, err = eval.UltCheck()
	assertValidUlt(t, acts, err, []action.Action{})
}

func assertValidSkill(t *testing.T, act action.Action, err error, validact action.Action) {
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if act.Target != validact.Target || act.TargetEvaluator != validact.TargetEvaluator || act.Type != validact.Type {
		t.Errorf("incorrect action %s. should be: %s", spew.Sprint(act), spew.Sprint(validact))
		t.FailNow()
	}
}

func assertValidUlt(t *testing.T, acts []action.Action, err error, validacts []action.Action) {
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if len(acts) != len(validacts) {
		t.Errorf("incorrect number of action (%d). should be %d", len(acts), len(validacts))
		t.FailNow()
	}

	for i, k := range acts {
		if k.Target != validacts[i].Target || k.TargetEvaluator != validacts[i].TargetEvaluator || k.Type != validacts[i].Type {
			t.Errorf("incorrect action %s. should be: %s", spew.Sprint(k), spew.Sprint(validacts[i]))
			t.FailNow()
		}
	}
}
