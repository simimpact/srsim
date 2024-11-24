import { event } from "@srsim/ts-types";
import React from "react";
import { LogCard } from "./LogCard";
import { FcRight } from "react-icons/fc";

type Prop = {
  names: { [key: string]: string };
  event: {
    name: "ActionStart";
    event: event.ActionStart;
  };
};

export const ActionStart = ({ names, event }: Prop) => {
  let e = event;
  return (
    <LogCard
      className="pl-2"
      msg={`${names[e.event.owner] ?? e.event.owner} used ${e.event.attack_type}`}
      icon={<FcRight />}
    />
  );
};
