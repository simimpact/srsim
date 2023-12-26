package eval

import (
	"fmt"

	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/logic/gcs/ast"
	"github.com/simimpact/srsim/pkg/model"
)

// Functions for writing more flexible scripts.
func (e *Eval) initConditionalFuncs(env *Env) {
	funcs := map[string]func(*ast.CallExpr, *Env) (Obj, error){
		// modifier
		"has_modifier":   e.hasModifier,
		"modifier_count": e.modifierCount,
		// attribute
		"ult_ready":       e.ultReady,
		"skill_points":    e.skillPoints,
		"energy":          e.energy,
		"max_energy":      e.maxEnergy,
		"hp_ratio":        e.hpRatio,
		"weakness_broken": e.weaknessBroken,
		"has_weakness":    e.hasWeakness,
		"stance":          e.stance,
		"max_stance":      e.maxStance,
		"is_alive":        e.isAlive,
		// shield
		"has_shield":  e.hasShield,
		"is_shielded": e.isShielded,
		// turn
		// TODO: whos_next()?
		// info
		"skill_ready": e.skillReady,
		"element":     e.element,
		// target
		"is_valid":     e.isValid,
		"is_character": e.isCharacter,
		"is_enemy":     e.isEnemy,
		"enemies":      e.enemies,
		"characters":   e.characters,
		"adjacent_to":  e.adjacentTo,
	}
	for name, fn := range funcs {
		env.setBuiltinFunc(name, fn)
	}
}

func (e *Eval) initEnums(env *Env) {
	enums := []map[string]int32{
		model.Property_value,
		model.StatusType_value,
		model.Path_value,
		model.DamageType_value,
		model.TargetType_value,
	}
	for _, enumMap := range enums {
		for name, value := range enumMap {
			env.setv(name, &number{ival: int64(value)})
		}
	}
}

func (e *Eval) initCharNames(env *Env) {
	for _, k := range e.engine.Characters() {
		char, err := e.engine.CharacterInfo(k)
		if err != nil {
			fmt.Println(err)
			return
		}
		env.setv(string(char.Key), &number{ival: int64(k)})
	}
}

// modifier

// has_modifier(target, mod)
func (e *Eval) hasModifier(c *ast.CallExpr, env *Env) (Obj, error) {
	objs, err := e.validateArguments(c.Args, env, typNum, typStr)
	if err != nil {
		return nil, err
	}
	target := key.TargetID(objs[0].(*number).ival)
	modifier := key.Modifier(objs[1].(*strval).str)

	if !e.engine.IsValid(target) {
		return nil, fmt.Errorf("target %d is invalid", target)
	}
	return bton(e.engine.HasModifier(target, modifier)), nil
}

// modifier_count(target, type)
func (e *Eval) modifierCount(c *ast.CallExpr, env *Env) (Obj, error) {
	objs, err := e.validateArguments(c.Args, env, typNum, typNum)
	if err != nil {
		return nil, err
	}
	target := key.TargetID(objs[0].(*number).ival)
	status := model.StatusType(objs[1].(*number).ival)

	if !e.engine.IsValid(target) {
		return nil, fmt.Errorf("target %d is invalid", target)
	}
	return &number{ival: int64(e.engine.ModifierStatusCount(target, status))}, nil
}

// attribute

// ult_ready(char)
func (e *Eval) ultReady(c *ast.CallExpr, env *Env) (Obj, error) {
	objs, err := e.validateArguments(c.Args, env, typNum)
	if err != nil {
		return nil, err
	}
	target := key.TargetID(objs[0].(*number).ival)

	if !e.engine.IsCharacter(target) {
		return nil, fmt.Errorf("target %d is not a character", target)
	}
	return bton(e.engine.EnergyRatio(target) >= 1), nil
}

// skill_points()
func (e *Eval) skillPoints(c *ast.CallExpr, env *Env) (Obj, error) {
	if _, err := e.validateArguments(c.Args, env); err != nil {
		return nil, err
	}

	return &number{ival: int64(e.engine.SP())}, nil
}

// energy(target)
func (e *Eval) energy(c *ast.CallExpr, env *Env) (Obj, error) {
	objs, err := e.validateArguments(c.Args, env, typNum)
	if err != nil {
		return nil, err
	}
	target := key.TargetID(objs[0].(*number).ival)

	if !e.engine.IsValid(target) {
		return nil, fmt.Errorf("target %d is invalid", target)
	}
	return &number{
		fval:    e.engine.Energy(target),
		isFloat: true,
	}, nil
}

