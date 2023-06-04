package combat

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/model"
)

func (mgr *Manager) Attack(atk info.Attack, effect model.AttackEffect) {
	// start an attack
	if !mgr.isInAttack && isAttackStartable(atk.AttackType) {
		// TODO: make this a struct?
		mgr.isInAttack = true
		mgr.attackInfo = attackInfo{
			attacker:     atk.Source,
			targets:      atk.Targets,
			attackType:   atk.AttackType,
			damageType:   atk.DamageType,
			attackEffect: effect,
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
	mgr.isInAttack = false
	mgr.event.AttackEnd.Emit(event.AttackEndEvent{
		Attacker:     mgr.attackInfo.attacker,
		Targets:      mgr.attackInfo.targets,
		AttackType:   mgr.attackInfo.attackType,
		AttackEffect: mgr.attackInfo.attackEffect,
		DamageType:   mgr.attackInfo.damageType,
	})
}

func isAttackStartable(t model.AttackType) bool {
	return t != model.AttackType_DOT && t != model.AttackType_PURSUED
}
