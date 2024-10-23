"use client";
import React from "react";
import { ViewerContext } from "./provider";
import { GraphCard, InProgress, OverviewStatsBarGraph } from "ui";
import { RollupGrid } from "./rollup";
import { ExecutorContext } from "../exec/provider";

export default function Viewer() {
  const { state } = React.useContext(ViewerContext);
  const { supplier } = React.useContext(ExecutorContext);
  const exec = supplier();
  if (state.error !== null && state.error !== "") {
    return <pre>{state.error}</pre>;
  }
  if (state.data === null) {
    return <div>No results yet...</div>;
  }
  if (state.data.statistics === undefined) {
    return <div>No stats available yet...</div>;
  }
  return (
    <div className="p-3 flex flex-col min-h-screen">
      {exec.running() ? (
        <InProgress
          val={state.progress ?? 0}
          className="mb-2"
          cancel={() => {
            exec.cancel();
          }}
        />
      ) : null}
      <RollupGrid data={state.data.statistics} />
      <div className="flex flex-col h-full mt-2">
        {state.data.statistics.damage_dealt_by_cycle === undefined ? null : (
          <GraphCard title="Average Damage By Cycle">
            <OverviewStatsBarGraph
              dataKey="avg_dmg_by_cycle"
              data={state.data.statistics.damage_dealt_by_cycle}
            />
          </GraphCard>
        )}
      </div>
    </div>
  );
}
