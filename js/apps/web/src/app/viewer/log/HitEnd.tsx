import { event } from "@srsim/ts-types";
import React from "react";
import { JSONTree } from "react-json-tree";
import { FcSupport } from "react-icons/fc";
import {
  Button,
  Card,
  CardContent,
  CardHeader,
  CardTitle,
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "ui";

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
    <Dialog>
      <DialogTrigger asChild>
        <div className="ml-6 cursor-pointer p-1 rounded-sm border-0 bg-muted mt-2 w-fit flex flex-row items-center gap-x-2">
          <FcSupport />
          <div>{`${names[e.event.attacker] ?? e.event.attacker}'s ${e.event.attack_type} hit ${names[e.event.defender] ?? e.event.defender}, dealt ${e.event.total_damage.toLocaleString()} [${e.event.key}]`}</div>
        </div>
      </DialogTrigger>
      <DialogContent className="max-h-[70%] max-w-[70%] overflow-scroll">
        <DialogHeader>
          <DialogTitle>HitEnd Details: {e.event.key}</DialogTitle>
        </DialogHeader>
        <DialogDescription>
          <JSONTree data={e.event.hit} />
        </DialogDescription>
      </DialogContent>
    </Dialog>
  );
};
