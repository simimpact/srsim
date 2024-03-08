package info

import (
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
	"google.golang.org/protobuf/types/known/structpb"
)

type Enemy struct {
	Key        key.Enemy                  `json:"key"`
	Level      int                        `json:"level"`
	Rank       model.EnemyRank            `json:"rank"`
	Parameters map[string]*structpb.Value `json:"parameters"`
}
