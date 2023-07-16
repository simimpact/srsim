// Code generated by tygo. DO NOT EDIT.
/* eslint-disable @typescript-eslint/consistent-type-definitions */
/* eslint-disable @typescript-eslint/consistent-indexed-object-style */
/* eslint-disable prettier/prettier */

import { Handler } from "./handler";
import * as info from "./info";

//////////
// source: attribute.go

export type HPChangeEventHandler = Handler<HPChange>;
export interface HPChange {
  target: string;
  old_hp_ratio: number /* float64 */;
  new_hp_ratio: number /* float64 */;
  old_hp: number /* float64 */;
  new_hp: number /* float64 */;
  is_hp_change_by_damage: boolean;
}
export type LimboWaitHealEventHandler = Handler<LimboWaitHeal>;
export interface LimboWaitHeal {
  target: string;
  is_cancelled: boolean;
}
export type EnergyChangeEventHandler = Handler<EnergyChange>;
export interface EnergyChange {
  target: string;
  old_energy: number /* float64 */;
  new_energy: number /* float64 */;
}
export type StanceChangeEventHandler = Handler<StanceChange>;
export interface StanceChange {
  target: string;
  old_stance: number /* float64 */;
  new_stance: number /* float64 */;
}
export type StanceBreakEventHandler = Handler<StanceBreak>;
export interface StanceBreak {
  target: string;
  source: string;
}
export type StanceResetEventHandler = Handler<StanceReset>;
export interface StanceReset {
  target: string;
}
export type SPChangeEventHandler = Handler<SPChange>;
export interface SPChange {
  key: string;
  old_sp: number /* int */;
  new_sp: number /* int */;
}

//////////
// source: combat.go

export type AttackStartEventHandler = Handler<AttackStart>;
export interface AttackStart {
  key: string;
  attacker: string;
  targets: string[];
  attack_type: string;
  damage_type: string;
}
export type AttackEndEventHandler = Handler<AttackEnd>;
export interface AttackEnd {
  key: string;
  attacker: string;
  targets: string[];
  attack_type: string;
  damage_type: string;
}
export type HitStartEventHandler = Handler<HitStart>;
export interface HitStart {
  attacker: string;
  defender: string;
  hit?: info.Hit;
}
export type HitEndEventHandler = Handler<HitEnd>;
export interface HitEnd {
  key: string;
  hit_index: number /* int */;
  attacker: string;
  defender: string;
  attack_type: string;
  damage_type: string;
  base_damage: number /* float64 */;
  bonus_damage: number /* float64 */;
  defence_multiplier: number /* float64 */;
  resistance: number /* float64 */;
  vulnerability: number /* float64 */;
  toughness_multiplier: number /* float64 */;
  fatigue: number /* float64 */;
  all_damage_reduce: number /* float64 */;
  crit_damage: number /* float64 */;
  total_damage: number /* float64 */;
  hp_damage: number /* float64 */;
  shield_damage: number /* float64 */;
  hp_ratio_remaining: number /* float64 */;
  is_crit: boolean;
  use_snapshot: boolean;
}
export type HealStartEventHandler = Handler<HealStart>;
export interface HealStart {
  key: string;
  target?: info.StatsEncoded;
  healer?: info.StatsEncoded;
  base_heal: info.HealMap;
  heal_value: number /* float64 */;
  use_snapshot: boolean;
}
export type HealEndEventHandler = Handler<HealEnd>;
export interface HealEnd {
  key: string;
  target: string;
  healer: string;
  heal_amount: number /* float64 */;
  overflow_heal_amount: number /* float64 */;
  use_snapshot: boolean;
}

//////////
// source: modifier.go

export type ModifierAddedEventHandler = Handler<ModifierAdded>;
export interface ModifierAdded {
  target: string;
  modifier: info.Modifier;
  chance: number /* float64 */;
}
export type ModifierResistedEventHandler = Handler<ModifierResisted>;
export interface ModifierResisted {
  target: string;
  source: string;
  modifier: string;
  chance: number /* float64 */;
  base_chance: number /* float64 */;
  effect_hit_rate: number /* float64 */;
  effect_res: number /* float64 */;
  debuff_res: number /* float64 */;
}
export type ModifierRemovedEventHandler = Handler<ModifierRemoved>;
export interface ModifierRemoved {
  target: string;
  modifier: info.Modifier;
}
export type ModifierExtendedDurationEventHandler = Handler<ModifierExtendedDuration>;
export interface ModifierExtendedDuration {
  target: string;
  modifier: info.Modifier;
  old_value: number /* int */;
  new_value: number /* int */;
}
export type ModifierExtendedCountEventHandler = Handler<ModifierExtendedCount>;
export interface ModifierExtendedCount {
  target: string;
  modifier: info.Modifier;
  old_value: number /* float64 */;
  new_value: number /* float64 */;
}

