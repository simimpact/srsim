package target

import "github.com/simimpact/srsim/pkg/model"

type ExecutableAction struct {
	Execute    func()
	SPDelta    int
	AttackType model.AttackType
	IsInsert   bool
}

type ExecutableUlt struct {
	Execute func()
}
