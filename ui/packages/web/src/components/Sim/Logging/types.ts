// common metadata for logs
// outside sim
// quote kyle
// This is the stuff where you don't even need to execute the sim, basically metadata like "who are the
//characters", "what are there starting stats/gear", "who are the enemies", "what day was this sim ran",
// "what git commit/version was used", etc
// TODO: link with current m7 db structs from othi's backend
// TODO: db struct codegen
interface Logable {
  loggedDate: string;
  characters: SimCharacter[];
  enemies: SimEnemy[];
  commit: string;
  gear: GearConfig;
}

interface GearConfig {}

interface Log<T extends Logable> {
  eventName: string;
  logType: LogType;
  chunks: T[];
}

const LOG_TYPE = {
  CONFIG: "CONFIG",
  EVENT: "EVENT",
  RESULT: "RESULT",
  ERROR: "ERROR",
} as const;
type LogType = (typeof LOG_TYPE)[keyof typeof LOG_TYPE];

interface SimCharacter {
  name: string;
  id: string | number;
  rarity: number;
  computedAttack: number;
}
interface SimEnemy {
  name: string;
  id: string | number;
  level: number;
  def: number;
}
