package combat

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/model"
)

func (mgr *Manager) Attack(atk info.Attack, effect model.AttackEffect) {
	// TODO: validate, dead check

	// start an attack
	if !mgr.isInAttack && isAttackStartable(atk.AttackType) {
		// TODO: make this a struct?
		mgr.isInAttack = true
		mgr.attacker = atk.Source
		mgr.targets = atk.Targets
		mgr.attackType = atk.AttackType
		mgr.damageType = atk.DamageType
		mgr.attackEffect = effect

		mgr.event.AttackStart.Emit(event.AttackStartEvent{
			Attacker:   atk.Source,
			AttackType: atk.AttackType,
			Targets:    mgr.targets,
			DamageType: mgr.damageType,
		})
	}

	for _, target := range atk.Targets {
		mgr.performHit(mgr.newHit(target, atk))
	}
}

func (mgr *Manager) EndAttack() {
	mgr.isInAttack = false
	mgr.event.AttackEnd.Emit(event.AttackEndEvent{
		Attacker:     mgr.attacker,
		Targets:      mgr.targets,
		AttackType:   mgr.attackType,
		AttackEffect: mgr.attackEffect,
		DamageType:   mgr.damageType,
	})
}

func isAttackStartable(t model.AttackType) bool {
	return t != model.AttackType_DOT && t != model.AttackType_PURSUED
}
