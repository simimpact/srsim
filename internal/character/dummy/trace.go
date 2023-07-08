package dummy

import (
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/engine/target/character"
)

var traces = character.TraceMap{
	"101": {
		Ascension: 2,
	},
	"102": {
		Ascension: 4,
	},
	"103": {
		Ascension: 6,
	},
	"201": {
		Stat:   prop.ATKPercent,
		Amount: 0.04,
		Level:  1,
	},
	"202": {
		Stat:      prop.QuantumDamagePercent,
		Amount:    0.032,
		Ascension: 2,
	},
	"203": {
		Stat:      prop.ATKPercent,
		Amount:    0.0400,
		Ascension: 3,
	},
	"204": {
		Stat:      prop.DEFPercent,
		Amount:    0.0500,
		Ascension: 3,
	},
	"205": {
		Stat:      prop.ATKPercent,
		Amount:    0.0600,
		Ascension: 5,
	},
	"206": {
		Stat:      prop.QuantumDamagePercent,
		Amount:    0.0480,
		Ascension: 5,
	},
	"207": {
		Stat:      prop.ATKPercent,
		Amount:    0.0600,
		Ascension: 5,
	},
	"208": {
		Stat:      prop.DEFPercent,
		Amount:    0.0750,
		Ascension: 6,
	},
	"209": {
		Stat:   prop.QuantumDamagePercent,
		Amount: 0.0640,
		Level:  75,
	},
	"210": {
		Stat:   prop.ATKPercent,
		Amount: 0.0800,
		Level:  80,
	},
}

func (c *char) a2() {
	if !c.info.Traces["101"] {
		return
	}
}

func (c *char) a4() {
	if !c.info.Traces["102"] {
		return
	}
}

func (c *char) a6() {
	if !c.info.Traces["103"] {
		return
	}
}
