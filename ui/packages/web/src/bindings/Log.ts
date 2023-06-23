export type Event = "TurnReset" | "TurnEnd" | "SPChange";

export interface Log {
  abc: string;
  bar: number;
  bazz: number;
  children: ChildLog[];
  eventIndex: number;
  eventName: Event;
  fooo: string;
  sss: string;
  [k: string]: unknown;
}

export interface ChildLog {
  abc: string;
  bar: number;
  bazz: number;
  eventIndex: number;
  eventName: Event;
  fooo: string;
  sss: string;
  [k: string]: unknown;
}