package modifier

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

type Instance struct {
	name                     key.Modifier
	owner                    key.TargetID
	source                   key.TargetID
	state                    any
	tickImmediately          bool
	canTickImmediatelyPhase2 bool
	duration                 int
	count                    float64
	maxCount                 float64
	countAddWhenStack        float64
	stats                    info.PropMap
	debuffRES                info.DebuffRESMap
	weakness                 info.WeaknessMap
	renewTurn                int
	manager                  *Manager
	listeners                Listeners
	statusType               model.StatusType
	flags                    []model.BehaviorFlag
	modifySnapshot           bool
	canDispel                bool
}

func (mgr *Manager) newInstance(owner key.TargetID, mod info.Modifier, renew int) *Instance {
	config := modifierCatalog[mod.Name]
	mi := &Instance{
		owner:                    owner,
		name:                     mod.Name,
		source:                   mod.Source,
		state:                    mod.State,
		tickImmediately:          mod.TickImmediately,
		duration:                 mod.Duration,
		count:                    mod.Count,
		maxCount:                 mod.MaxCount,
		countAddWhenStack:        mod.CountAddWhenStack,
		stats:                    mod.Stats,
		debuffRES:                mod.DebuffRES,
		weakness:                 mod.Weakness,
		manager:                  mgr,
		listeners:                config.Listeners,
		statusType:               config.StatusType,
		flags:                    config.BehaviorFlags,
		modifySnapshot:           config.CanModifySnapshot,
		canTickImmediatelyPhase2: false,
		renewTurn:                renew,
		canDispel:                config.CanDispel,
	}

	if mi.stats == nil {
		mi.stats = info.NewPropMap()
	}
	if mi.debuffRES == nil {
		mi.debuffRES = info.NewDebuffRESMap()
	}
	if mi.weakness == nil {
		mi.weakness = info.NewWeaknessMap()
	}

	// Apply defaults from config as fallback
	if mi.countAddWhenStack == 0 {
		mi.countAddWhenStack = config.CountAddWhenStack
	}
	if mi.maxCount == 0 {
		mi.maxCount = config.MaxCount
	}
	if mi.count == 0 {
		mi.count = config.Count
	}
	if mi.duration == 0 {
		mi.duration = config.Duration
	}

	// default "infinite" cases
	if mi.duration <= 0 {
		mi.duration = -1
	}
	if mi.count <= 0 {
		mi.count = -1
	}
	if mi.maxCount <= 0 {
		mi.maxCount = -1
	}

	// cases that can stack. Need to ensure count is stackable pending what was specified on instance
	if config.Stacking == Replace || config.Stacking == ReplaceBySource || config.Stacking == Merge {
		if mi.count <= 0 && mi.countAddWhenStack > 0 {
			mi.count = mi.countAddWhenStack
		}
	}

	return mi
}

// Add a property to this modifier instance. Will cause all modifiers attached to the owner of this
// modifier to execute OnPropertyChange listener
func (mi *Instance) AddProperty(prop prop.Property, amt float64) {
	mi.stats.Modify(prop, amt)
	mi.manager.emitPropertyChange(mi.owner)
}

func (mi *Instance) SetProperty(prop prop.Property, amt float64) {
	old := mi.stats[prop]
	mi.stats.Set(prop, amt)
	if old != mi.stats[prop] {
		mi.manager.emitPropertyChange(mi.owner)
	}
}

// Add a new debuffRES for the given behavior flag
func (mi *Instance) AddDebuffRES(flag model.BehaviorFlag, amt float64) {
	mi.debuffRES.Modify(flag, amt)
}

// Adds a new weakness to this modifier. In stats snapshots, the modifier owner will now be listed
// as weak to this damage type.
func (mi *Instance) AddWeakness(weakness model.DamageType) {
	mi.weakness[weakness] = true
}

// Removes the given weakness from the modifier's weakness list. NOTE: This does not remove
// the weakness from the target if it is applied by another modifier or is in the modifier's base
// stats.
func (mi *Instance) RemoveWeakness(weakness model.DamageType) {
	delete(mi.weakness, weakness)
}

// Get the current value of a property set by this modifier instance
func (mi *Instance) GetProperty(prop prop.Property) float64 {
	return mi.stats[prop]
}

// Get the current value of a debuff res set by this modifier instance
func (mi *Instance) GetDebuffRES(flags ...model.BehaviorFlag) float64 {
	return mi.debuffRES.GetDebuffRES(flags...)
}

// check if this modifier instance has applied this specific weakness type to the target
func (mi *Instance) HasWeakness(dmgType model.DamageType) bool {
	return mi.weakness[dmgType]
}

// Remove this modifier instance
func (mi *Instance) RemoveSelf() {
	mi.manager.RemoveSelf(mi.owner, mi)
}

// Name of the modifier this instance represents (what modifier config it is associated with)
func (mi *Instance) Name() key.Modifier {
	return mi.name
}

// TargetID for who created this modifier instance
func (mi *Instance) Source() key.TargetID {
	return mi.source
}

// Return the current stats of the target this modifier is attached to
func (mi *Instance) OwnerStats() *info.Stats {
	return mi.Engine().Stats(mi.owner)
}

// TargetID for who this modifier is currently attached to
func (mi *Instance) Owner() key.TargetID {
	return mi.owner
}

// Returns the state struct associated with this modifier instance (created in AddModifier call)
// This state struct is untyped. Up to modifier logic to type assert to the desired struct type
func (mi *Instance) State() any {
	return mi.state
}

// How long before this instance expires. Will be removed when Duration == 0. A negative duration
// means  that this modifier will never expire
func (mi *Instance) Duration() int {
	return mi.duration
}

// How many stacks are associated with this modifier instance. If count reaches 0, the modifier
// will be removed from the target. A count of -1 means that a count was never specified and/or
// a StackingBehavior that does not support stacking was used.
func (mi *Instance) Count() float64 {
	return mi.count
}

// The maximum amount of stacks the associated modifier instance can have.
func (mi *Instance) MaxCount() float64 {
	return mi.maxCount
}

// What status type this modifier instance is (copied from the modifier config)
func (mi *Instance) StatusType() model.StatusType {
	return mi.statusType
}

// What BehaviorFlags exist for this modifier instance (copied from the modifier config)
func (mi *Instance) BehaviorFlags() []model.BehaviorFlag {
	return mi.flags
}

// Returns a copy of the config associated with this instance (config created via modifier.Register)
func (mi *Instance) Config() Config {
	return modifierCatalog[mi.name]
}

// Engine access for various operations
func (mi *Instance) Engine() engine.Engine {
	return mi.manager.engine
}

// Convert modifier instance to model version for read-only access to the instance
func (mi *Instance) ToModel() info.Modifier {
	props := info.NewPropMap()
	res := info.NewDebuffRESMap()

	props.AddAll(mi.stats)
	res.AddAll(mi.debuffRES)

	return info.Modifier{
		Name:              mi.name,
		Source:            mi.source,
		State:             mi.state,
		Duration:          mi.duration,
		TickImmediately:   mi.tickImmediately,
		Count:             mi.count,
		MaxCount:          mi.maxCount,
		CountAddWhenStack: mi.countAddWhenStack,
		Stats:             props,
		DebuffRES:         res,
		CanDispel:         mi.canDispel,
	}
}
