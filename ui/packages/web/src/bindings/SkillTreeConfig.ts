export type Anchor =
  | "Point01"
  | "Point02"
  | "Point03"
  | "Point04"
  | "Point05"
  | "Point06"
  | "Point07"
  | "Point08"
  | "Point09"
  | "Point10"
  | "Point11"
  | "Point12"
  | "Point13"
  | "Point14"
  | "Point15"
  | "Point16"
  | "Point17"
  | "Point18";

export type AssetPath = string;

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

export interface SkillTreeConfig {
  ability_name: string;
  anchor: Anchor;
  avatar_id: number;
  avatar_promotion_limit: (number | null)[];
  default_unlock: boolean[];
  icon_path: AssetPath;
  level: number[];
  level_up_skill_id: number[];
  material_list: Item[][];
  max_level: number;
  param_list: string[];
  point_desc: ParameterizedDescription;
  point_id: number;
  point_name: string;
  point_trigger_key: string;
  point_type: number;
  pre_point: number[];
  status_add_list: AbilityProperty[];
}

export interface Item {
  item_id: number;
  item_num: number;
}

export interface AbilityProperty {
  property_type: Property;
  value: Param;
}

export interface Param {
  value: number;
}