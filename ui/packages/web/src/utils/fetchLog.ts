import {
  ActionEnd,
  ActionStart,
  BattleStart,
  CurrentGaugeCostChange,
  GaugeChange,
  Initialize,
  InsertEnd,
  InsertStart,
  Termination,
  TurnEnd,
  TurnReset,
  TurnStart,
  TurnTargetsAdded,
} from "@srsim/types/src/event";

const cliApi = (endpoint: string) => `http://localhost:8382${endpoint}`;

interface Result {
  schema_version: { major: string; minor: string };
  sim_version: string;
  modified: boolean;
  debug_seed: string;
  config: {
    setting: {
      cycle_limit: number;
    };
    characters: unknown;
    enemies: unknown;
    gcs1: string;
  };
  statistics: unknown;
}

export type SimLog =
  | {
      name: "Initialize";
      event: Initialize;
    }
  | {
      name: "BattleStart";
      event: BattleStart;
    }
  | {
      name: "TurnStart";
      event: TurnStart;
    }
  | {
      name: "TurnEnd";
      event: TurnEnd;
    }
  | {
      name: "Termination";
      event: Termination;
    }
  | {
      name: "ActionStart";
      event: ActionStart;
    }
  | {
      name: "ActionEnd";
      event: ActionEnd;
    }
  | {
      name: "InsertStart";
      event: InsertStart;
    }
  | {
      name: "InsertEnd";
      event: InsertEnd;
    }
  | {
      name: "TurnTargetsAdded";
      event: TurnTargetsAdded;
    }
  | {
      name: "TurnReset";
      event: TurnReset;
    }
  | {
      name: "GaugeChange";
      event: GaugeChange;
    }
  | {
      name: "CurrentGaugeCostChange";
      event: CurrentGaugeCostChange;
    };

async function fetchLog(): Promise<SimLog[]> {
  const req = await fetch(cliApi("/log"));
  if (req.ok) {
    const asText = await req.text();
    const binding: any[] = asText.split("\n");
    const events: SimLog[] = []; // TODO: TS
    binding.forEach(line => {
      if (line != "") {
        const data = JSON.parse(line as string) as SimLog; // TODO: TS
        events.push(data);
      }
    });
    return Promise.resolve(events);
  } else {
    return Promise.reject("do you have the cli running ?");
  }
}

async function fetchResult() {
  const req = await fetch(cliApi("/result"));
  if (req.ok) {
    const data = JSON.parse(await req.text()) as Result;
    return data;
  } else throw Error("do you have the cli running ?");
}

export { fetchLog, fetchResult };
