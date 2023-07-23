export type SkillType = "Normal" | "BPSkill" | "Ultra" | "Talent" | "MazeNormal" | "Maze";

export type SKillEffect =
  | "SingleAttack"
  | "AoEAttack"
  | "MazeAttack"
  | "Blast"
  | "Impair"
  | "Bounce"
  | "Enhance"
  | "Support"
  | "Defence"
  | "Restore";

export type Element = "Fire" | "Ice" | "Physical" | "Wind" | "Lightning" | "Quantum" | "Imaginary";

export interface AvatarSkillConfig {
  attack_type?: SkillType | null;
  bpadd?: Param | null;
  bpneed?: Param | null;
  cool_down: number;
  delay_ratio: Param;
  extra_effect_idlist: number[];
  init_cool_down: number;
  level: number[];
  level_up_cost_list: number[];
  max_level: number;
  param_list: string[][];
  rated_rank_id: number[];
  rated_skill_tree_id: number[];
  show_damage_list: number[];
  show_heal_list: number[];
  show_stance_list: Param[];
  simple_extra_effect_idlist: number[];
  simple_param_list: Param[][];
  simple_skill_desc: string;
  skill_combo_value_delta?: Param | null;
  skill_desc: ParameterizedDescription;
  skill_effect: SKillEffect;
  skill_icon: AssetPath;
  skill_id: number;
  skill_name: string;
  skill_need: string;
  skill_tag: string;
  skill_trigger_key: string;
  skill_type_desc: string;
  spbase?: Param | null;
  spmultiple_ratio: Param;
  stance_damage_type?: Element | null;
  ultra_skill_icon: AssetPath;
}

export type ParameterizedDescription = string[];

export type AssetPath = string;

export interface AvatarRankConfig {
  desc: ParameterizedDescription;
  icon_path: AssetPath;
  name: string;
  param: string[];
  rank: number;
  rank_ability: string[];
  rank_id: number;
  skill_add_level_list: SkillAddLevelList;
  trigger: string;
  unlock_cost: Item[];
}

export type SkillAddLevelList = Record<string, number>;

export interface Item {
  item_id: number;
  item_num: number;
}

export interface AvatarPropertyConfig {
  icon_path: string;
  is_battle_display?: boolean | null;
  is_display?: boolean | null;
  main_relic_filter?: number | null;
  order: number;
  property_classify?: number | null;
  property_name: string;
  property_name_filter: string;
  property_name_relic: string;
  property_name_skill_tree: string;
  property_type: string;
  sub_relic_filter?: number | null;
}

export type Path =
  | "Destruction"
  | "Hunt"
  | "Erudition"
  | "Harmony"
  | "Nihility"
  | "Preservation"
  | "Abundance";

/**
 * metadata for light cones
 */

export interface EquipmentConfig {
  avatar_base_type: Path;
  coin_cost: number;
  equipment_desc: string;
  equipment_id: number;
  equipment_name: string;
  exp_provide: number;
  exp_type: number;
  max_promotion: number;
  max_rank: number;
  rank_up_cost_list: number[];
  rarity: number;
  release: boolean;
  skill_id: number;
}

export interface AvatarConfig {
  avatar_base_type: Path;
  avatar_desc: string;
  avatar_id: number;
  avatar_name: string;
  avatar_votag: string;
  damage_type: Element;
  damage_type_resistance: DamageTypeResistance[];
  rank_idlist: number[];
  rarity: number;
  release: boolean;
  skill_list: number[];
  spneed: number;
}

export interface DamageTypeResistance {
  damage_type: Element;
  value: Param;
}

export interface Param {
  value: number;
}
