import { SimConfig, SimResult } from "@srsim/ts-types/src/generated/index.model";
import { Executor } from "./Executor";

export class NullExecutor implements Executor {
  constructor() {}

  public ready(): Promise<boolean> {
    return new Promise(resolve => resolve(false));
  }
  public running(): boolean {
    return false;
  }
  public validate(cfg: string): Promise<SimConfig> {
    return new Promise((_, reject) => reject("null executor - cannot validate"));
  }
  public sample(cfg: string, seed: string): Promise<any> {
    return new Promise((_, reject) => reject("null executor - cannot create sample"));
  }
  public run(
    cfg: string,
    iterations: number,
    updateResult: (result: SimResult, iters: number, hash: string) => void //hash is currently unused; we add iters because currently not in PB
  ): Promise<boolean | void> {
    return new Promise((_, reject) => reject("null executor - cannot run"));
  }
  public cancel(): void {}
  public buildInfo(): { hash: string; date: string } {
    return {
      hash: "",
      date: "",
    };
  }
}