//////////
// source: shield.go

export type ShieldAddedEventHandler = Handler<ShieldAdded>;
export interface ShieldAdded {
  id: string;
  info: info.Shield;
  shield_health: number /* float64 */;
}
export type ShieldRemovedEventHandler = Handler<ShieldRemoved>;
export interface ShieldRemoved {
  id: string;
  target: string;
}
export type ShieldChangeEventHandler = Handler<ShieldChange>;
export interface ShieldChange {
  id: string;
  target: string;
  old_hp: number /* float64 */;
  new_hp: number /* float64 */;
}

//////////
// source: sim.go

export type InitializeEventHandler = Handler<Initialize>;
export interface Initialize {
  config?: any /* model.SimConfig */;
  seed: number /* int64 */;
}
export type CharactersAddedEventHandler = Handler<CharactersAdded>;
export interface CharactersAdded {
  characters: CharInfo[];
}
export interface CharInfo {
  id: string;
  info?: info.Character;
}
export type EnemiesAddedEventHandler = Handler<EnemiesAdded>;
export interface EnemiesAdded {
  enemies: EnemyInfo[];
}
export interface EnemyInfo {
  id: string;
  info?: info.Enemy;
}
export type BattleStartEventHandler = Handler<BattleStart>;
export interface BattleStart {
  char_info: { [key: string]: info.Character};
  enemy_info: { [key: string]: info.Enemy};
  char_stats: (info.StatsEncoded | undefined)[];
  enemy_stats: (info.StatsEncoded | undefined)[];
  neutral_stats: (info.StatsEncoded | undefined)[];
}
export type Phase1StartEventHandler = Handler<Phase1Start>;
export interface Phase1Start {
}
export type Phase1EndEventHandler = Handler<Phase1End>;
export interface Phase1End {
}
export type Phase2StartEventHandler = Handler<Phase2Start>;
export interface Phase2Start {
}
export type Phase2EndEventHandler = Handler<Phase2End>;
export interface Phase2End {
}
export type TurnStartEventHandler = Handler<TurnStart>;
export interface TurnStart {
  active: string;
  target_type: string;
  delta_av: number /* float64 */;
  total_av: number /* float64 */;
  turn_order: TurnStatus[];
}
export type TurnEndEventHandler = Handler<TurnEnd>;
export interface TurnEnd {
  characters: (info.StatsEncoded | undefined)[];
  enemies: (info.StatsEncoded | undefined)[];
  neutrals: (info.StatsEncoded | undefined)[];
}
export type TerminationEventHandler = Handler<Termination>;
export interface Termination {
  total_av: number /* float64 */;
  reason: string;
}
export type ActionStartEventHandler = Handler<ActionStart>;
export interface ActionStart {
  owner: string;
  attack_type: string;
  is_insert: boolean;
}
export type ActionEndEventHandler = Handler<ActionEnd>;
export interface ActionEnd {
  owner: string;
  targets: { [key: string]: boolean};
  attack_type: string;
  is_insert: boolean;
}
export type InsertStartEventHandler = Handler<InsertStart>;
export interface InsertStart {
  key: string;
  owner: string;
  abort_flags: string[];
  priority: info.InsertPriority;
}
export type InsertEndEventHandler = Handler<InsertEnd>;
export interface InsertEnd {
  key: string;
  owner: string;
  targets: { [key: string]: boolean};
  abort_flags: string[];
  priority: info.InsertPriority;
}
export type TargetDeathEventHandler = Handler<TargetDeath>;
export interface TargetDeath {
  target: string;
  killer: string;
}

//////////
// source: turn.go

export type TurnTargetsAddedEventHandler = Handler<TurnTargetsAdded>;
export interface TurnTargetsAdded {
  targets: string[];
  turn_order: TurnStatus[];
}
export type TurnResetEventHandler = Handler<TurnReset>;
export interface TurnReset {
  reset_target: string;
  gauge_cost: number /* float64 */;
  turn_order: TurnStatus[];
}
export type GaugeChangeEventHandler = Handler<GaugeChange>;
export interface GaugeChange {
  target: string;
  old_gauge: number /* float64 */;
  new_gauge: number /* float64 */;
  turn_order: TurnStatus[];
}
export type CurrentGaugeCostChangeEventHandler = Handler<CurrentGaugeCostChange>;
export interface CurrentGaugeCostChange {
  old_cost: number /* float64 */;
  new_cost: number /* float64 */;
}
export interface TurnStatus {
  id: string;
  gauge: number /* float64 */;
  av: number /* float64 */;
}
