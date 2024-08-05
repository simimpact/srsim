package kafka

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	E1 = "kafka-e1"
	E2 = "kafka-e2"
	E4 = "kafka-e4"
)

func init() {
	modifier.Register(E1, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeHitAll: E1Listener,
		},
		StatusType: model.StatusType_STATUS_DEBUFF,
	})

	modifier.Register(E2, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
	})

	modifier.Register(E4, modifier.Config{
		Listeners: modifier.Listeners{
			OnAfterHitAll: E4Listen,
		},
	})
}

func (c *char) initEidolons() {
	for _, ally := range c.engine.Characters() {
		c.engine.AddModifier(ally, info.Modifier{
			Name:   E2,
			Source: c.id,
			Stats: info.PropMap{
				prop.DOTDamagePercent: 0.25,
			},
		})
	}
}

func E1Listener(mod *modifier.Instance, e event.HitStart) {
	if e.Hit.AttackType == model.AttackType_DOT {
		e.Hit.Attacker.AddProperty(E1, prop.AllDamageTaken, 0.30)
	}
}

/*
DM uses a listener like this but for OnAfterBeingHitAll and then has another modifier/listener that would then apply this
to any new enemies that spawn. Current impl potentially more performant than original approach?
func E4Listener(mod *modifier.Instance, e event.HitEnd) {
	isDot := e.AttackType == model.AttackType_DOT
	isShock := e.DamageType == model.DamageType_THUNDER
	ownedByKafka := e.Attacker == mod.Source()
	if isDot && isShock && ownedByKafka {
		mod.Engine().ModifyEnergy(info.ModifyAttribute{
			Key: E4,
			Target: mod.Owner(),
			Source: mod.Owner(),
			Amount: 2,
		})
	}
}*/

func E4Listen(mod *modifier.Instance, e event.HitEnd) {
	isDot := e.AttackType == model.AttackType_DOT
	isShock := e.DamageType == model.DamageType_THUNDER
	if isDot && isShock {
		mod.Engine().ModifyEnergy(info.ModifyAttribute{
			Key:    E4,
			Target: mod.Owner(),
			Source: mod.Owner(),
			Amount: 2,
		})
	}
}
