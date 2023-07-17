package shield

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

func (mgr *Manager) AddShield(id key.Shield, shield info.Shield) {

	// Get the stats for the source and target
	source := mgr.attr.Stats(shield.Source)
	maxShield := mgr.targets[shield.Target][mgr.MaxShield(shield.Target)]
	target := mgr.attr.Stats(shield.Target)

	// Compute shield baseHP from ShieldMap property values
	baseHP := 0.0

	for k, v := range shield.BaseShield {
		switch k {
		case model.ShieldFormula_SHIELD_BY_SHIELDER_ATK:
			baseHP += v * source.ATK()
		case model.ShieldFormula_SHIELD_BY_SHIELDER_DEF:
			baseHP += v * source.DEF()
		case model.ShieldFormula_SHIELD_BY_SHIELDER_MAX_HP:
			baseHP += v * source.HP()
		case model.ShieldFormula_SHIELD_BY_TARGET_MAX_HP:
			baseHP += v * target.HP()
		case model.ShieldFormula_SHIELD_BY_SHIELDER_TOTAL_SHIELD:
			baseHP += v * maxShield.hp
		}
	}

	// Compute final shieldHP using shield HP formula
	shieldHP := baseHP * (1 + source.GetProperty(prop.ShieldBoost)) * (1 + target.GetProperty(prop.ShieldTaken))

	// Create new instance to add to list of shields for target
	newInstance := &Instance{name: id, hp: shieldHP}

	switch isMatching, index := mgr.CheckMatching(id, shield); isMatching {
	// Replace shield at matching index
	case isMatching == true:
		mgr.targets[shield.Target][index] = newInstance
	// Add new shield to targets list
	case isMatching == false:
		mgr.targets[shield.Target] = append(mgr.targets[shield.Target], newInstance)
	}

	// Event emission
	mgr.event.ShieldAdded.Emit(event.ShieldAdded{
		ID:           id,
		Info:         shield,
		ShieldHealth: baseHP,
	})
}

// Check list of shields for a matching shield
func (mgr *Manager) CheckMatching(id key.Shield, shield info.Shield) (bool, int) {
	for i, shields := range mgr.targets[shield.Target] {
		if shields.name == id {
			return true, i
		}
	}
	return false, 0
}
