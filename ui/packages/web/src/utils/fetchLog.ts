import * as event from "@srsim/types/src/event";

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

export type SimLog =
  | {
      name: "HPChange";
      event: event.HPChange;
    }
  | {
      name: "LimboWaitHeal";
      event: event.LimboWaitHeal;
    }
  | {
      name: "EnergyChange";
      event: event.EnergyChange;
    }
  | {
      name: "StanceChange";
      event: event.StanceChange;
    }
  | {
      name: "StanceBreak";
      event: event.StanceBreak;
    }
  | {
      name: "StanceReset";
      event: event.StanceReset;
    }
  | {
      name: "SPChange";
      event: event.SPChange;
    }
  | {
      name: "AttackStart";
      event: event.AttackStart;
    }
  | {
      name: "AttackEnd";
      event: event.AttackEnd;
    }
  | {
      name: "HitStart";
      event: event.HitStart;
    }
  | {
      name: "HitEnd";
      event: event.HitEnd;
    }
  | {
      name: "HealStart";
      event: event.HealStart;
    }
  | {
      name: "HealEnd";
      event: event.HealEnd;
    }
  | {
      name: "ModifierAdded";
      event: event.ModifierAdded;
    }
  | {
      name: "ModifierResisted";
      event: event.ModifierResisted;
    }
  | {
      name: "ModifierRemoved";
      event: event.ModifierRemoved;
    }
  | {
      name: "ModifierExtendedDuration";
      event: event.ModifierExtendedDuration;
    }
  | {
      name: "ModifierExtendedCount";
      event: event.ModifierExtendedCount;
    }
  | {
      name: "ShieldAdded";
      event: event.ShieldAdded;
    }
  | {
      name: "ShieldRemoved";
      event: event.ShieldRemoved;
    }
  | {
      name: "ShieldChange";
      event: event.ShieldChange;
    }
  | {
      name: "Initialize";
      event: event.Initialize;
    }
  | {
      name: "CharactersAdded";
      event: event.CharactersAdded;
    }
  | {
      name: "EnemiesAdded";
      event: event.EnemiesAdded;
    }
  | {
      name: "BattleStart";
      event: event.BattleStart;
    }
  | {
      name: "Phase1Start";
      event: event.Phase1Start;
    }
  | {
      name: "Phase1End";
      event: event.Phase1End;
    }
  | {
      name: "Phase2Start";
      event: event.Phase2Start;
    }
  | {
      name: "Phase2End";
      event: event.Phase2End;
    }
  | {
      name: "TurnStart";
      event: event.TurnStart;
    }
  | {
      name: "TurnEnd";
      event: event.TurnEnd;
    }
  | {
      name: "Termination";
      event: event.Termination;
    }
  | {
      name: "ActionStart";
      event: event.ActionStart;
    }
  | {
      name: "ActionEnd";
      event: event.ActionEnd;
    }
  | {
      name: "InsertStart";
      event: event.InsertStart;
    }
  | {
      name: "InsertEnd";
      event: event.InsertEnd;
    }
  | {
      name: "TargetDeath";
      event: event.TargetDeath;
    }
  | {
      name: "TurnTargetsAdded";
      event: event.TurnTargetsAdded;
    }
  | {
      name: "TurnReset";
      event: event.TurnReset;
    }
  | {
      name: "GaugeChange";
      event: event.GaugeChange;
    }
  | {
      name: "CurrentGaugeCostChange";
      event: event.CurrentGaugeCostChange;
    };
