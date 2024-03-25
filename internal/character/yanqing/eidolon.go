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
	E4            = "yanqing-e4"
	E1 key.Attack = "yanqing-e1"
)

func init() {
	modifier.Register(E4, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
	})
}

func (c *char) E1Listener(e event.AttackEnd) {
	if c.info.Eidolon >= 1 && c.engine.HasModifier(e.Targets[0], common.Freeze) {
		c.engine.Attack(info.Attack{
			Key:        E1,
			Source:     e.Attacker,
			Targets:    e.Targets,
			DamageType: model.DamageType_ICE,
			AttackType: e.AttackType,
			BaseDamage: info.DamageMap{model.DamageFormula_BY_ATK: 0.6},
		})
	}
}

func (c *char) E4Listener(e event.HPChange) {
	if c.info.Eidolon >= 4 && e.NewHPRatio >= 0.8 {
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   E4,
			Stats:  info.PropMap{prop.IcePEN: 0.12},
			Source: c.id,
		})
	} else {
		c.engine.RemoveModifier(c.id, E4)
	}
}

func E6Listener(mod *modifier.Instance, target key.TargetID) {
	cinfo, _ := mod.Engine().CharacterInfo(mod.Owner())
	if cinfo.Eidolon >= 6 {
		mod.Engine().AddModifier(mod.Owner(), mod.ToModel())
	}
}
