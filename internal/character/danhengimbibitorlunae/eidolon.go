package danhengimbibitorlunae

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	E6Effect            = "danhengimbibitorlunae-e6effect" // MAvatar_DanHengIL_00_Rank06_ImaginaryPenetrate
	E6       key.Reason = "danhengimbibitorlunae-e6"
)

// E1 increase talent max stack by 4 and gain 1 extra stack when hit
// E2 100% forward and 1 more ult stack
// E4 skill buff 1 more turn
// E6 20% imaginary pen for attack3 when ally use ult, max 3 stack

func init() {
	// imaginary pen for attack3,change by ally ult count
	modifier.Register(E6Effect, modifier.Config{
		StatusType:        model.StatusType_STATUS_BUFF,
		Stacking:          modifier.ReplaceBySource,
		MaxCount:          3,
		CountAddWhenStack: 1,
		Listeners: modifier.Listeners{
			OnBeforeHitAll: E6BeforeHit,
			OnAfterAttack:  RemoveE6,
		},
		CanModifySnapshot: true,
	})
}

func E6BeforeHit(mod *modifier.Instance, e event.HitStart) {
	temp, _ := mod.Engine().CharacterInstance(mod.Source())
	if temp.(*char).attackLevel == 3 {
		e.Hit.Attacker.AddProperty(E6, prop.ImaginaryPEN, 0.2*mod.Count())
	}
}
func RemoveE6(mod *modifier.Instance, e event.AttackEnd) {
	mod.RemoveSelf()
}

// count ally ult
func (c *char) E6ActionEndListener(e event.ActionEnd) {
	if e.AttackType == model.AttackType_ULT && e.Owner != c.id && c.info.Eidolon >= 6 {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   E6Effect,
			Source: c.id,
		})
	}
}
