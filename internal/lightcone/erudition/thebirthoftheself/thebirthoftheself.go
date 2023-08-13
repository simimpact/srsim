package thebirthoftheself

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
	birth = "the-birth-of-the-self"
)

// Increases DMG dealt by the wearer's follow-up attacks by 24%.
// If the current HP of the target enemy is below or equal to 50%,
// increases DMG dealt by follow-up attacks by an extra 24%.

func init() {
	lightcone.Register(key.TheBirthoftheSelf, lightcone.Config{
		CreatePassive: Create,
		Rarity:        4,
		Path:          model.Path_ERUDITION,
		Promotions:    promotions,
	})
	modifier.Register(birth, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeHit: buffFollowUpAtk,
		},
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	dmgAmt := 0.18 + 0.06*float64(lc.Imposition)

	engine.AddModifier(owner, info.Modifier{
		Name:   birth,
		Source: owner,
		State:  dmgAmt,
	})
}

func buffFollowUpAtk(mod *modifier.Instance, e event.HitStart) {
	// if hit not follow-up : bypass.
	if e.Hit.AttackType != model.AttackType_INSERT {
		return
	}
	dmgAmt := mod.State().(float64)
	// 2x damage buff if hit enemy hp <50%
	if mod.Engine().HPRatio(e.Defender) <= 0.5 {
		dmgAmt = 2 * dmgAmt
	}
	e.Hit.Attacker.AddProperty(birth, prop.AllDamagePercent, dmgAmt)
}
