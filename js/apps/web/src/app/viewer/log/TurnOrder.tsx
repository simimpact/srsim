import { event } from "@srsim/ts-types";
import React from "react";
type Prop = {
  names: { [key: string]: string };
  turns: event.TurnStatus[];
};
export const TurnOrder = ({ names, turns }: Prop) => {
  const rows = turns.map((e, i) => {
    return (
      <div className="flex flex-row items-center w-full gap-x-2" key={i}>
        <div>{`${names[e.id] ?? e.id}: gauge = ${e.gauge.toFixed(1)}, av = ${e.av.toFixed(0)}`}</div>
      </div>
    );
  });
  return <div className="flex flex-col">{rows}</div>;
};
