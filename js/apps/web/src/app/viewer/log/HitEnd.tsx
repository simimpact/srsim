import { event } from "@srsim/ts-types";
import React from "react";
import { JSONTree } from "react-json-tree";
import { FcSupport } from "react-icons/fc";
import { LogCard } from "./LogCard";

type Prop = {
  names: { [key: string]: string };
  event: {
    name: "HitEnd";
    event: event.HitEnd;
  };
};

export const HitEnd = ({ names, event }: Prop) => {
  let e = event;
  return (
    <LogCard
      className="pl-10 cursor-pointer"
      msg={`${names[e.event.attacker] ?? e.event.attacker}'s ${e.event.attack_type} hit ${names[e.event.defender] ?? e.event.defender}, dealt ${e.event.total_damage.toLocaleString()} [${e.event.key}]`}
      icon={<FcSupport />}
    >
      <JSONTree data={e.event.hit} />
    </LogCard>
  );
};
