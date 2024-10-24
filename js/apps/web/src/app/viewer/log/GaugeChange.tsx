import { event } from "@srsim/ts-types";
import React from "react";
import { LogCard } from "./LogCard";
import { FaExchangeAlt } from "react-icons/fa";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogTrigger,
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from "ui";
import { JSONTree } from "react-json-tree";

type Prop = {
  names: { [key: string]: string };
  event: {
    name: "GaugeChange";
    event: event.GaugeChange;
  };
};

export const GaugeChange = ({ names, event }: Prop) => {
  let e = event.event;
  let clone: event.TurnStatus[] = JSON.parse(JSON.stringify(e.turn_order));

  clone.forEach((e, i) => {
    const n = names[e.id];
    clone[i]!.id = n ?? e.id;
  });

  return (
    <Dialog>
      <DialogTrigger asChild>
        <div>
          <LogCard
            className="pl-10 cursor-pointer"
            msg={`${names[e.target] ?? e.target} gauge changed: ${e.old_gauge.toFixed(1)} to ${e.new_gauge.toFixed(1)} (source ${e.source} [${e.key}]).`}
            icon={<FaExchangeAlt />}
          ></LogCard>
        </div>
      </DialogTrigger>
      <DialogContent className="max-h-[70%] max-w-[70%] overflow-auto">
        <DialogDescription>
          <JSONTree data={clone} />
        </DialogDescription>
      </DialogContent>
    </Dialog>
  );
};
