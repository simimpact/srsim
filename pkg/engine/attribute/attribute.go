// package attribute provides attribute service which is used to keep track of
// character stats
package attribute

import (
	"fmt"

	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

type Service struct {
	targetBaseStats map[key.TargetID][]float64
}

//TargetBaseAttrFunc shold return the base stats given a targetID
type TargetBaseAttrFunc func(key.TargetID) []float64

func New(targets []key.TargetID, baseFn TargetBaseAttrFunc) (*Service, error) {
	s := &Service{
		targetBaseStats: make(map[key.TargetID][]float64),
	}

	for _, v := range targets {
		base := baseFn(v)
		if base == nil {
			return nil, fmt.Errorf("baseFn did not provide base stats for target id %v", v)
		}
		//sanity check?
		if len(base) != len(model.AttributeType_value) {
			return nil, fmt.Errorf("baseFn provided an attribute slice with invalid len; expecting %v, got %v", len(model.AttributeType_value), len(base))
		}

		s.targetBaseStats[v] = base
	}


	return s, nil
}


func (s *Service) Stat(id key.TargetID, stat model.AttributeType) float64 {
	v, ok := s.targetBaseStats[id]
	if !ok {
		return 0
	}
	//TODO: do we need this sanity check here?
	if _, ok := model.AttributeType_name[int32(stat)]; !ok {
		return 0
	}
	return v[stat]
}