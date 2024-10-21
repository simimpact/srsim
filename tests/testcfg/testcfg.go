package testcfg

import (
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

//nolint:exhaustruct // unused options ok left uninitialized
func TestConfigTwoElites() *model.SimConfig {
	return &model.SimConfig{
		Settings: &model.SimulatorSettings{
			CycleLimit: 10,
		},
		Characters: []*model.Character{},
		Enemies:    []*model.Enemy{StandardEnemy(), StandardEnemy()},
	}
}

//nolint:exhaustruct // unused options ok left uninitialized
func StandardEnemy() *model.Enemy {
	return &model.Enemy{
		Key:   key.DummyEnemy.String(),
		Level: 10,
	}
}
