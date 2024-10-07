package tingyun

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	IsTingyun            = "tingyun-is-tingyun"
	TingyunNormalMonitor = "tingyun-normal-monitor"
	ProcTalent           = "tingyun-benediction-proc-talent"
)

func init() {
	modifier.Register(TingyunNormalMonitor, modifier.Config{
		Listeners: modifier.Listeners{
			OnAfterAttack: doProcTalent, // Talent's Pursued
		},
	})
	modifier.Register(IsTingyun, modifier.Config{})
}

func (c *char) initTalent() {
	c.engine.AddModifier(c.id, info.Modifier{
		Name:   IsTingyun,
		Source: c.id,
	})
}

func doProcTalent(mod *modifier.Instance, e event.AttackEnd) {
	if e.AttackType == model.AttackType_NORMAL && mod.Engine().HasModifier(mod.Owner(), IsTingyun) {
		st := mod.State().(*skillState)
		mod.Engine().Attack(info.Attack{
			Key:        ProcTalent,
			Targets:    e.Targets,
			Source:     mod.Owner(),
			AttackType: model.AttackType_PURSUED,
			DamageType: model.DamageType_THUNDER,
			BaseDamage: info.DamageMap{model.DamageFormula_BY_ATK: st.talentPursuedMltp},
		})
	}
}
