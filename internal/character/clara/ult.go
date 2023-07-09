package clara

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Ult        key.Modifier = "clara-ult" // MAvatar_Klara_00_Ultra_WarriorMode
	UltCounter key.Modifier = "clara-ult-enhanced-counter"
)

func init() {
	modifier.Register(Ult, modifier.Config{
		BehaviorFlags: []model.BehaviorFlag{model.BehaviorFlag_STAT_BURST},
		Stacking:      modifier.Refresh,
		StatusType:    model.StatusType_STATUS_BUFF,
		Duration:      2,
	})

	// stack count is managed in talent.go
	modifier.Register(UltCounter, modifier.Config{
		Stacking: modifier.ReplaceBySource,
	})
}

// After Clara uses Ultimate, DMG dealt to her is reduced by an extra *%, and
// she has greatly increased chances of being attacked by enemies for 2 turn(s).
// In addition, Svarog's Counter is enhanced. When an ally is attacked, Svarog
// immediately launches a Counter, and its DMG multiplier against the enemy
// increases by *%. Enemies adjacent to it take 50% of the DMG dealt to the
// target enemy. Enhanced Counter(s) can take effect 2 time(s).
func (c *char) Ult(target key.TargetID, state info.ActionState) {
	c.e2()

	c.engine.AddModifier(c.id, info.Modifier{
		Name:            Ult,
		Source:          c.id,
		Duration:        2,
		TickImmediately: true,
		Stats:           info.PropMap{prop.AggroPercent: 5, prop.AllDamageReduce: ultCut[c.info.UltLevelIndex()]},
	})

	enhanceCounterNum := 2.0
	if c.info.Eidolon >= 6 {
		enhanceCounterNum = 3.0
	}

	c.engine.AddModifier(c.id, info.Modifier{
		Name:     UltCounter,
		Source:   c.id,
		Count:    enhanceCounterNum,
		MaxCount: 3,
	})

	c.engine.ModifyEnergy(c.id, 5.0)
}
