package theunreachableside

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
	side    key.Modifier = "the-unreachable-side"
	dmgBuff key.Modifier = "the-unreachable-side-dmg-buff"
)

// Increases the wearer's CRIT rate by 18% and increases their Max HP by 18%.
// When the wearer is attacked or consumes their own HP, their DMG increases by 24%.
// This effect is removed after the wearer uses an attack.

func init() {
	lightcone.Register(key.TheUnreachableSide, lightcone.Config{
		CreatePassive: Create,
		Rarity:        5,
		Path:          model.Path_DESTRUCTION,
		Promotions:    promotions,
	})
	modifier.Register(side, modifier.Config{
		Listeners: modifier.Listeners{
			OnAfterBeingAttacked: buffOnAttacked,
			OnHPChange:           buffOnHPConsume,
		},
	})
	modifier.Register(dmgBuff, modifier.Config{
		Listeners: modifier.Listeners{
			OnAfterAttack: removeBuff,
		},
		StatusType: model.StatusType_STATUS_BUFF,
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	crNHpAmt := 0.15 + 0.03*float64(lc.Imposition)
	dmgAmt := 0.20 + 0.04*float64(lc.Imposition)
	engine.AddModifier(owner, info.Modifier{
		Name:   side,
		Source: owner,
		Stats: info.PropMap{
			prop.CritChance: crNHpAmt,
			prop.HPPercent:  crNHpAmt,
		},
		State: dmgAmt,
	})
}

func buffOnAttacked(mod *modifier.Instance, e event.AttackEnd) {
	dmgAmt := mod.State().(float64)
	// add damage buff here.
	addDmgBuff(mod.Engine(), dmgBuff, mod.Owner(), dmgAmt)
}

func buffOnHPConsume(mod *modifier.Instance, e event.HPChange) {
	// bypass if: hpchangebydmg is true, hpchange source isn't lc holder, hpchange increase hp.
	if e.IsHPChangeByDamage || e.Target != mod.Owner() || e.NewHP > e.OldHP {
		return
	}
	dmgAmt := mod.State().(float64)
	// add damage buff here.
	addDmgBuff(mod.Engine(), dmgBuff, mod.Owner(), dmgAmt)
}

func removeBuff(mod *modifier.Instance, e event.AttackEnd) {
	mod.RemoveSelf()
}

func addDmgBuff(engine engine.Engine, dmgBuff key.Modifier, owner key.TargetID, dmgAmt float64) {
	engine.AddModifier(owner, info.Modifier{
		Name:   dmgBuff,
		Source: owner,
		Stats:  info.PropMap{prop.AllDamagePercent: dmgAmt},
	})
}
