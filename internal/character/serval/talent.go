package serval

import (
	"github.com/simimpact/srsim/internal/global/common"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	ServalTalentListener key.Modifier = "serval-talent-listener"
	ServalTalent         key.Attack   = "serval-talent"
)

func init() {
	modifier.Register(ServalTalentListener, modifier.Config{
		Listeners: modifier.Listeners{
			OnAfterHit: onAfterHit,
		},
	})
}

func onAfterHit(mod *modifier.Instance, e event.HitEnd) {
	if mod.Engine().HasModifier(e.Defender, common.Shock) {
		energyGain := 0.0
		if mod.Engine().HasBehaviorFlag(e.Defender, model.BehaviorFlag_STAT_DOT_ELECTRIC) {
			temp, _ := mod.Engine().CharacterInstance(mod.Owner())
			c := temp.(*char)
			// Every time Serval's Talent is triggered to deal Additional DMG, she regenerates 4 Energy.
			if c.info.Eidolon >= 2 {
				energyGain += 4.0
			}
			mod.Engine().Attack(info.Attack{
				Key:        ServalTalent,
				Source:     e.Attacker,
				Targets:    []key.TargetID{e.Defender},
				DamageType: model.DamageType_THUNDER,
				AttackType: model.AttackType_PURSUED,
				BaseDamage: info.DamageMap{
					model.DamageFormula_BY_ATK: talent[c.info.SkillLevelIndex()],
				},
				StanceDamage: 0.0,
				EnergyGain:   energyGain,
			})
		}
	}

}