// max_energy(target)
func (e *Eval) maxEnergy(c *ast.CallExpr, env *Env) (Obj, error) {
	objs, err := e.validateArguments(c.Args, env, typNum)
	if err != nil {
		return nil, err
	}
	target := key.TargetID(objs[0].(*number).ival)

	if !e.engine.IsValid(target) {
		return nil, fmt.Errorf("target %d is invalid", target)
	}
	return &number{
		fval:    e.engine.MaxEnergy(target),
		isFloat: true,
	}, nil
}

// hp_ratio(target)
func (e *Eval) hpRatio(c *ast.CallExpr, env *Env) (Obj, error) {
	objs, err := e.validateArguments(c.Args, env, typNum)
	if err != nil {
		return nil, err
	}
	target := key.TargetID(objs[0].(*number).ival)

	if !e.engine.IsValid(target) {
		return nil, fmt.Errorf("target %d is invalid", target)
	}
	return &number{
		fval:    e.engine.HPRatio(target),
		isFloat: true,
	}, nil
}

// shield

// has_shield(target, key)
func (e *Eval) hasShield(c *ast.CallExpr, env *Env) (Obj, error) {
	objs, err := e.validateArguments(c.Args, env, typNum, typStr)
	if err != nil {
		return nil, err
	}
	target := key.TargetID(objs[0].(*number).ival)
	key := key.Shield(objs[1].(*strval).str)

	if !e.engine.IsValid(target) {
		return nil, fmt.Errorf("target %d is invalid", target)
	}
	return bton(e.engine.HasShield(target, key)), nil
}

// is_shielded(target)
func (e *Eval) isShielded(c *ast.CallExpr, env *Env) (Obj, error) {
	objs, err := e.validateArguments(c.Args, env, typNum)
	if err != nil {
		return nil, err
	}
	target := key.TargetID(objs[0].(*number).ival)

	if !e.engine.IsValid(target) {
		return nil, fmt.Errorf("target %d is invalid", target)
	}
	return bton(e.engine.IsShielded(target)), nil
}

// info

// skill_ready(char)
func (e *Eval) skillReady(c *ast.CallExpr, env *Env) (Obj, error) {
	objs, err := e.validateArguments(c.Args, env, typNum)
	if err != nil {
		return nil, err
	}
	target := key.TargetID(objs[0].(*number).ival)

	result, err := e.engine.CanUseSkill(target)
	if err != nil {
		return nil, err
	}
	return bton(result), nil
}

// target

// is_valid(targer)
func (e *Eval) isValid(c *ast.CallExpr, env *Env) (Obj, error) {
	objs, err := e.validateArguments(c.Args, env, typNum)
	if err != nil {
		return nil, err
	}
	target := key.TargetID(objs[0].(*number).ival)

	return bton(e.engine.IsValid(target)), nil
}

// is_character(char)
func (e *Eval) isCharacter(c *ast.CallExpr, env *Env) (Obj, error) {
	objs, err := e.validateArguments(c.Args, env, typNum)
	if err != nil {
		return nil, err
	}
	target := key.TargetID(objs[0].(*number).ival)

	return bton(e.engine.IsCharacter(target)), nil
}

// is_enemy(target)
func (e *Eval) isEnemy(c *ast.CallExpr, env *Env) (Obj, error) {
	objs, err := e.validateArguments(c.Args, env, typNum)
	if err != nil {
		return nil, err
	}
	target := key.TargetID(objs[0].(*number).ival)

	return bton(e.engine.IsEnemy(target)), nil
}

// enemies()
func (e *Eval) enemies(c *ast.CallExpr, env *Env) (Obj, error) {
	if _, err := e.validateArguments(c.Args, env); err != nil {
		return nil, err
	}

	enemies := e.engine.Enemies()
	if len(enemies) == 0 {
		return &mapval{}, nil
	}
	result := &mapval{array: make([]Obj, 0, len(enemies))}
	for _, enemy := range enemies {
		result.array = append(result.array, &number{ival: int64(enemy)})
	}
	return result, nil
}

