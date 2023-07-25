package enemy

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

func (mgr *Manager) AddEnemy(id key.TargetID, enemy *model.Enemy) error {
	lvl := int(enemy.Level)

	// TODO: placeholder. should generate curve from dm (leaving to whomever implements enemy)
	baseStats := info.NewPropMap()
	baseStats.Modify(prop.HPBase, enemy.Hp)
	baseStats.Modify(prop.DEFBase, 200.0+10.0*float64(lvl))

	debuffRES := info.NewDebuffRESMap()
	for _, res := range enemy.DebuffRes {
		debuffRES.Modify(res.Flag, res.Amount)
	}

	weakness := info.NewWeaknessMap()
	for _, w := range enemy.Weaknesses {
		weakness.Add(w)
	}

	// add 20% res to any type we are not weak to
	for i := 1; i < len(model.DamageType_name); i++ {
		t := model.DamageType(i)
		if weakness.Has(t) {
			continue
		}

		switch t {
		case model.DamageType_PHYSICAL:
			baseStats.Modify(prop.PhysicalDamageRES, 0.2)
		case model.DamageType_FIRE:
			baseStats.Modify(prop.FireDamageRES, 0.2)
		case model.DamageType_ICE:
			baseStats.Modify(prop.IceDamageRES, 0.2)
		case model.DamageType_THUNDER:
			baseStats.Modify(prop.ThunderDamageRES, 0.2)
		case model.DamageType_WIND:
			baseStats.Modify(prop.WindDamageRES, 0.2)
		case model.DamageType_QUANTUM:
			baseStats.Modify(prop.QuantumDamageRES, 0.2)
		case model.DamageType_IMAGINARY:
			baseStats.Modify(prop.ImaginaryDamageRES, 0.2)
		}
	}

	mgr.attr.AddTarget(id, info.Attributes{
		Level:         lvl,
		Stance:        enemy.Toughness,
		MaxStance:     enemy.Toughness,
		BaseStats:     baseStats,
		BaseDebuffRES: debuffRES,
		Weakness:      weakness,
		HPRatio:       1.0,
		Energy:        0,
		MaxEnergy:     0,
	})

	info := info.Enemy{
		Level:     lvl,
		MaxStance: enemy.Toughness,
		Weakness:  weakness,
		DebuffRES: debuffRES,
	}
	mgr.info[id] = info
	return nil
}
