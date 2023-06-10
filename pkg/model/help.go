package model

// returns true if the AttackType is part of an attack and not a biproduct AttackType
func (t AttackType) IsQualified() bool {
	return t != AttackType_DOT && t != AttackType_PURSUED && t != AttackType_ELEMENT_DAMAGE
}
