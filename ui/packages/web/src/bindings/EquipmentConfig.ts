export type Path = "Destruction" | "Hunt" | "Erudition" | "Harmony" | "Nihility" | "Preservation" | "Abundance";

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