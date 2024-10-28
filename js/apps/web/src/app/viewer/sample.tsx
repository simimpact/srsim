"use client";
import React from "react";
import { ExecutorContext } from "../exec/provider";
import { model, SimLog } from "@srsim/ts-types";
import { Button } from "ui";
import { ActionStart, GaugeChange, HitEnd, SPChange, TurnStart } from "./log";

type PropType = {
  data: model.SimResult;
};

const defaultLogs = ["TurnStart", "ActionStart", "HitEnd", "SPChange", "GaugeChange"];

export const Sample = ({ data }: PropType) => {
  const [logCat, setLogCat] = React.useState<string[]>(defaultLogs);
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

  const rows = logs.map((e, i) => {
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
      default:
        if (logCat.indexOf(e.name) === -1) {
          return null;
        }
    }

    switch (e.name) {
      case "TurnStart":
        return <TurnStart key={i} names={nameMap} event={e} />;
      case "ActionStart":
        return <ActionStart key={i} names={nameMap} event={e} />;
      case "HitEnd":
        return <HitEnd key={i} names={nameMap} event={e} />;
      case "SPChange":
        return <SPChange key={i} names={nameMap} event={e} />;
      case "GaugeChange":
        return <GaugeChange key={i} names={nameMap} event={e} />;
      default:
        return null;
    }
  });

  return (
    <div className="flex flex-col">
      <pre>{rows}</pre>
    </div>
  );
};
