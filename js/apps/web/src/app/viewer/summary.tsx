"use client";
import React from "react";
import { GraphCard, OverviewStatsBarGraph } from "ui";
import { RollupGrid } from "./rollup";
import { model } from "@srsim/ts-types";

type PropType = {
  data: model.SimResult;
};

export const Summary = ({ data }: PropType) => {
  if (data.statistics === undefined) {
    return <div>Waiting for data...</div>;
  }
  return (
    <div className="flex flex-col min-h-screen">
      <RollupGrid data={data.statistics} />
      <div className="flex flex-col h-full mt-2">
        {data.statistics.damage_dealt_by_cycle === undefined ? null : (
          <GraphCard title="Average Damage By Cycle">
            <OverviewStatsBarGraph
              dataKey="avg_dmg_by_cycle"
              data={data.statistics.damage_dealt_by_cycle}
            />
          </GraphCard>
        )}
      </div>
    </div>
  );
};
