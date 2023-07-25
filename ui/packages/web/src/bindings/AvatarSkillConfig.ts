export type SkillType = "Normal" | "BPSkill" | "Ultra" | "Talent" | "MazeNormal" | "Maze";

export type ParameterizedDescription = string[];

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

export type AssetPath = string;

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

export interface Param {
  value: number;
}