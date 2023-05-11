package action

import "github.com/simimpact/srsim/pkg/key"

type Action struct {
	Type   key.ActionType
	target key.TargetID
}