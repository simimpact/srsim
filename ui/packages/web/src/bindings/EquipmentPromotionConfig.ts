export interface EquipmentPromotionConfig {
  base_attack: number[];
  base_attack_add: number[];
  base_defence: number[];
  base_defence_add: number[];
  base_hp: number[];
  base_hpadd: number[];
  equipment_id: number;
  max_level: number[];
  promotion: number[];
  promotion_cost_list: Item[][];
  world_level_require: number[];
}

export interface Item {
  item_id: number;
  item_num: number;
}