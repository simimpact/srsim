package trailblazerimaginary

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	talent         key.Reason   = "trailblazerimaginary-talent"
	A2             key.Reason   = "trailblazerimaginary-a2"
	E2Buff         key.Modifier = "trailblazerimaginary-e2buff"
	E4ListenerBuff key.Modifier = "trailblazerimaginary-e4listener"
	E4Buff         key.Modifier = "trailblazerimaginary-e4buff"
)

func init() {
	modifier.Register(E2Buff, modifier.Config{
		Duration:   3,
		StatusType: model.StatusType_STATUS_BUFF,
	})
	modifier.Register(E4ListenerBuff, modifier.Config{
		Listeners: modifier.Listeners{
			OnPropertyChange: E4Listener,
			OnRemove:         E4Remove,
		},
		StatusType: model.StatusType_UNKNOWN_STATUS,
	})
	modifier.Register(E4Buff, modifier.Config{
		Stacking:   modifier.Replace,
		StatusType: model.StatusType_STATUS_BUFF,
	})
}

func (c *char) buffListener(e event.TurnStart) {
	if e.Active != c.id {
		return
	}
	c.ultLifeTime--
	if c.ultLifeTime <= 0 {
		for _, target := range c.engine.Characters() {
			c.engine.RemoveModifier(target, BackupDancer)
		}
	}
}

func (c *char) talentListener(e event.StanceBreak) {
	c.engine.ModifyEnergy(info.ModifyAttribute{
		Key:    talent,
		Source: c.id,
		Target: c.id,
		Amount: talentEnergy[c.info.TalentLevelIndex()],
	})
	if c.info.Traces["103"] {
		c.engine.ModifyGaugeNormalized(info.ModifyAttribute{
			Key:    A2,
			Target: e.Target,
			Source: c.id,
			Amount: 0.3,
		})
	}
}

func E4Listener(mod *modifier.Instance) {
	for _, target := range mod.Engine().Characters() {
		if target == mod.Owner() {
			continue
		}
		mod.Engine().AddModifier(target, info.Modifier{
			Name:   E4Buff,
			Source: mod.Owner(),
			Stats: info.PropMap{
				prop.BreakEffect: 0.15 * mod.Engine().Stats(mod.Owner()).BreakEffect(),
			},
		})
	}
}

func E4Remove(mod *modifier.Instance) {
	for _, target := range mod.Engine().Characters() {
		if target == mod.Owner() {
			continue
		}
		mod.Engine().RemoveModifier(target, E4Buff)
	}
}
