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