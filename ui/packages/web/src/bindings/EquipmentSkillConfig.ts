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

export type ParameterizedDescription = string[];

/**
 * skill info for light cones
 */

export interface EquipmentSkillConfig {
  ability_name: string;
  /**
   * merge
   */
  ability_property: AbilityProperty[][];
  /**
   * merge
   */
  param_list: string[][];
  skill_desc: ParameterizedDescription;
  skill_id: number;
  /**
   * merge
   */
  skill_name: string;
}

export interface AbilityProperty {
  property_type: Property;
  value: Param;
}

export interface Param {
  value: number;
}