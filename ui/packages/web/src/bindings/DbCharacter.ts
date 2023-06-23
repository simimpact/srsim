export type Element = "Fire" | "Ice" | "Physical" | "Wind" | "Lightning" | "Quantum" | "Imaginary";

export type AssetPath = string;

export type Path = "Destruction" | "Hunt" | "Erudition" | "Harmony" | "Nihility" | "Preservation" | "Abundance";

export interface DbCharacter {
  element: Element;
  icon: AssetPath;
  id: number;
  max_sp: number;
  name: string;
  path: Path;
  portrait: AssetPath;
  preview: AssetPath;
  ranks: string[];
  rarity: number;
  skill_trees: string[];
  /**
   * skillIds
   */
  skills: string[];
  tag: string;
  [k: string]: unknown;
}