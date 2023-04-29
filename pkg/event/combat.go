package event

type combat struct {
}

type CombatEvent struct {
}

func (c *combat) EmitCombatEvent(ev *CombatEvent) {

}

func (c *combat) SubCombatEvent(key string, f func(ev *CombatEvent)) {

}
