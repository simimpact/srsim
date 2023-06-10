package target

import "github.com/simimpact/srsim/pkg/model"

type ExecutableAction struct {
	Execute     func()
	SPChange    int
	SkillEffect model.SkillEffect
	AttackType  model.AttackType
	IsInsert    bool
}

type ExecutableUlt struct {
	Execute     func()
	SkillEffect model.SkillEffect
}
