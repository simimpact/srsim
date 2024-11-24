package prop

// GameToProp is a helper for use in scripts and reference when implementing
var GameToProp = map[string]Property{
	"Unknow":                    Invalid,
	"BaseHP":                    HPBase,
	"HPAddedRatio":              HPPercent,
	"HPDelta":                   HPFlat,
	"HPConvert":                 HPConvert,
	"BaseAttack":                ATKBase,
	"AttackAddedRatio":          ATKPercent,
	"AttackDelta":               ATKFlat,
	"AttackConvert":             ATKConvert,
	"BaseDefence":               DEFBase,
	"DefenceAddedRatio":         DEFPercent,
	"DefenceDelta":              DEFFlat,
	"DefenceConvert":            DEFConvert,
	"BaseSpeed":                 SPDBase,
	"SpeedAddedRatio":           SPDPercent,
	"SpeedDelta":                SPDFlat,
	"SpeedConvert":              SPDConvert,
	"AllDamageTypeAddedRatio":   AllDamagePercent,
	"AllDamageReduce":           AllDamageReduce,
	"DotDamageAddedRatio":       DOTDamagePercent,
	"FatigueRatio":              Fatigue,
	"CriticalChanceBase":        CritChance,
	"CriticalDamageBase":        CritDMG,
	"PhysicalAddedRatio":        PhysicalDamagePercent,
	"FireAddedRatio":            FireDamagePercent,
	"IceAddedRatio":             IceDamagePercent,
	"ThunderAddedRatio":         ThunderDamagePercent,
	"QuantumAddedRatio":         QuantumDamagePercent,
	"ImaginaryAddedRatio":       ImaginaryDamagePercent,
	"WindAddedRatio":            WindDamagePercent,
	"PhysicalResistanceDelta":   PhysicalDamageRES,
	"FireResistanceDelta":       FireDamageRES,
	"IceResistanceDelta":        IceDamageRES,
	"ThunderResistanceDelta":    ThunderDamageRES,
	"QuantumResistanceDelta":    QuantumDamageRES,
	"ImaginaryResistanceDelta":  ImaginaryDamageRES,
	"WindResistanceDelta":       WindDamageRES,
	"AllDamageTypeResistance":   AllDamageRES,
	"AllDamageTypePenetrate":    AllDamagePEN,
	"PhysicalPenetrate":         PhysicalPEN,
	"FirePenetrate":             FirePEN,
	"IcePenetrate":              IcePEN,
	"ThunderPenetrate":          ThunderPEN,
	"QuantumPenetrate":          QuantumPEN,
	"ImaginaryPenetrate":        ImaginaryPEN,
	"WindPenetrate":             WindPEN,
	"PhysicalTakenRatio":        PhysicalDamageTaken,
	"FireTakenRatio":            FireDamageTaken,
	"IceTakenRatio":             IceDamageTaken,
	"ThunderTakenRatio":         ThunderDamageTaken,
	"QuantumTakenRatio":         QuantumDamageTaken,
	"ImaginaryTakenRatio":       ImaginaryDamageTaken,
	"WindTakenRatio":            WindDamageTaken,
	"AllDamageTypeTakenRatio":   AllDamageTaken,
	"MinimumFatigueRatio":       MinFatigue,
	"StanceBreakAddedRatio":     AllStanceDMGPercent,
	"HealRatioBase":             HealBoost,
	"HealRatioConvert":          HealBoostConvert,
	"HealTakenRatio":            HealTaken,
	"ShieldAddedRatio":          ShieldBoost,
	"ShieldTakenRatio":          ShieldTaken,
	"StatusProbabilityBase":     EffectHitRate,
	"StatusProbabilityConvert":  EffectHitRateConvert,
	"StatusResistanceBase":      EffectRES,
	"StatusResistanceConvert":   EffectRESConvert,
	"SPRatioBase":               EnergyRegen,
	"SPRatioConvert":            EnergyRegenConvert,
	"BreakDamageAddedRatioBase": BreakEffect,
	"AggroBase":                 AggroBase,
	"AggroAddedRatio":           AggroPercent,
	"AggroDelta":                AggroFlat,
}
