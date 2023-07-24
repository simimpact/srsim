export interface SimConfig {
  settings: { cycle_limit: number };
  characters: CharacterConfig[];
  enemies: EnemyConfig[];
}

export interface CharacterConfig {
  key: string;
  level: number;
  max_level: number;
  eidols: number;
  traces: string[];
  abilities: {
    attack: number;
    skill: number;
    ult: number;
    talent: number;
  };
  light_cone: {
    key: string;
    level: number;
    max_level: number;
    imposition: number;
  };
  relics: Relic[] | undefined;
}

export interface EnemyConfig {
  level: number;
  hp: number;
  toughness: number;
  weakness: ElementConfig[];
}

type ElementConfig = "WIND" | "QUANTUM" | "ICE" | "LIGHTNING" | "FIRE" | "IMAGINARY" | "PHYSICAL";

export interface Relic {
  key: string;
  main_stat: RelicParam<EquimentMainStat>;
  sub_stats: RelicParam<EquimentSubStat>[];
}

interface RelicParam<T extends EquimentMainStat | EquimentSubStat> {
  stat: T;
  amount: number;
}
type EquimentMainStat = string;
type EquimentSubStat = string;
