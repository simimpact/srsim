export type Event = "TurnReset" | "TurnEnd" | "SPChange";

export interface Log {
  bar: number;
  bazz: number;
  eventIndex: number;
  eventName: Event;
  fooo: string;
  [k: string]: unknown;
}