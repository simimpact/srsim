package patienceisallyouneed

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/lightcone"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	patience key.Modifier = "patience-is-all-you-need"
	spdBuff  key.Modifier = "patience-is-all-you-need-spd-buff"
	erode                 = "erode"
)

type state struct {
	spdBuff, dotDmg float64
}

// Increases DMG dealt by the wearer by 24%. After every attack launched by wearer,
// their SPD increases by 4.8%, stacking up to 3 times. If the wearer hits an enemy target
// that is not afflicted by Erode, there is a 100% base chance to inflict Erode to the target.
// Enemies afflicted with Erode are also considered to be Shocked and will receive
// Lightning DoT at the start of each turn equal to 60% of the wearer's ATK, lasting for 1 turn(s).

func init() {
	lightcone.Register(key.PatienceIsAllYouNeed, lightcone.Config{
		CreatePassive: Create,
		Rarity:        5,
		Path:          model.Path_NIHILITY,
		Promotions:    promotions,
	})
	modifier.Register(patience, modifier.Config{
		Listeners: modifier.Listeners{
			OnAfterHit:    inflictErode,
			OnAfterAttack: addSpeedBuff,
		},
	})
	modifier.Register(spdBuff, modifier.Config{
		BehaviorFlags: []model.BehaviorFlag{
			model.BehaviorFlag_STAT_SPEED_UP,
		},
		Stacking: modifier.ReplaceBySource,
	})
	modifier.Register(erode, modifier.Config{
		Listeners: modifier.Listeners{
			OnPhase1: addDotDmg,
		},
		BehaviorFlags: []model.BehaviorFlag{
			model.BehaviorFlag_STAT_DOT_ELECTRIC,
			model.BehaviorFlag_STAT_DOT,
		},
		Stacking:          modifier.ReplaceBySource,
		TickMoment:        modifier.ModifierPhase1End,
		StatusType:        model.StatusType_STATUS_DEBUFF,
		MaxCount:          1,
		CountAddWhenStack: 1,
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	dmgBuff := 0.2 + 0.04*float64(lc.Imposition)
	modState := state{
		spdBuff: 0.04 + 0.008*float64(lc.Imposition),
		dotDmg:  0.5 + 0.1*float64(lc.Imposition),
	}

	engine.AddModifier(owner, info.Modifier{
		Name:   patience,
		Source: owner,
		Stats:  info.PropMap{prop.AllDamagePercent: dmgBuff},
		State:  &modState,
	})
}

func addSpeedBuff(mod *modifier.Instance, e event.AttackEnd) {
	state := mod.State().(*state)
	mod.Engine().AddModifier(mod.Owner(), info.Modifier{
		Name:              spdBuff,
		Source:            mod.Owner(),
		MaxCount:          3,
		CountAddWhenStack: 1,
		Stats:             info.PropMap{prop.SPDPercent: state.spdBuff},
	})
}

func inflictErode(mod *modifier.Instance, e event.HitEnd) {
	state := mod.State().(*state)
	// if already inflicted by erode, bypass
	if mod.Engine().HasModifier(e.Defender, erode) {
		return
	}
	mod.Engine().AddModifier(e.Defender, info.Modifier{
		Name:     erode,
		Source:   mod.Owner(),
		Duration: 1,
		Chance:   1,
		State:    state.dotDmg,
	})
}

func addDotDmg(mod *modifier.Instance) {
	dotDmg := mod.State().(float64)
	mod.Engine().Attack(info.Attack{
		Key:        erode,
		Targets:    []key.TargetID{mod.Owner()},
		Source:     mod.Source(),
		AttackType: model.AttackType_DOT,
		DamageType: model.DamageType_THUNDER,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: dotDmg,
		},
	})
}
