package testrelic

import (
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

func MusketeerStatlessSet() []*model.Relic {
	return ComposeRelicSet(MusketeerStatless(), MusketeerStatless(), MusketeerStatless(), MusketeerStatless())
}

func MusketeerStatless() *model.Relic {
	// A statless relic for testing 2/4-set effects
	return &model.Relic{
		Key: key.MusketeerOfWildWheat.String(),
		MainStat: &model.RelicStat{
			Stat:   model.Property_HP_FLAT,
			Amount: 0,
		},
		SubStats: nil,
	}
}