// weakness_broken(target)
func (e *Eval) weaknessBroken(c *ast.CallExpr, env *Env) (Obj, error) {
	objs, err := e.validateArguments(c.Args, env, typNum)
	if err != nil {
		return nil, err
	}
	target := key.TargetID(objs[0].(*number).ival)

	if !e.engine.IsEnemy(target) {
		return nil, fmt.Errorf("target %d is not an enemy", target)
	}
	return bton(e.engine.Stance(target) == 0), nil
}

// has_weakness(target, element)
func (e *Eval) hasWeakness(c *ast.CallExpr, env *Env) (Obj, error) {
	objs, err := e.validateArguments(c.Args, env, typNum, typNum)
	if err != nil {
		return nil, err
	}
	target := key.TargetID(objs[0].(*number).ival)
	element := model.DamageType(objs[1].(*number).ival)

	if !e.engine.IsEnemy(target) {
		return nil, fmt.Errorf("target %d is not an enemy", target)
	}
	stats := e.engine.Stats(target)
	return bton(stats.IsWeakTo(element)), nil
}

// element(char)
func (e *Eval) element(c *ast.CallExpr, env *Env) (Obj, error) {
	objs, err := e.validateArguments(c.Args, env, typNum)
	if err != nil {
		return nil, err
	}
	target := key.TargetID(objs[0].(*number).ival)

	if !e.engine.IsCharacter(target) {
		return nil, fmt.Errorf("target %d is not a character", target)
	}
	charInfo, err := e.engine.CharacterInfo(target)
	if err != nil {
		return nil, err
	}
	return &number{ival: int64(charInfo.Element)}, nil
}

// characters()
func (e *Eval) characters(c *ast.CallExpr, env *Env) (Obj, error) {
	if _, err := e.validateArguments(c.Args, env); err != nil {
		return nil, err
	}

	characters := e.engine.Characters()
	if len(characters) == 0 {
		return &mapval{}, nil
	}
	result := &mapval{array: make([]Obj, 0, len(characters))}
	for _, enemy := range characters {
		result.array = append(result.array, &number{ival: int64(enemy)})
	}
	return result, nil
}

// adjacent_to(target)
func (e *Eval) adjacentTo(c *ast.CallExpr, env *Env) (Obj, error) {
	objs, err := e.validateArguments(c.Args, env, typNum)
	if err != nil {
		return nil, err
	}
	target := key.TargetID(objs[0].(*number).ival)

	if !e.engine.IsValid(target) {
		return nil, fmt.Errorf("target %d is invalid", target)
	}
	targets := e.engine.AdjacentTo(target)
	if len(targets) == 0 {
		return &mapval{}, nil
	}
	result := &mapval{array: make([]Obj, 0, len(targets))}
	for _, enemy := range targets {
		result.array = append(result.array, &number{ival: int64(enemy)})
	}
	return result, nil
}

// stance(target)
func (e *Eval) stance(c *ast.CallExpr, env *Env) (Obj, error) {
	objs, err := e.validateArguments(c.Args, env, typNum)
	if err != nil {
		return nil, err
	}
	target := key.TargetID(objs[0].(*number).ival)

	if !e.engine.IsEnemy(target) {
		return nil, fmt.Errorf("target %d is not an enemy", target)
	}
	return &number{
		fval:    e.engine.Stance(target),
		isFloat: true,
	}, nil
}

// max_stance(target)
func (e *Eval) maxStance(c *ast.CallExpr, env *Env) (Obj, error) {
	objs, err := e.validateArguments(c.Args, env, typNum)
	if err != nil {
		return nil, err
	}
	target := key.TargetID(objs[0].(*number).ival)

	if !e.engine.IsEnemy(target) {
		return nil, fmt.Errorf("target %d is not an enemy", target)
	}
	return &number{
		fval:    e.engine.MaxStance(target),
		isFloat: true,
	}, nil
}

// is_alive(target)
func (e *Eval) isAlive(c *ast.CallExpr, env *Env) (Obj, error) {
	objs, err := e.validateArguments(c.Args, env, typNum)
	if err != nil {
		return nil, err
	}
	target := key.TargetID(objs[0].(*number).ival)

	if !e.engine.IsValid(target) {
		return nil, fmt.Errorf("target %d is invalid", target)
	}
	return bton(e.engine.IsAlive(target)), nil
}
