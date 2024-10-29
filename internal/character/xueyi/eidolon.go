package xueyi

import (
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	E4 = "xueyi-e4"
)

func init() {
	modifier.Register(E4, modifier.Config{
		Stacking:   modifier.ReplaceBySource,
		Duration:   2,
		StatusType: model.StatusType_STATUS_BUFF,
	})
}
