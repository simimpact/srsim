package combat

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
)

func (mgr *Manager) Attack(atk info.Attack) {
	if len(atk.Targets) == 0 || mgr.attr.HPRatio(atk.Source) == 0 {
		return
	}

	// start an attack
	if !mgr.isInAttack && atk.AttackType.IsQualified() {
		mgr.isInAttack = true
		mgr.attackInfo = attackInfo{
			attacker:   atk.Source,
			targets:    atk.Targets,
			attackType: atk.AttackType,
			damageType: atk.DamageType,
		}

		mgr.event.AttackStart.Emit(event.AttackStartEvent{
			Attacker:   atk.Source,
			AttackType: atk.AttackType,
			Targets:    atk.Targets,
			DamageType: atk.DamageType,
		})
	}

	for _, target := range atk.Targets {
		mgr.performHit(mgr.newHit(target, atk))
	}
}

func (mgr *Manager) EndAttack() {
	if mgr.isInAttack {
		mgr.isInAttack = false
		mgr.event.AttackEnd.Emit(event.AttackEndEvent{
			Attacker:   mgr.attackInfo.attacker,
			Targets:    mgr.attackInfo.targets,
			AttackType: mgr.attackInfo.attackType,
			DamageType: mgr.attackInfo.damageType,
		})
	}
}
