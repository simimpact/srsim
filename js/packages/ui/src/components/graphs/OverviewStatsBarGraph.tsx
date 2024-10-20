"use client";
import React from "react";
import {
  AnimatedAxis, // any of these can be non-animated equivalents
  AnimatedGrid,
  AnimatedBarSeries,
  XYChart,
  Tooltip,
  Axis,
} from "@visx/xychart";
import { model } from "@srsim/ts-types";
import { OverviewStats } from "@srsim/ts-types/src/generated/index.model";

type OverviewStatsBarGraphsProps = {
  dataKey: string;
  data: model.OverviewStats[];
};

interface OrderedOverviewStats {
  index: number;
  data: OverviewStats;
}

const accessors = {
  xAccessor: (d: OrderedOverviewStats) => d.index,
  yAccessor: (d: OrderedOverviewStats) => d.data.mean,
};

export function OverviewStatsBarGraph(props: OverviewStatsBarGraphsProps) {
  let max = Number.MIN_VALUE;
  const data: OrderedOverviewStats[] = props.data.map((e, i) => {
    if (e.mean! >= max) {
      max = e.mean ?? max;
    }
    return {
      index: i,
      data: e,
    };
  });

  //calculate axies; call it 10 grids starting at 0
  let axisVals: number[] = [];
  for (let i = 0; i < 11; i++) {
    axisVals.push(i * (max / 10));
  }

  return (
    <XYChart height={300} xScale={{ type: "band" }} yScale={{ type: "linear" }}>
      <AnimatedAxis orientation="bottom" />
      <AnimatedGrid columns={false} numTicks={4} />
      <AnimatedBarSeries dataKey={props.dataKey} data={data} {...accessors} radius={1} />
      <Axis orientation="left" tickValues={axisVals}></Axis>
      <Tooltip
        snapTooltipToDatumX
        snapTooltipToDatumY
        showVerticalCrosshair
        showSeriesGlyphs
        renderTooltip={({ tooltipData, colorScale }) => {
          const data = tooltipData!.nearestDatum!.datum as OrderedOverviewStats;
          return <div>{data.data.mean!}</div>;
        }}
      />
    </XYChart>
  );
}
