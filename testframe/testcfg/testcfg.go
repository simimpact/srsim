package testcfg

import (
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

func TestConfigTwoElites() *model.SimConfig {

	return &model.SimConfig{
		Iterations:  1,
		WorkerCount: 1,
		Settings: &model.SimulatorSettings{
			CycleLimit: 10,
			TtkMode:    false,
		},
		Characters: []*model.Character{DanHung()},
		Enemies:    []*model.Enemy{StandardEnemy(), StandardEnemy()},
	}
}

func DanHung() *model.Character {
	return &model.Character{
		Key:      key.DanHeng.String(),
		Level:    80,
		MaxLevel: 80,
		Eidols:   0,
		Traces:   nil,
		Talents:  []uint32{1, 1, 1, 1},
		Cone: &model.LightCone{
			Key:        key.OnlySilenceRemains.String(),
			Level:      80,
			MaxLevel:   80,
			Imposition: 1,
		},
		Relics:      MusketeerSet(),
		StartEnergy: 50,
	}
}

func MusketeerSet() []*model.Relic {
	return nil
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
