package blade

import (
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	E2 key.Modifier = "blade-e2"
	E4 key.Modifier = "blade-e4"
)

func init() {
	modifier.Register(E2, modifier.Config{
		Stacking:   modifier.Replace,
		Duration:   3,
		StatusType: model.StatusType_STATUS_BUFF,
	})

	modifier.Register(E4, modifier.Config{
		Stacking:          modifier.Replace,
		MaxCount:          2,
		CountAddWhenStack: 1,
		StatusType:        model.StatusType_STATUS_BUFF,
		Listeners: modifier.Listeners{
			OnAdd: func(mod *modifier.Instance) {
				mod.SetProperty(prop.HPPercent, 0.2*mod.Count())
			},
		},
	})
}
