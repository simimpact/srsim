export type Path = "Destruction" | "Hunt" | "Erudition" | "Harmony" | "Nihility" | "Preservation" | "Abundance";

export type Element = "Fire" | "Ice" | "Physical" | "Wind" | "Lightning" | "Quantum" | "Imaginary";

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