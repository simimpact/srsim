package model

// returns true if the AttackType is part of an attack and not a biproduct AttackType
func (t AttackType) IsQualified() bool {
	return t != AttackType_DOT && t != AttackType_PURSUED && t != AttackType_ELEMENT_DAMAGE
}

// returns true if the rank of the enemy is elite or higher
func (r EnemyRank) IsElite() bool {
	return r > EnemyRank_MINION
}
