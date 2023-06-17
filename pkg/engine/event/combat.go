package event

import (
	"bytes"
	"fmt"
	"github.com/simimpact/srsim/pkg/engine/event/handler"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
	"strconv"
)

type AttackStartEventHandler = handler.EventHandler[AttackStartEvent]
type AttackStartEvent struct {
	Attacker   key.TargetID
	Targets    []key.TargetID
	AttackType model.AttackType
	DamageType model.DamageType
}

func (e AttackStartEvent) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString("{\"Event\": \"AttackStartEvent\",")
	buf.WriteString("\"Attacker\": " + strconv.Itoa(int(e.Attacker)))
	buf.WriteString(fmt.Sprintf(",\"Targets\": %v", e.Targets))
	buf.WriteString(",\"AttackType\": \"" + e.AttackType.String())
	buf.WriteString("\",\"DamageType\": \"" + e.DamageType.String())
	buf.WriteByte('"')
	buf.WriteByte('}')
	fmt.Println("Hello")
	fmt.Println(buf.String())
	return buf.Bytes(), nil
}

type AttackEndEventHandler = handler.EventHandler[AttackEndEvent]
type AttackEndEvent struct {
	Attacker   key.TargetID
	Targets    []key.TargetID
	AttackType model.AttackType
	DamageType model.DamageType
}

type HitStartEventHandler = handler.EventHandler[HitStartEvent]
type HitStartEvent struct {
	Attacker key.TargetID
	Defender key.TargetID
	Hit      *info.Hit
}

type HitEndEventHandler = handler.EventHandler[HitEndEvent]
type HitEndEvent struct {
	Attacker         key.TargetID
	Defender         key.TargetID
	AttackType       model.AttackType
	DamageType       model.DamageType
	BaseDamage       float64
	BonusDamage      float64
	TotalDamage      float64
	HPDamage         float64
	ShieldDamage     float64
	HPRatioRemaining float64
	IsCrit           bool
	UseSnapshot      bool
}

type HealStartEventHandler = handler.MutableEventHandler[HealStartEvent]
type HealStartEvent struct {
	Target      *info.Stats
	Healer      *info.Stats
	BaseHeal    info.HealMap
	HealValue   float64
	UseSnapshot bool
}

type HealEndEventHandler = handler.EventHandler[HealEndEvent]
type HealEndEvent struct {
	Target             key.TargetID
	Healer             key.TargetID
	HealAmount         float64
	OverflowHealAmount float64
	UseSnapshot        bool
}
