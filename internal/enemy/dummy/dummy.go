package dummy

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/target/enemy"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

func init() {
	enemy.Register(key.DummyEnemy, enemy.Config{
		Create: NewInstance,
		Rank:   model.EnemyRank_ELITE,
		Curve:  enemy.Curve1,
		Base: enemy.BaseStats{
			ATK:        18,
			DEF:        210,
			HP:         150,
			SPD:        100,
			EffectRES:  0.1,
			MinFatigue: 0.2,
			CritDMG:    0.2,
			CritChance: 0.0,
			Stance:     300,
		},
	})
}

type Attack int

const (
	None Attack = iota
	SingleAttack
	BounceAttack
	BlastAttack
	AoeAttack
)

type impl struct {
	engine        engine.Engine
	id            key.TargetID
	info          info.Enemy
	attack        Attack
	hitCount      int
	damagePercent float64
	damageType    model.DamageType
	energyGain    float64
}

func NewInstance(engine engine.Engine, id key.TargetID, enemyInfo info.Enemy) info.EnemyInstance {
	var attack Attack
	switch enemyInfo.Parameters["attack"].GetStringValue() {
	case "", "NONE":
		attack = None
	case "SINGLE":
		attack = SingleAttack
	case "BOUNCE":
		attack = BounceAttack
	case "BLAST":
		attack = BlastAttack
	case "AOE":
		attack = AoeAttack
	default:
		panic(
			"Unknown Attack value in dummy parameters: " + enemyInfo.Parameters["attack"].GetStringValue())
	}

	hit := enemyInfo.Parameters["hit_count"].GetNumberValue()
	if hit <= 0 {
		hit = 1
	}

	energy := enemyInfo.Parameters["energy"].GetNumberValue()
	if energy < 0 {
		energy = 0
	}

	dmg := enemyInfo.Parameters["damage_percent"].GetNumberValue()
	if dmg < 0 {
		dmg = 0
	}

	var dmgType model.DamageType
	switch enemyInfo.Parameters["damage_type"].GetStringValue() {
	case "", "PHYSICAL":
		dmgType = model.DamageType_PHYSICAL
	case "FIRE":
		dmgType = model.DamageType_FIRE
	case "ICE":
		dmgType = model.DamageType_ICE
	case "THUNDER":
		dmgType = model.DamageType_THUNDER
	case "WIND":
		dmgType = model.DamageType_WIND
	case "QUANTUM":
		dmgType = model.DamageType_QUANTUM
	case "IMAGINARY":
		dmgType = model.DamageType_IMAGINARY
	default:
		panic(
			"Unknown DamageType value in dummy parameters: " + enemyInfo.Parameters["damage_type"].GetStringValue())
	}

	return &impl{
		engine:        engine,
		id:            id,
		info:          enemyInfo,
		attack:        attack,
		hitCount:      int(hit),
		damagePercent: dmg,
		damageType:    dmgType,
		energyGain:    energy,
	}
}
