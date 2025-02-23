package inherentlyunjustdestiny

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
	Check      = "inherently-unjust-destiny"
	ShieldCdmg = "inherently-unjust-destiny-cdmg-buff"
	AllIn      = "inherently-unjust-destiny-vuln-debuff"
)

type state struct {
	chance         float64
	vuln           float64
	applieddynamic bool
}

// Increases the wearer's DEF by 40/46/52/58/64%. When the wearer provides a Shield to an ally,
// the wearer's CRIT DMG increases by 40/46/52/58/64%, lasting for 2 turn(s).
// When the wearer's follow-up attack hits an enemy target, there is a 100/115/130/145/160% base chance
// to increase the DMG taken by the attacked enemy target by 10/11.5/13/14.5/16%, lasting for 2 turn(s).

func init() {
	lightcone.Register(key.InherentlyUnjustDestiny, lightcone.Config{
		CreatePassive: Create,
		Rarity:        5,
		Path:          model.Path_PRESERVATION,
		Promotions:    promotions,
	})

	modifier.Register(Check, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeHitAll: applyVuln,
		},
	})

	modifier.Register(ShieldCdmg, modifier.Config{
		Stacking:   modifier.Replace,
		StatusType: model.StatusType_STATUS_BUFF,
		CanDispel:  true,
	})

	modifier.Register(AllIn, modifier.Config{
		Stacking:   modifier.ReplaceBySource,
		StatusType: model.StatusType_STATUS_DEBUFF,
		CanDispel:  true,
		Listeners: modifier.Listeners{
			OnBeforeBeingHitAll: func(mod *modifier.Instance, e event.HitStart) {
				e.Hit.Defender.AddProperty(AllIn, prop.AllDamageTaken, mod.State().(state).vuln)
			},
		},
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	defncdmgAmt := 0.34 + 0.06*float64(lc.Imposition)
	chanceAmt := 0.85 + 0.15*float64(lc.Imposition)
	vulnAmt := chanceAmt * 0.1

	engine.AddModifier(owner, info.Modifier{
		Name:   Check,
		Source: owner,
		Stats:  info.PropMap{prop.DEFPercent: defncdmgAmt},
		State: &state{
			chance:         chanceAmt,
			vuln:           vulnAmt,
			applieddynamic: false,
		},
	})

	// Special shield listener for owner providing shields
	engine.Events().ShieldAdded.Subscribe(func(event event.ShieldAdded) {
		if event.Info.Source == owner {
			engine.AddModifier(owner, info.Modifier{
				Name:     ShieldCdmg,
				Source:   owner,
				Duration: 2,
				Stats:    info.PropMap{prop.CritDMG: defncdmgAmt},
			})
		}
	})
}

func applyVuln(mod *modifier.Instance, e event.HitStart) {
	st := mod.State().(*state)
	if e.Hit.AttackType == model.AttackType_INSERT {
		mod.Engine().AddModifier(e.Defender, info.Modifier{
			Name:     AllIn,
			Source:   mod.Owner(),
			Duration: 2,
			Stats:    info.PropMap{prop.AllDamageTaken: st.vuln},
			Chance:   st.chance,
		})
	}
	// Logic to apply the debuff on the first hit, ensuring that this hit also benefits from the vuln effect
	if !mod.Engine().HasModifierFromSource(e.Defender, mod.Owner(), AllIn) {
		st.applieddynamic = false
	} else if !st.applieddynamic {
		e.Hit.Defender.AddProperty(AllIn, prop.AllDamageTaken, st.vuln)
		st.applieddynamic = true
	}
}
