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
	BackupDancer   key.Modifier = "trailblazerimaginary-backupdancer"
	SuperBreak     key.Attack   = "trailblazerimaginary-superbreak"
	talent         key.Reason   = "trailblazerimaginary-talent"
	A2             key.Reason   = "trailblazerimaginary-a2"
	E2Buff         key.Modifier = "trailblazerimaginary-e2buff"
	E4ListenerBuff key.Modifier = "trailblazerimaginary-e4listener"
	E4Buff         key.Modifier = "trailblazerimaginary-e4buff"
)

func init() {
	modifier.Register(BackupDancer, modifier.Config{
		Stacking: modifier.Replace,
		Listeners: modifier.Listeners{
			OnBeforeHit: superBreakListener,
		},
		StatusType: model.StatusType_STATUS_BUFF,
	})
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

var vulner = []float64{0, 0.6, 0.5, 0.4, 0.3, 0.2}

func superBreakListener(mod *modifier.Instance, e event.HitStart) {
	sourceInfo, _ := mod.Engine().CharacterInfo(mod.Source())
	attackerInfo, _ := mod.Engine().CharacterInfo(e.Attacker)
	attackerStats := mod.Engine().Stats(e.Attacker)

	damage := baseBreakDamage[attackerInfo.Level] * (1 + attackerStats.BreakEffect()) * e.Hit.StanceDamage

	enemyCount := len(mod.Engine().Enemies())
	if enemyCount > 5 {
		enemyCount = 5
	}
	if sourceInfo.Traces["101"] {
		damage *= (1 + vulner[enemyCount])
	}

	if mod.Engine().Stance(e.Defender) == 0 {
		mod.Engine().Attack(info.Attack{
			Key:          SuperBreak,
			Source:       e.Attacker,
			Targets:      []key.TargetID{e.Defender},
			DamageType:   e.Hit.DamageType,
			AttackType:   model.AttackType_ELEMENT_DAMAGE,
			BaseDamage:   info.DamageMap{},
			DamageValue:  damage,
			AsPureDamage: true,
		})
	}
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
