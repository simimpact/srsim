// package turn provides a naiva implementation of the TurnManager
// TODO: this is a basic place holder implementation
package turn

import (
	"math"

	"github.com/simimpact/srsim/pkg/key"
)

type statService interface {
	Speed(key.TargetID) float64
}

type target struct {
	id    key.TargetID
	gauge float64
}

type TurnCtrl struct {
	targets []*target //maybe should be a linked list here or w/e idk
	stats   statService
}

func New(ids []key.TargetID) *TurnCtrl {
	t := &TurnCtrl{}
	for _, v := range ids {
		t.targets = append(t.targets, &target{
			id: v,
		})
	}

	return t
}

func (c *TurnCtrl) findTargetPos(id key.TargetID) int {
	for i, v := range c.targets {
		if v.id == id {
			return i
		}
	}
	panic("wtf")
}

func (c *TurnCtrl) AdvanceTargetTurn(id key.TargetID) {
	i := c.findTargetPos(id)
	//some logic here to move it to the front
	top := c.targets[i]
	c.targets = append(c.targets[:i], c.targets[i+1:]...)
	c.targets = append([]*target{top}, c.targets...)
}

func (c *TurnCtrl) AdvanceTargetAVPercent(id key.TargetID, per float64) {
	idx := c.findTargetPos(id)
	c.targets[idx].gauge -= 10000 * per
}

func (c *TurnCtrl) DelayTargetAVPercent(id key.TargetID, per float64) {
	idx := c.findTargetPos(id)
	c.targets[idx].gauge += 10000 * per
}

func (c *TurnCtrl) whoNextIndex() int {
	idx := -1
	lowest := math.MaxFloat64
	for i, v := range c.targets {
		av := v.gauge / c.stats.Speed(v.id)
		//< here should only pick up the first lowest av in the list
		if av < lowest {
			lowest = av
			idx = i
		}
	}
	if idx == -1 {
		panic("wtf")
	}
	return idx
}

func (c *TurnCtrl) reduceGaugeWithAV(av float64) {
	for i, v := range c.targets {
		c.targets[i].gauge -= av * c.stats.Speed(v.id)
	}
}

func (c *TurnCtrl) resetGauge(i int) {
	c.targets[i].gauge = 10000
}

func (c *TurnCtrl) AdvanceTurn() key.TargetID {
	idx := c.whoNextIndex()
	t := c.targets[idx]
	c.reduceGaugeWithAV(t.gauge / c.stats.Speed(t.id))
	c.resetGauge(idx)
	return t.id
}

func (c *TurnCtrl) CurrentCycle() int {
	//TODO: ????
	return 0
}