package enemy

import (
	"fmt"

	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
	"google.golang.org/protobuf/types/known/structpb"
)

func (mgr *Manager) AddEnemy(id key.TargetID, enemy *model.Enemy) error {
	config, ok := enemyCatalog[key.Enemy(enemy.Key)]
	if !ok {
		return fmt.Errorf("invalid enemy: %v", enemy.Key)
	}

	lvl := int(enemy.Level)
	if lvl <= 0 {
		lvl = 1
	}

	baseStats := newBaseStats(&config, enemy, lvl)

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

	stance := config.Base.Stance
	if enemy.BaseStats != nil && enemy.BaseStats.Stance > 0 {
		stance = enemy.BaseStats.Stance
	}
	stance *= Curve(config.Curve)[lvl].StanceScaling

	mgr.attr.AddTarget(id, info.Attributes{
		Level:         lvl,
		Stance:        stance,
		MaxStance:     stance,
		BaseStats:     baseStats,
		BaseDebuffRES: debuffRES,
		Weakness:      weakness,
		HPRatio:       1.0,
		Energy:        0,
		MaxEnergy:     0,
	})

	rank := config.Rank
	if enemy.Rank != model.EnemyRank_RANK_INVALID {
		rank = enemy.Rank
	}

	params := enemy.Parameters.GetFields()
	if params == nil {
		params = make(map[string]*structpb.Value)
	}

	info := info.Enemy{
		Key:        key.Enemy(enemy.Key),
		Level:      lvl,
		Rank:       rank,
		Parameters: params,
	}

	mgr.info[id] = info
	mgr.instances[id] = config.Create(mgr.engine, id, info)
	return nil
}

func newBaseStats(config *Config, enemy *model.Enemy, lvl int) info.PropMap {
	out := info.NewPropMap()
	lc := Curve(config.Curve)[lvl]

	atk := config.Base.ATK
	if enemy.BaseStats != nil && enemy.BaseStats.Atk > 0 {
		atk = enemy.BaseStats.Atk
	}
	out.Modify(prop.ATKBase, atk*lc.ATKScaling)

	def := config.Base.DEF
	if enemy.BaseStats != nil && enemy.BaseStats.Def > 0 {
		def = enemy.BaseStats.Def
	}
	out.Modify(prop.DEFBase, def*lc.DEFScaling)

	hp := config.Base.HP
	if enemy.BaseStats != nil && enemy.BaseStats.Hp > 0 {
		hp = enemy.BaseStats.Hp
	}
	out.Modify(prop.HPBase, hp*lc.HPScaling)

	spd := config.Base.SPD
	if enemy.BaseStats != nil && enemy.BaseStats.Spd > 0 {
		spd = enemy.BaseStats.Spd
	}
	out.Modify(prop.SPDBase, spd*lc.SPDScaling)

	er := config.Base.EffectRES
	if enemy.BaseStats != nil && enemy.BaseStats.EffectRes > 0 {
		er = enemy.BaseStats.EffectRes
	}
	out.Modify(prop.EffectRES, er+lc.EffectRES)
	out.Modify(prop.EffectHitRate, lc.EffectHitRate)

	cc := config.Base.CritChance
	if enemy.BaseStats != nil && enemy.BaseStats.CritChance > 0 {
		cc = enemy.BaseStats.CritChance
	}
	out.Modify(prop.CritChance, cc)

	cd := config.Base.CritDMG
	if enemy.BaseStats != nil && enemy.BaseStats.CritDmg > 0 {
		cd = enemy.BaseStats.CritDmg
	}
	out.Modify(prop.CritDMG, cd)

	fatigue := config.Base.MinFatigue
	if enemy.BaseStats != nil && enemy.BaseStats.MinFatigue > 0 {
		fatigue = enemy.BaseStats.MinFatigue
	}
	out.Modify(prop.MinFatigue, fatigue)

	return out
}
