package asta

import "github.com/simimpact/srsim/pkg/engine/modifier"

const (
	asta_talent = "asta-talent"
)

func init() {
	modifier.Register(asta_talent, modifier.Config{
		Listeners: modifier.Listeners{
			OnPhase1: phase1TalentListener,
		},
	})
}

func (c *char) initTalent() {

}

func phase1TalentListener(mod *modifier.Instance) {

}
