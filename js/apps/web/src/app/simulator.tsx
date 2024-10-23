"use client";
import React from "react";
import { Button, Editor } from "@ui/components";
import { Executor } from "@srsim/executor";
import { useRouter } from "next/navigation";
import { throttle } from "lodash-es";
import { model } from "@srsim/ts-types";
import { ViewerContext } from "./viewer/provider";
import { ExecutorContext } from "./exec/provider";

export default function Simulator() {
  const { supplier } = React.useContext(ExecutorContext);
  return <SimulatorCore exec={supplier()} />;
}

type SimulatorCoreProps = {
  exec: Executor;
};

const DEFAULT_VIEWER_THROTTLE = 100;
// TODO: this should be an user option
const DEFAULT_ITERS = 1000;
const cfgKey = "user_local_cfg";

const SimulatorCore = ({ exec }: SimulatorCoreProps) => {
  const router = useRouter();
  const [cfg, setCfg] = React.useState<string>("");
  React.useEffect(() => {
    const saved = window.localStorage.getItem(cfgKey);
    if (saved === null) {
      window.localStorage.setItem(cfgKey, "");
      return;
    }
    setCfg(saved);
  }, []);
  React.useEffect(() => {
    localStorage.setItem(cfgKey, cfg);
  }, [cfg]);
  const { dispatch } = React.useContext(ViewerContext);

  const run = () => {
    const updateResult = throttle(
      (res: model.SimResult, hash: string) => {
        console.log("updating result", res);
        dispatch({
          type: "SET_RESULT",
          payload: {
            result: res,
            progress: (100 * (res.statistics?.iterations ?? 0)) / DEFAULT_ITERS,
          },
        });
      },
      DEFAULT_VIEWER_THROTTLE,
      { leading: true, trailing: true }
    );

    exec.run(cfg, DEFAULT_ITERS, updateResult).catch(err => {
      dispatch({
        type: "SET_ERROR",
        payload: {
          error: err,
          config: cfg,
        },
      });
    });

    router.push("/viewer");
  };
  return (
    <div className="m-3">
      <Editor cfg={cfg} onChange={v => setCfg(v)} className="mb-2"></Editor>
      <div className="sticky bottom-0 flex flex-col gap-y-1 z-10">
        <Button variant="secondary" onClick={() => run()}>
          Run
        </Button>
      </div>
    </div>
  );
};
