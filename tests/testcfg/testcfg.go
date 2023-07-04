package testcfg

import (
	"github.com/simimpact/srsim/pkg/model"
)

func TestConfigTwoElites() *model.SimConfig {
	return &model.SimConfig{
		Settings: &model.SimulatorSettings{
			CycleLimit: 10,
		},
		Characters: []*model.Character{},
		Enemies:    []*model.Enemy{StandardEnemy(), StandardEnemy()},
	}
}

func StandardEnemy() *model.Enemy {
	return &model.Enemy{
		Level:      90,
		Hp:         200000,
		Toughness:  420,
		Weaknesses: nil,
		DebuffRes:  nil,
	}
}
