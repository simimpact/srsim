export type Event = "TurnReset" | "TurnEnd" | "SPChange";

export interface Log {
  bar: number;
  bazz: number;
  eventName: Event;
  fooo: string;
  [k: string]: unknown;
}