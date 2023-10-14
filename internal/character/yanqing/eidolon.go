package yanqing

import (
	"github.com/simimpact/srsim/internal/global/common"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Eidolon            = "yanqing-eidolon"
	E4                 = "yanqing-e4"
	E1      key.Attack = "yanqing-e1"
)

func (c *char) initEidolon() {
	modifier.Register(Eidolon, modifier.Config{
		StatusType: model.StatusType_UNKNOWN_STATUS,
		Listeners: modifier.Listeners{
			OnAfterAttack: E1Listener,
			OnHPChange:    E4Listener,
		},
	})
	modifier.Register(E4, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
	})
	c.engine.AddModifier(c.id, info.Modifier{
		Name:   Eidolon,
		Source: c.id,
	})
}
func E1Listener(mod *modifier.Instance, e event.AttackEnd) {
	cinfo, _ := mod.Engine().CharacterInfo(mod.Owner())
	if cinfo.Eidolon >= 1 && mod.Engine().HasModifier(e.Targets[0], common.Freeze) {
		mod.Engine().Attack(info.Attack{
			Key:        E1,
			Source:     e.Attacker,
			Targets:    e.Targets,
			DamageType: model.DamageType_ICE,
			AttackType: e.AttackType,
			BaseDamage: info.DamageMap{model.DamageFormula_BY_ATK: 0.6},
		})
	}
}
func E4Listener(mod *modifier.Instance, e event.HPChange) {
	cinfo, _ := mod.Engine().CharacterInfo(mod.Owner())
	if cinfo.Eidolon >= 4 && e.NewHPRatio >= 0.8 {
		mod.Engine().AddModifier(mod.Owner(), info.Modifier{
			Name:  E4,
			Stats: info.PropMap{prop.IcePEN: 0.12},
		})
	} else {
		mod.Engine().RemoveModifier(mod.Owner(), E4)
	}
}
func E6Listener(mod *modifier.Instance, target key.TargetID) {
	cinfo, _ := mod.Engine().CharacterInfo(mod.Owner())
	if cinfo.Eidolon >= 6 {
		mod.Engine().AddModifier(mod.Owner(), mod.ToModel())
	}
}
