"use client";
import React, { ReactNode } from "react";
import { ExecutorContext } from "../exec/provider";
import { model, SimLog } from "@srsim/ts-types";
import { Button } from "ui";
import { FcChargeBattery, FcRight, FcSportsMode, FcSupport } from "react-icons/fc";

type PropType = {
  data: model.SimResult;
};

export const Sample = ({ data }: PropType) => {
  const [logs, setLogs] = React.useState<SimLog[] | null>(null);
  const { supplier } = React.useContext(ExecutorContext);
  const exec = supplier();

  const getSamples = () => {
    const cfg = data.config;
    if (cfg === undefined) {
      console.warn("unexpected cfg not defined");
      return;
    }
    exec.sample(JSON.stringify(cfg), data.debugSeed ?? "0").then((res: SimLog[]) => {
      setLogs(res);
    });
  };

  if (logs === null) {
    return (
      <div>
        <Button onClick={getSamples}>Generate</Button>
      </div>
    );
  }

  let nameMap: { [key: string]: string } = {};

  const rows = logs.map(e => {
    let name = "";
    switch (e.name) {
      case "CharactersAdded":
        //map char names
        e.event.characters.forEach(c => {
          nameMap[c.id] = c.info?.key ?? "unknown";
        });
        return null;
      case "EnemiesAdded":
        //map enemy names
        e.event.enemies.forEach(c => {
          nameMap[c.id] = c.info?.key ?? "unknown";
        });
        return null;
      case "TurnStart":
        return (
          <LogCard
            msg={`${nameMap[e.event.active] ?? e.event.active}'s turn (current av: ${e.event.total_av.toFixed(0)})`}
            icon={<FcSportsMode />}
          />
        );
      case "ActionStart":
        return (
          <LogCard
            className="pl-2"
            msg={`${nameMap[e.event.owner] ?? e.event.owner} used ${e.event.attack_type}`}
            icon={<FcRight />}
          />
        );
      case "HitEnd":
        return (
          <LogCard
            className="pl-6"
            msg={`${nameMap[e.event.attacker] ?? e.event.attacker}'s ${e.event.attack_type} hit ${nameMap[e.event.defender] ?? e.event.defender}, dealt ${e.event.total_damage.toLocaleString()} [${e.event.key}]`}
            icon={<FcSupport />}
          />
        );
      case "SPChange":
        let t = "gained";
        let diff = e.event.new_sp - e.event.old_sp;
        if (diff < 0) {
          t = "lost";
          diff = -1 * diff;
        }
        return (
          <LogCard
            msg={`${t} ${diff} SP (source: ${nameMap[e.event.source] ?? e.event.source} [${e.event.key}]) `}
            icon={<FcChargeBattery />}
            className="pl-2"
          />
        );
      default:
        return null;
    }
  });

  return (
    <div>
      <pre>{rows}</pre>
    </div>
  );
};

type LogCardProp = {
  msg: string;
  icon: ReactNode | null;
  className?: string;
};

const LogCard = ({ msg, icon, className }: LogCardProp) => {
  return (
    <div className={"rounded-sm border-1 bg-muted flex flex-row w-full " + className}>
      {icon}
      <div>{msg}</div>
    </div>
  );
};
