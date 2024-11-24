package afterthecharmonyfall

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
	charmonyFall            = "after-the-charmony-fall"
	charmonyFallBreakEffect = "after-the-charmony-fall-break-effect"
	charmonyFallSpd         = "after-the-charmony-fall-spd"
)

// Increases the wearer's Break Effect by 28%.
// After the wearer uses Ultimate, increases SPD by 8%, lasting for 2 turn(s).
func init() {
	lightcone.Register(key.AftertheCharmonyFall, lightcone.Config{
		CreatePassive: Create,
		Rarity:        4,
		Path:          model.Path_ERUDITION,
		Promotions:    promotions,
	})
	modifier.Register(charmonyFallSpd, modifier.Config{
		Stacking:      modifier.ReplaceBySource,
		StatusType:    model.StatusType_STATUS_BUFF,
		CanDispel:     true,
		BehaviorFlags: []model.BehaviorFlag{model.BehaviorFlag_STAT_SPEED_UP},
		Listeners: modifier.Listeners{
			OnAfterAction: addSpdBuff,
		},
	})
}

func Create(engine engine.Engine, owner key.TargetID, lc info.LightCone) {
	amt := 0.21 + 0.07*float64(lc.Imposition)
	engine.AddModifier(owner, info.Modifier{
		Name:   charmonyFallBreakEffect,
		Source: owner,
		Stats:  info.PropMap{prop.BreakEffect: amt},
		State:  float64(lc.Imposition),
	})
}

func addSpdBuff(mod *modifier.Instance, e event.ActionEnd) {
	if e.AttackType == model.AttackType_ULT {
		amt := 0.06 + 0.02*mod.State().(float64)
		mod.Engine().AddModifier(mod.Owner(), info.Modifier{
			Name:     charmonyFallSpd,
			Source:   mod.Owner(),
			Duration: 2,
			Stats:    info.PropMap{prop.SPDPercent: amt},
		})
	}
}
