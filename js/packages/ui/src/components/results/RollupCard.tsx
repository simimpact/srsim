import React, { memo, ReactNode } from "react";
import { Card, CardContent, CardHeader, CardTitle, Skeleton } from "@ui/components";
import { cn } from "@ui/lib/utils";

type AuxStat = {
  title: string;
  value?: string;
};

type CardProps = {
  title: string;
  color: string;
  value?: string;
  label?: string;
  auxStats?: Array<AuxStat>;
  tooltip?: string | JSX.Element;
  click?: (() => void) | null;
};

const CardTemplate = ({
  title,
  color,
  value,
  label,
  auxStats,
  tooltip,
  click = null,
}: CardProps) => {
  const interactable = click !== null;
  const c = cn(
    "flex flex-auto flex-row items-stretch justify-between pl-2 " +
      `${color} ${interactable ? " hover:cursor-pointer" : ""}`
  );

  return (
    <div className={"flex basis-1/4 flex-auto min-w-fit "}>
      <Card className={c} onClick={() => interactable && value != undefined && click()}>
        <div className="flex flex-col justify-start bg-muted w-full">
          <CardHeader className="pb-2">
            <CardTitle>{title}</CardTitle>
          </CardHeader>
          <CardContent>
            <CardValue value={value} label={label} />
            <CardAux aux={auxStats} />
          </CardContent>
        </div>
      </Card>
    </div>
  );
};

export const RollupCard = memo(CardTemplate);

const CardValue = ({ value, label }: { value?: number | string | null; label?: string }) => {
  const out = value == null ? 1234 : value;
  const valueClass = cn("text-5xl font-bold tabular-nums", { "bp4-skeleton": value == null });

  let lbl: ReactNode;
  if (label != null) {
    lbl = <div className="flex items-start text-base text-gray-400">{label}</div>;
  }

  return (
    <div className="flex flex-row py-2 gap-1 justify-start">
      <div className={valueClass}>{out}</div>
      {lbl}
    </div>
  );
};

const CardAux = ({ aux }: { aux?: Array<AuxStat> }) => {
  if (aux == null) {
    return null;
  }

  return (
    <div className="grid grid-cols-3 gap-x-5 pt-1 justify-start text-sm font-mono min-w-fit">
      {aux.map(e => (
        <AuxItem key={e.title} stat={e} />
      ))}
    </div>
  );
};

const AuxItem = ({ stat }: { stat: AuxStat }) => {
  if (stat.value === undefined) {
    return;
  }

  const cls = cn("font-black text-current text-sm text-bp4-light-gray-500", {
    "bp4-skeleton": stat.value == null,
  });

  return (
    <div className="flex flex-row items-start gap-3">
      <div className="text-gray-400">{stat.title}</div>
      {stat.value === undefined ? (
        <Skeleton />
      ) : (
        <div className="font-black text-current text-sm text-gray-500">{stat.value}</div>
      )}
    </div>
  );
};
