export type ParameterizedDescription = string[];

export type Property =
  | "MaxHP"
  | "Attack"
  | "Defence"
  | "Speed"
  | "CriticalChance"
  | "CriticalDamage"
  | "BreakDamageAddedRatio"
  | "BreakDamageAddedRatioBase"
  | "HealRatio"
  | "MaxSP"
  | "SPRatio"
  | "StatusProbability"
  | "StatusResistance"
  | "CriticalChanceBase"
  | "CriticalDamageBase"
  | "HealRatioBase"
  | "StanceBreakAddedRatio"
  | "SPRatioBase"
  | "StatusProbabilityBase"
  | "StatusResistanceBase"
  | "PhysicalAddedRatio"
  | "PhysicalResistance"
  | "FireAddedRatio"
  | "FireResistance"
  | "IceAddedRatio"
  | "IceResistance"
  | "ThunderAddedRatio"
  | "ThunderResistance"
  | "WindAddedRatio"
  | "WindResistance"
  | "QuantumAddedRatio"
  | "QuantumResistance"
  | "ImaginaryAddedRatio"
  | "ImaginaryResistance"
  | "BaseHP"
  | "HPDelta"
  | "HPAddedRatio"
  | "BaseAttack"
  | "AttackDelta"
  | "AttackAddedRatio"
  | "BaseDefence"
  | "DefenceDelta"
  | "DefenceAddedRatio"
  | "BaseSpeed"
  | "HealTakenRatio"
  | "PhysicalResistanceDelta"
  | "FireResistanceDelta"
  | "IceResistanceDelta"
  | "ThunderResistanceDelta"
  | "WindResistanceDelta"
  | "QuantumResistanceDelta"
  | "ImaginaryResistanceDelta"
  | "SpeedDelta"
  | "SpeedAddedRatio"
  | "AllDamageTypeAddedRatio";

export interface RelicSetSkillConfig {
  ability_name: ParameterizedDescription[];
  ability_param_list: string[][];
  property_list: RelicParam[][];
  require_num: number[];
  set_id: number;
  skill_desc: ParameterizedDescription[];
}

export interface RelicParam {
  param: Param;
  property: Property;
}

export interface Param {
  value: number;
}