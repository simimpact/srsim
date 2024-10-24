import { event } from "@srsim/ts-types";
import React from "react";
import { LogCard } from "./LogCard";
import { FaExchangeAlt } from "react-icons/fa";
import { TurnOrder } from "./TurnOrder";

type Prop = {
  names: { [key: string]: string };
  event: {
    name: "GaugeChange";
    event: event.GaugeChange;
  };
};

export const GaugeChange = ({ names, event }: Prop) => {
  let e = event.event;
  return (
    <LogCard
      className="pl-10 cursor-pointer"
      msg={`${names[e.target] ?? e.target} gauge changed: ${e.old_gauge.toFixed(1)} to ${e.new_gauge.toFixed(1)} (source ${e.source} [${e.key}]).`}
      icon={<FaExchangeAlt />}
    >
      <TurnOrder names={names} turns={e.turn_order} />
    </LogCard>
  );
};
