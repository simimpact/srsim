package event

import (
	"github.com/simimpact/srsim/pkg/engine/event/handler"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

type AttackStartEventHandler = handler.EventHandler[AttackStart]
type AttackStart struct {
	Attacker   key.TargetID     `json:"attacker"`
	Targets    []key.TargetID   `json:"targets"`
	AttackType model.AttackType `json:"attack_type"`
	DamageType model.DamageType `json:"damage_type"`
}

type AttackEndEventHandler = handler.EventHandler[AttackEnd]
type AttackEnd struct {
	Attacker   key.TargetID     `json:"attacker"`
	Targets    []key.TargetID   `json:"targets"`
	AttackType model.AttackType `json:"attack_type"`
	DamageType model.DamageType `json:"damage_type"`
}

type HitStartEventHandler = handler.EventHandler[HitStart]
type HitStart struct {
	Attacker key.TargetID `json:"attacker"`
	Defender key.TargetID `json:"defender"`
	Hit      *info.Hit    `json:"hit"`
}

type HitEndEventHandler = handler.EventHandler[HitEnd]
type HitEnd struct {
	Attacker            key.TargetID     `json:"attacker"`
	Defender            key.TargetID     `json:"defender"`
	AttackType          model.AttackType `json:"attack_type"`
	DamageType          model.DamageType `json:"damage_type"`
	BaseDamage          float64          `json:"base_damage"`
	BonusDamage         float64          `json:"bonus_damage"`
	DefenceMultiplier   float64          `json:"defence_multiplier"`
	Resistance          float64          `json:"resistance"`
	Vulnerability       float64          `json:"vulnerability"`
	ToughnessMultiplier float64          `json:"toughness_multiplier"`
	Fatigue             float64          `json:"fatigue"`
	AllDamageReduce     float64          `json:"all_damage_reduce"`
	CritDamage          float64          `json:"crit_damage"`
	TotalDamage         float64          `json:"total_damage"`
	HPDamage            float64          `json:"hp_damage"`
	ShieldDamage        float64          `json:"shield_damage"`
	HPRatioRemaining    float64          `json:"hp_ratio_remaining"`
	IsCrit              bool             `json:"is_crit"`
	UseSnapshot         bool             `json:"use_snapshot"`
}

type HealStartEventHandler = handler.MutableEventHandler[HealStart]
type HealStart struct {
	Target      *info.Stats  `json:"target"`
	Healer      *info.Stats  `json:"healer"`
	BaseHeal    info.HealMap `json:"base_heal"`
	HealValue   float64      `json:"heal_value"`
	UseSnapshot bool         `json:"use_snapshot"`
}

type HealEndEventHandler = handler.EventHandler[HealEnd]
type HealEnd struct {
	Target             key.TargetID `json:"target"`
	Healer             key.TargetID `json:"healer"`
	HealAmount         float64      `json:"heal_amount"`
	OverflowHealAmount float64      `json:"overflow_heal_amount"`
	UseSnapshot        bool         `json:"use_snapshot"`
}
