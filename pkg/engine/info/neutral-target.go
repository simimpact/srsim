package info

import (
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

type NeutralTarget struct {
	Key     key.NeutralTarget `json:"key"`
	Element model.DamageType  `json:"element"`
}
