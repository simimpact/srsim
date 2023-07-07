package clara

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	// MAvatar_Klara_00_BPSkill_Revenge
	TalentMark key.Modifier = "clara-talent-mark"
	// MAvatar_Klara_00_Passive_DamageReduce
	TalentRes key.Modifier = "clara-talent-dmg-res"
)

// Under the protection of Svarog, DMG taken by Clara when hit by enemy
// attacks is reduced by *%. Svarog will mark enemies who attack Clara with
// his Mark of Counter and retaliate with a Counter, dealing Physical DMG
// equal to *% of Clara's ATK.
func init() {
	modifier.Register(TalentRes, modifier.Config{})
}

func (c *char) talentCounter(e event.AttackEnd) {
	// MAvatar_Klara_00_PassiveATK_Mark
    // follow-up attack + add counter mark
    c.engine.InsertAbility(info.Insert{
        Execute: func() {
            c.engine.Attack(info.Attack{
                Source: c.id,
                // TODO: defined attacker
				// Targets:    []key.TargetID{target},
                DamageType: model.DamageType_PHYSICAL,
                AttackType: model.AttackType_INSERT,
                // BaseDamage: ,
                StanceDamage: 30.0,
                EnergyGain: 5.0,
            })
        },
        Source: c.id,
        Priority: info.CharInsertAttackOthers,
        // dying check ?
        // AbortFlags: ,
    })
}
