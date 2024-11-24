import { event } from "@srsim/ts-types";
import React from "react";
import { LogCard } from "./LogCard";
import { FcChargeBattery } from "react-icons/fc";

type Prop = {
  names: { [key: string]: string };
  event: {
    name: "SPChange";
    event: event.SPChange;
  };
};

export const SPChange = ({ names, event }: Prop) => {
  let e = event;
  let t = "gained";
  let diff = e.event.new_sp - e.event.old_sp;
  if (diff < 0) {
    t = "lost";
    diff = -1 * diff;
  }
  return (
    <LogCard
      msg={`${t} ${diff} SP (source: ${names[e.event.source] ?? e.event.source} [${e.event.key}]), current ${e.event.new_sp} `}
      icon={<FcChargeBattery />}
      className="pl-2"
    />
  );
};
