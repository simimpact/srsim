package inthenight

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
	IntheNight key.Modifier = "in_the_night"
)

type Amts struct {
	dmg float64
	cd  float64
}

// Increases the wearer's CRIT Rate by 18/21/24/27/30%. While the wearer is in
// battle, for every 10 SPD that exceeds 100, the DMG of the wearer's Basic ATK
// and Skill is increased by 6/7/8/9/10% and the CRIT DMG of their Ultimate is
// increased by 12/14/16/18/20%. This effect can stack up to 6 time(s).
func init() {
	lightcone.Register(key.IntheNight, lightcone.Config{
		CreatePassive: Create,
		Rarity:        5,
		Path:          model.Path_HUNT,
		Promotions:    promotions,
	})

	modifier.Register(IntheNight, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeHit: onBeforeHit,
		},
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	cr := 0.15 + 0.03*float64(lc.Imposition)
	dmg := 0.05 + 0.01*float64(lc.Imposition)
	cd := 0.1 + 0.02*float64(lc.Imposition)

	engine.AddModifier(owner, info.Modifier{
		Name:   IntheNight,
		Source: owner,
		Stats: info.PropMap{
			prop.CritChance: cr,
		},
		State: Amts{dmg: dmg, cd: cd},
	})
}

func onBeforeHit(mod *modifier.ModifierInstance, e event.HitStartEvent) {
	spd := e.Hit.Attacker.SPD()

	if spd >= 110 {
		// calculate stacks
		stacks := 0
		spd -= 100
		for i := 0; i < 6 && spd >= 10; i++ {
			spd -= 10
			stacks++
		}

		// modify damage
		switch e.Hit.AttackType {
		case model.AttackType_NORMAL, model.AttackType_SKILL:
			e.Hit.Attacker.AddProperty(prop.AllDamagePercent, float64(stacks)*mod.State().(Amts).dmg)
		case model.AttackType_ULT:
			e.Hit.Attacker.AddProperty(prop.CritDMG, float64(stacks)*mod.State().(Amts).cd)
		}
	}
}
