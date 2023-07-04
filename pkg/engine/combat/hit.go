package combat

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
)

// PERFORMHIT
// Updates states of all involved parties based upon given hit info.
// 1. Calculates damage to be dealt based upon given hit info
// 2. If shield is present, absorb as much damage as possible using shield.AbsorbDamage
// 3. If shield is not present/cannot take all damage, apply damage to target HP
// 4. Modify stance and energy based upon hit info
// 5. Emit DamageResultEvent (includes modified state and logging variables)

func (mgr *Manager) performHit(hit *info.Hit) {
	mgr.event.HitStart.Emit(event.HitStart{
		Attacker: hit.Attacker.ID(),
		Defender: hit.Defender.ID(),
		Hit:      hit,
	})

	crit := crit(hit, mgr.rdm)

	base := baseDamage(hit)*hit.HitRatio + hit.DamageValue
	bonus := bonusDamage(hit)
	defMult := defMult(hit)
	res := res(hit)
	vul := vul(hit)
	toughnessMultiplier := toughness(hit)
	fatigue := 1 - hit.Attacker.GetProperty(prop.Fatigue)
	allDamageReduce := damageReduce(hit)
	critDmg := critDmg(hit, crit)

	total := base * bonus * defMult * res * vul * toughnessMultiplier * fatigue * allDamageReduce * critDmg

	hpUpdate := mgr.shld.AbsorbDamage(hit.Defender.ID(), total)
	mgr.attr.ModifyHPByAmount(hit.Defender.ID(), hit.Attacker.ID(), -total, true)
	mgr.attr.ModifyStance(hit.Defender.ID(), hit.Attacker.ID(), -hit.StanceDamage*hit.HitRatio)
	if mgr.target.IsCharacter(hit.Attacker.ID()) {
		mgr.attr.ModifyEnergy(hit.Attacker.ID(), hit.EnergyGain*hit.HitRatio)
	} else {
		mgr.attr.ModifyEnergy(hit.Defender.ID(), hit.EnergyGain*hit.HitRatio)
	}

	mgr.event.HitEnd.Emit(event.HitEnd{
		Attacker:            hit.Attacker.ID(),
		Defender:            hit.Defender.ID(),
		AttackType:          hit.AttackType,
		DamageType:          hit.DamageType,
		TotalDamage:         total,
		ShieldDamage:        total - hpUpdate,
		HPDamage:            hpUpdate,
		HPRatioRemaining:    mgr.attr.HPRatio(hit.Defender.ID()),
		BaseDamage:          base,
		BonusDamage:         bonus,
		DefenceMultiplier:   defMult,
		Resistance:          res,
		Vulnerability:       vul,
		ToughnessMultiplier: toughnessMultiplier,
		Fatigue:             fatigue,
		AllDamageReduce:     allDamageReduce,
		CritDamage:          critDmg,
		IsCrit:              crit,
		UseSnapshot:         hit.UseSnapshot,
	})
}

func (mgr *Manager) newHit(target key.TargetID, atk info.Attack) *info.Hit {
	// set HitRatio to 1 if unspecified
	ratio := atk.HitRatio
	if ratio <= 0 {
		ratio = 1
	}

	// make a copy of the base damage info
	baseDamage := make(info.DamageMap, len(atk.BaseDamage))
	for k, v := range atk.BaseDamage {
		baseDamage[k] = v
	}

	return &info.Hit{
		Attacker:     mgr.attr.Stats(atk.Source),
		Defender:     mgr.attr.Stats(target),
		AttackType:   atk.AttackType,
		DamageType:   atk.DamageType,
		BaseDamage:   baseDamage,
		EnergyGain:   atk.EnergyGain,
		StanceDamage: atk.StanceDamage,
		HitRatio:     ratio,
		AsPureDamage: atk.AsPureDamage,
		DamageValue:  atk.DamageValue,
		UseSnapshot:  atk.UseSnapshot,
	}
}
