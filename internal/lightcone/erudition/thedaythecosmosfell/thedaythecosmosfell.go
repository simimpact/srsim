package thedaythecosmosfell

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
	cosmosAtk     = "day-cosmos-fell-atk"
	cosmosCritDmg = "day-cosmos-fell-cdmg"
)

// Increases the wearer's ATK by 16%.
// When the wearer uses an attack and at least 2 attacked enemies have the corresponding Weakness,
// the wearer's CRIT DMG increases by 20%, lasting for 2 turn(s).
func init() {
	lightcone.Register(key.TheDayTheCosmosFell, lightcone.Config{
		CreatePassive: Create,
		Rarity:        4,
		Path:          model.Path_ERUDITION,
		Promotions:    promotions,
	})
	modifier.Register(cosmosCritDmg, modifier.Config{
		Stacking:   modifier.ReplaceBySource,
		StatusType: model.StatusType_STATUS_BUFF,
		Listeners: modifier.Listeners{
			OnAfterAttack: weaknessCheck,
		},
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	amt := 0.14 + 0.02*float64(lc.Imposition)
	engine.AddModifier(owner, info.Modifier{
		Name:   cosmosAtk,
		Source: owner,
		Stats:  info.PropMap{prop.ATKPercent: amt},
		State:  lc.Imposition,
	})
}

func weaknessCheck(mod *modifier.Instance, e event.AttackEnd) {
	matchingWeakness := 0
	charInfo, _ := mod.Engine().CharacterInfo(mod.Owner())
	for _, target := range e.Targets {
		if mod.Engine().Stats(target).IsWeakTo(charInfo.Element) {
			matchingWeakness++
		}
	}

	amt := 0.15 + 0.05*mod.State().(float64)
	if matchingWeakness >= 2 {
		mod.Engine().AddModifier(mod.Owner(), info.Modifier{
			Name:     cosmosCritDmg,
			Source:   mod.Owner(),
			Stats:    info.PropMap{prop.CritDMG: amt},
			Duration: 2,
		})
	}
}
