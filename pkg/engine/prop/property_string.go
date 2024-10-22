// Code generated by "stringer -type=Property"; DO NOT EDIT.

package prop

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Invalid-0]
	_ = x[HPBase-1]
	_ = x[HPPercent-2]
	_ = x[HPFlat-3]
	_ = x[HPConvert-4]
	_ = x[ATKBase-5]
	_ = x[ATKPercent-6]
	_ = x[ATKFlat-7]
	_ = x[ATKConvert-8]
	_ = x[DEFBase-9]
	_ = x[DEFPercent-10]
	_ = x[DEFFlat-11]
	_ = x[DEFConvert-12]
	_ = x[SPDBase-13]
	_ = x[SPDPercent-14]
	_ = x[SPDFlat-15]
	_ = x[SPDConvert-16]
	_ = x[CritChance-17]
	_ = x[CritDMG-18]
	_ = x[EnergyRegen-19]
	_ = x[EnergyRegenConvert-20]
	_ = x[EffectHitRate-21]
	_ = x[EffectHitRateConvert-22]
	_ = x[EffectRES-23]
	_ = x[EffectRESConvert-24]
	_ = x[HealBoost-25]
	_ = x[HealBoostConvert-26]
	_ = x[HealTaken-27]
	_ = x[ShieldBoost-28]
	_ = x[ShieldTaken-29]
	_ = x[AggroBase-30]
	_ = x[AggroPercent-31]
	_ = x[AggroFlat-32]
	_ = x[BreakEffect-33]
	_ = x[AllDamageRES-34]
	_ = x[PhysicalDamageRES-35]
	_ = x[FireDamageRES-36]
	_ = x[IceDamageRES-37]
	_ = x[ThunderDamageRES-38]
	_ = x[QuantumDamageRES-39]
	_ = x[ImaginaryDamageRES-40]
	_ = x[WindDamageRES-41]
	_ = x[AllDamagePEN-42]
	_ = x[PhysicalPEN-43]
	_ = x[FirePEN-44]
	_ = x[IcePEN-45]
	_ = x[ThunderPEN-46]
	_ = x[QuantumPEN-47]
	_ = x[ImaginaryPEN-48]
	_ = x[WindPEN-49]
	_ = x[AllDamageTaken-50]
	_ = x[PhysicalDamageTaken-51]
	_ = x[FireDamageTaken-52]
	_ = x[IceDamageTaken-53]
	_ = x[ThunderDamageTaken-54]
	_ = x[QuantumDamageTaken-55]
	_ = x[ImaginaryDamageTaken-56]
	_ = x[WindDamageTaken-57]
	_ = x[AllDamagePercent-58]
	_ = x[DOTDamagePercent-59]
	_ = x[FireDamagePercent-60]
	_ = x[IceDamagePercent-61]
	_ = x[ThunderDamagePercent-62]
	_ = x[QuantumDamagePercent-63]
	_ = x[ImaginaryDamagePercent-64]
	_ = x[WindDamagePercent-65]
	_ = x[PhysicalDamagePercent-66]
	_ = x[AllStanceDMGPercent-67]
	_ = x[AllDamageReduce-90]
	_ = x[Fatigue-91]
	_ = x[MinFatigue-92]
}

const (
	_Property_name_0 = "InvalidHPBaseHPPercentHPFlatHPConvertATKBaseATKPercentATKFlatATKConvertDEFBaseDEFPercentDEFFlatDEFConvertSPDBaseSPDPercentSPDFlatSPDConvertCritChanceCritDMGEnergyRegenEnergyRegenConvertEffectHitRateEffectHitRateConvertEffectRESEffectRESConvertHealBoostHealBoostConvertHealTakenShieldBoostShieldTakenAggroBaseAggroPercentAggroFlatBreakEffectAllDamageRESPhysicalDamageRESFireDamageRESIceDamageRESThunderDamageRESQuantumDamageRESImaginaryDamageRESWindDamageRESAllDamagePENPhysicalPENFirePENIcePENThunderPENQuantumPENImaginaryPENWindPENAllDamageTakenPhysicalDamageTakenFireDamageTakenIceDamageTakenThunderDamageTakenQuantumDamageTakenImaginaryDamageTakenWindDamageTakenAllDamagePercentDOTDamagePercentFireDamagePercentIceDamagePercentThunderDamagePercentQuantumDamagePercentImaginaryDamagePercentWindDamagePercentPhysicalDamagePercentAllStanceDMGPercent"
	_Property_name_1 = "AllDamageReduceFatigueMinFatigue"
)

var (
	_Property_index_0 = [...]uint16{0, 7, 13, 22, 28, 37, 44, 54, 61, 71, 78, 88, 95, 105, 112, 122, 129, 139, 149, 156, 167, 185, 198, 218, 227, 243, 252, 268, 277, 288, 299, 308, 320, 329, 340, 352, 369, 382, 394, 410, 426, 444, 457, 469, 480, 487, 493, 503, 513, 525, 532, 546, 565, 580, 594, 612, 630, 650, 665, 681, 697, 714, 730, 750, 770, 792, 809, 830, 849}
	_Property_index_1 = [...]uint8{0, 15, 22, 32}
)

func (i Property) String() string {
	switch {
	case 0 <= i && i <= 67:
		return _Property_name_0[_Property_index_0[i]:_Property_index_0[i+1]]
	case 90 <= i && i <= 92:
		i -= 90
		return _Property_name_1[_Property_index_1[i]:_Property_index_1[i+1]]
	default:
		return "Property(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
