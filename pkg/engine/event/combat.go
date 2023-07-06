package event

import (
	"github.com/simimpact/srsim/pkg/engine/event/handler"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

type AttackStartEventHandler = handler.EventHandler[AttackStart]
type AttackStart struct {
	Attacker   key.TargetID
	Targets    []key.TargetID
	AttackType model.AttackType
	DamageType model.DamageType
}

type AttackEndEventHandler = handler.EventHandler[AttackEnd]
type AttackEnd struct {
	Attacker   key.TargetID
	Targets    []key.TargetID
	AttackType model.AttackType
	DamageType model.DamageType
}

type HitStartEventHandler = handler.EventHandler[HitStart]
type HitStart struct {
	Attacker key.TargetID
	Defender key.TargetID
	Hit      *info.Hit
}

type HitEndEventHandler = handler.EventHandler[HitEnd]
type HitEnd struct {
	Attacker            key.TargetID
	Defender            key.TargetID
	AttackType          model.AttackType
	DamageType          model.DamageType
	BaseDamage          float64
	BonusDamage         float64
	DefenceMultiplier   float64
	Resistance          float64
	Vulnerability       float64
	ToughnessMultiplier float64
	Fatigue             float64
	AllDamageReduce     float64
	CritDamage          float64
	TotalDamage         float64
	HPDamage            float64
	ShieldDamage        float64
	HPRatioRemaining    float64
	IsCrit              bool
	UseSnapshot         bool
}

type HealStartEventHandler = handler.MutableEventHandler[HealStart]
type HealStart struct {
	Target      *info.Stats
	Healer      *info.Stats
	BaseHeal    info.HealMap
	HealValue   float64
	UseSnapshot bool
}

type HealEndEventHandler = handler.EventHandler[HealEnd]
type HealEnd struct {
	Target             key.TargetID
	Healer             key.TargetID
	HealAmount         float64
	OverflowHealAmount float64
	UseSnapshot        bool
}
