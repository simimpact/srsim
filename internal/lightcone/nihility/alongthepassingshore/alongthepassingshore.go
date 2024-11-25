package alongthepassingshore

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
	Check        = "along-the-passing-shore"
	MirageFizzle = "along-the-passing-shore-mirage-fizzle"
	Cooldown     = "along-the-passing-shore-mirage-fizzle-cooldown"
)

// Increases the wearer's CRIT DMG by 36/42/48/54/60%. When the wearer hits an enemy target,
// inflicts Mirage Fizzle on the enemy, lasting for 1 turn. Each time the wearer attacks,
// this effect can only trigger 1 time on each target. The wearer deals 24/28/32/36/40% increased DMG to targets
// afflicted with Mirage Fizzle, and the DMG dealt by Ultimate additionally increases by 24/28/32/36/40%.

func init() {
	lightcone.Register(key.AlongthePassingShore, lightcone.Config{
		CreatePassive: Create,
		Rarity:        5,
		Path:          model.Path_NIHILITY,
		Promotions:    promotions,
	})

	modifier.Register(Check, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeHit:    applyMirageFizzle,
			OnBeforeHitAll: applyDmgBonus,
			OnAfterAttack:  removeCooldown,
		},
	})

	modifier.Register(MirageFizzle, modifier.Config{
		Stacking:   modifier.ReplaceBySource,
		StatusType: model.StatusType_STATUS_DEBUFF,
		CanDispel:  true,
	})

	modifier.Register(Cooldown, modifier.Config{})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	cdmgAmt := 0.3 + 0.06*float64(lc.Imposition)
	dmgAmt := 0.2 + 0.04*float64(lc.Imposition)
	engine.AddModifier(owner, info.Modifier{
		Name:   Check,
		Source: owner,
		Stats:  info.PropMap{prop.CritDMG: cdmgAmt},
		State:  dmgAmt,
	})
}

func applyMirageFizzle(mod *modifier.Instance, e event.HitStart) {
	if mod.Engine().HasModifierFromSource(e.Defender, mod.Owner(), Cooldown) {
		return
	}
	success, _ := mod.Engine().AddModifier(e.Defender, info.Modifier{
		Name:     MirageFizzle,
		Source:   mod.Owner(),
		Duration: 1,
	})
	if success {
		mod.Engine().AddModifier(e.Defender, info.Modifier{
			Name:   Cooldown,
			Source: mod.Owner(),
		})
	}
}

func applyDmgBonus(mod *modifier.Instance, e event.HitStart) {
	if mod.Engine().HasModifierFromSource(e.Defender, mod.Owner(), MirageFizzle) {
		dmgBonus := mod.State().(float64)
		if e.Hit.AttackType == model.AttackType_ULT {
			dmgBonus *= 2
		}
		e.Hit.Attacker.AddProperty(MirageFizzle, prop.AllDamagePercent, dmgBonus)
	}
}

func removeCooldown(mod *modifier.Instance, e event.AttackEnd) {
	for _, trg := range mod.Engine().Enemies() {
		mod.Engine().RemoveModifier(trg, Cooldown)
	}
}
