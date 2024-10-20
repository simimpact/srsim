import { SimConfig, SimResult } from "@srsim/ts-types/src/generated/index.model";

export interface Executor {
  ready(): Promise<boolean>;
  running(): boolean;
  validate(cfg: string): Promise<SimConfig>;
  sample(cfg: string, seed: string): Promise<any>;
  run(
    cfg: string,
    iterations: number,
    updateResult: (result: SimResult, hash: string) => void //has is currently unused
  ): Promise<boolean | void>;
  cancel(): void;
  buildInfo(): { hash: string; date: string };
}
