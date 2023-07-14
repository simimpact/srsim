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
	if mgr.HasShield(shield.Target, id) == true {
		mgr.RemoveShield(id, shield.Target)
	}

	// 2. Get the stats for the source and target
	// 3. Compute shield HP/create ShieldInstance given the add params
	ShieldHP := 0.0

	for k, v := range shield.BaseShield {
		switch k {
		case model.ShieldFormula_SHIELD_BY_SHIELDER_ATK:
			ShieldHP += v * mgr.attr.Stats(shield.Source).ATK()
		case model.ShieldFormula_SHIELD_BY_SHIELDER_DEF:
			ShieldHP += v * mgr.attr.Stats(shield.Source).DEF()
		case model.ShieldFormula_SHIELD_BY_SHIELDER_MAX_HP:
			ShieldHP += v * mgr.attr.Stats(shield.Source).HP()
		case model.ShieldFormula_SHIELD_BY_TARGET_MAX_HP:
			ShieldHP += v * mgr.attr.Stats(shield.Target).HP()
		case model.ShieldFormula_SHIELD_BY_SHIELDER_TOTAL_SHIELD:
			ShieldHP += shield.ShieldValue
		}
	}

	newInstance := &Instance{name: id, HP: ShieldHP}

	// 4. add shield to mgr.targets[shield.target]
	mgr.targets[shield.Target] = append(mgr.targets[shield.Target], newInstance)

	// 5. emit ShieldAdded event

	// emit to signify shield added
	mgr.event.ShieldAdded.Emit(event.ShieldAdded{
		ID:           id,
		Info:         shield,
		ShieldHealth: ShieldHP,
	})
}
