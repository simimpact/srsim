package character

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

type Config struct {
	Create     func(engine engine.Engine, id key.TargetID, info info.Character) info.CharInstance
	Promotions []PromotionData
	Rarity     int
	Element    model.DamageType
	Path       model.Path
	MaxEnergy  float64
	Traces     TraceMap
	SkillInfo  SkillInfo
}

type TraceMap map[string]Trace

type Trace struct {
	Stat      prop.Property
	Amount    float64
	Ascension int
	Level     int
}

type PromotionData struct {
	MaxLevel   int
	ATKBase    float64
	ATKAdd     float64
	DEFBase    float64
	DEFAdd     float64
	HPBase     float64
	HPAdd      float64
	SPD        float64
	CritChance float64
	CritDMG    float64
	Aggro      float64
}

type SkillInfo struct {
	Attack    Attack
	Skill     Skill
	Ult       Ult
	Technique Technique
}

type SkillValidateFunc func(engine engine.Engine, char info.CharInstance) bool

type Attack struct {
	SPAdd      int
	TargetType model.TargetType
}

type Skill struct {
	SPNeed     int
	TargetType model.TargetType
	CanUse     SkillValidateFunc `exhaustruct:"optional"`
}

type Ult struct {
	TargetType model.TargetType
}

type Technique struct {
	TargetType model.TargetType
	IsAttack   bool
}

func (c Config) ascension(maxLvl, lvl int) int {
	if maxLvl <= 0 {
		maxLvl = lvl
	}

	for i, promo := range c.Promotions {
		if promo.MaxLevel >= maxLvl {
			return i
		}
	}
	return len(c.Promotions) - 1
}
