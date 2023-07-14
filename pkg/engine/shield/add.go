package shield

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

func (mgr *Manager) AddShield(id key.Shield, shield info.Shield) {
	// TODO: LOGIC FOR ADDING A SHIELD
	// 1. Check if the target already has this shield, if so remove old (mgr.RemoveShield)
	if mgr.HasShield(shield.Target, id) {
		mgr.RemoveShield(id, shield.Target)
	}

	// 2. Get the stats for the source and target
	// 3. Compute shield HP/create ShieldInstance given the add params
	shieldHP := 0.0

	for k, v := range shield.BaseShield {
		switch k {
		case model.ShieldFormula_SHIELD_BY_SHIELDER_ATK:
			shieldHP += v * mgr.attr.Stats(shield.Source).ATK()
		case model.ShieldFormula_SHIELD_BY_SHIELDER_DEF:
			shieldHP += v * mgr.attr.Stats(shield.Source).DEF()
		case model.ShieldFormula_SHIELD_BY_SHIELDER_MAX_HP:
			shieldHP += v * mgr.attr.Stats(shield.Source).HP()
		case model.ShieldFormula_SHIELD_BY_TARGET_MAX_HP:
			shieldHP += v * mgr.attr.Stats(shield.Target).HP()
		case model.ShieldFormula_SHIELD_BY_SHIELDER_TOTAL_SHIELD:
			shieldHP += shield.ShieldValue
		}
	}

	newInstance := &Instance{name: id, HP: shieldHP}

	// 4. add shield to mgr.targets[shield.target]
	mgr.targets[shield.Target] = append(mgr.targets[shield.Target], newInstance)

	// 5. emit ShieldAdded event

	// emit to signify shield added
	mgr.event.ShieldAdded.Emit(event.ShieldAdded{
		ID:           id,
		Info:         shield,
		ShieldHealth: shieldHP,
	})
}
