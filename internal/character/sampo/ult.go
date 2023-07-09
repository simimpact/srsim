package sampo

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	SampoDOTTaken key.Modifier = "sampo-dot-taken"
	Ult           key.Attack   = "sampo-ult"
)

func init() {
	modifier.Register(SampoDOTTaken, modifier.Config{
		Stacking:   modifier.ReplaceBySource,
		StatusType: model.StatusType_STATUS_DEBUFF,
		Listeners: modifier.Listeners{
			OnBeforeBeingHitAll: onBeforeBeingHitAll,
		},
	})
}

var ultHits = []float64{0.25, 0.25, 0.25, 0.25}

func (c *char) Ult(target key.TargetID, state info.ActionState) {
	for i, hitRatio := range ultHits {
		c.engine.Attack(info.Attack{
			Key:        Ult,
			HitIndex:   i,
			Source:     c.id,
			Targets:    c.engine.Enemies(),
			DamageType: model.DamageType_WIND,
			AttackType: model.AttackType_ULT,
			BaseDamage: info.DamageMap{
				model.DamageFormula_BY_ATK: ult[c.info.UltLevelIndex()],
			},
			StanceDamage: 60.0,
			EnergyGain:   5,
			HitRatio:     hitRatio,
		})
	}

	state.EndAttack()
	c.a4()

	for _, trg := range c.engine.Enemies() {
		c.engine.AddModifier(trg, info.Modifier{
			Name:            SampoDOTTaken,
			Source:          c.id,
			Chance:          1,
			Duration:        2,
			State:           ultDotTaken[c.info.UltLevelIndex()],
			TickImmediately: true,
		})
	}
}

func onBeforeBeingHitAll(mod *modifier.Instance, e event.HitStart) {
	if e.Hit.AttackType == model.AttackType_DOT {
		e.Hit.Defender.AddProperty(prop.AllDamageTaken, mod.State().(float64))
	}
}
