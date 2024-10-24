import React, { ReactNode } from "react";
import { cn } from "ui/src/lib/utils";

type LogCardProp = {
  msg: string;
  icon: ReactNode | null;
  className?: string;
  children?: ReactNode | null;
};

export const LogCard = ({ msg, icon, className, children = null }: LogCardProp) => {
  const c = cn("flex flex-row items-center w-full gap-x-2", className);
  return (
    <div className={c}>
      {icon}
      <div>{msg}</div>
      {children}
    </div>
  );
};
