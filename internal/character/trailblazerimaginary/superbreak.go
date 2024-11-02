package trailblazerimaginary

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	BackupDancer         key.Modifier = "trailblazerimaginary-backupdancer"
	SuperBreak           key.Attack   = "trailblazerimaginary-superbreak"
	StanceDamageRecorder key.Modifier = "trailblazerimaginary-stancedamagerecorder"
)

type recorderData struct {
	stanceDamage float64
}

func init() {
	modifier.Register(BackupDancer, modifier.Config{
		Stacking: modifier.Replace,
		Listeners: modifier.Listeners{
			OnAfterAttack:  triggerSuperBreak,
			OnBeforeAttack: startStanceDamageRecord,
		},
		StatusType: model.StatusType_STATUS_BUFF,
	})
	modifier.Register(StanceDamageRecorder, modifier.Config{
		StatusType: model.StatusType_UNKNOWN_STATUS,
		Stacking:   modifier.ReplaceBySource,
		Listeners: modifier.Listeners{
			OnBeforeBeingHitAll: stanceDamageListener,
		},
	})
}

func startStanceDamageRecord(mod *modifier.Instance, e event.AttackStart) {
	for _, target := range mod.Engine().Enemies() {
		mod.Engine().AddModifier(target, info.Modifier{
			Name:   StanceDamageRecorder,
			Source: mod.Owner(),
			State: &recorderData{
				stanceDamage: 0.0,
			},
		})
	}
}

func stanceDamageListener(mod *modifier.Instance, e event.HitStart) {
	if e.Attacker != mod.Source() {
		return
	}
	if mod.Engine().Stance(e.Defender) == 0 {
		mod.State().(*recorderData).stanceDamage += e.Hit.StanceDamage
	}
}

func getModifierFromSource(engine engine.Engine, target, source key.TargetID, modifier key.Modifier) []info.Modifier {
	modifierList := engine.GetModifiers(target, modifier)
	resList := []info.Modifier{}
	for _, mod := range modifierList {
		if mod.Source == source {
			resList = append(resList, mod)
		}
	}
	return resList
}

var vulner = []float64{0, 0.6, 0.5, 0.4, 0.3, 0.2}

func triggerSuperBreak(mod *modifier.Instance, e event.AttackEnd) {
	for _, target := range mod.Engine().Enemies() {
		recorder := getModifierFromSource(mod.Engine(), target, mod.Owner(), StanceDamageRecorder)[0]
		stanceDamage := recorder.State.(*recorderData).stanceDamage
		mod.Engine().RemoveModifierFromSource(target, mod.Owner(), StanceDamageRecorder)

		if stanceDamage == 0 {
			continue
		}

		sourceInfo, _ := mod.Engine().CharacterInfo(mod.Source())
		attackerInfo, _ := mod.Engine().CharacterInfo(mod.Owner())
		attackerStats := mod.Engine().Stats(e.Attacker)

		damage := baseBreakDamage[attackerInfo.Level] * (1 + attackerStats.BreakEffect()) * stanceDamage

		enemyCount := len(mod.Engine().Enemies())
		if enemyCount > 5 {
			enemyCount = 5
		}
		if sourceInfo.Traces["101"] {
			damage *= (1 + vulner[enemyCount])
		}

		mod.Engine().Attack(info.Attack{
			Key:          SuperBreak,
			Source:       mod.Owner(),
			Targets:      []key.TargetID{target},
			DamageType:   attackerInfo.Element,
			AttackType:   model.AttackType_ELEMENT_DAMAGE,
			BaseDamage:   info.DamageMap{},
			DamageValue:  damage,
			AsPureDamage: true,
			UseSnapshot:  true,
		})
	}
}
