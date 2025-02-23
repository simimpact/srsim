package hanya

import (
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	A2 = "hanya-a2"
	A4 = "hanya-a4"
	A6 = "hanya-a6"
)

func init() {
	modifier.Register(A2, modifier.Config{
		StatusType: model.StatusType_STATUS_BUFF,
		Stacking:   modifier.ReplaceBySource,
	})
}
