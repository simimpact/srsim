export type AssetPath = string;

export interface RelicSetConfig {
  release?: boolean | null;
  set_icon_figure_path: AssetPath;
  set_icon_path: AssetPath;
  set_id: number;
  set_name: string;
  set_skill_list: number[];
}