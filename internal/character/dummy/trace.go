package dummy

import (
	"github.com/simimpact/srsim/pkg/engine/target/character"
	"github.com/simimpact/srsim/pkg/model"
)

var traces = character.TraceMap{
	"1201101": {
		Ascension: 2,
	},
	"1201102": {
		Ascension: 4,
	},
	"1201103": {
		Ascension: 6,
	},
	"1201201": {
		Stat:   model.Property_ATK_PERCENT,
		Amount: 0.04,
		Level:  1,
	},
	"1201202": {
		Stat:      model.Property_QUANTUM_DMG_PERCENT,
		Amount:    0.032,
		Ascension: 2,
	},
	"1201203": {
		Stat:      model.Property_ATK_PERCENT,
		Amount:    0.0400,
		Ascension: 3,
	},
	"1201204": {
		Stat:      model.Property_DEF_PERCENT,
		Amount:    0.0500,
		Ascension: 3,
	},
	"1201205": {
		Stat:      model.Property_ATK_PERCENT,
		Amount:    0.0600,
		Ascension: 5,
	},
	"1201206": {
		Stat:      model.Property_QUANTUM_DMG_PERCENT,
		Amount:    0.0480,
		Ascension: 5,
	},
	"1201207": {
		Stat:      model.Property_ATK_PERCENT,
		Amount:    0.0600,
		Ascension: 5,
	},
	"1201208": {
		Stat:      model.Property_DEF_PERCENT,
		Amount:    0.0750,
		Ascension: 6,
	},
	"1201209": {
		Stat:   model.Property_QUANTUM_DMG_PERCENT,
		Amount: 0.0640,
		Level:  75,
	},
	"1201210": {
		Stat:   model.Property_ATK_PERCENT,
		Amount: 0.0800,
		Level:  80,
	},
}

func (c *char) a2() {
	if !c.info.Traces["1201101"] {
		return
	}
}

func (c *char) a4() {
	if !c.info.Traces["1201102"] {
		return
	}
}

func (c *char) a6() {
	if !c.info.Traces["1201103"] {
		return
	}
}
