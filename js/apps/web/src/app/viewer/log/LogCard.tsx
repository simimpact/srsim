import React, { ReactNode } from "react";
import { Dialog, DialogTrigger, DialogContent, DialogDescription } from "ui";
import { cn } from "ui/src/lib/utils";

type LogCardProp = {
  msg: string;
  icon: ReactNode | null;
  className?: string;
  children?: ReactNode | null;
};

export const LogCard = ({ msg, icon, className, children = null }: LogCardProp) => {
  const c = cn(
    "flex flex-row items-center w-full gap-x-2",
    className,
    children === null ? "" : "cursor-pointer"
  );

  if (children === null) {
    return (
      <div className={c}>
        {icon}
        <div>{msg}</div>
      </div>
    );
  }

  return (
    <Dialog>
      <DialogTrigger asChild>
        <div className={c}>
          {icon}
          <div>{msg}</div>
        </div>
      </DialogTrigger>
      <DialogContent className="max-h-[70%] max-w-[70%] overflow-auto">
        <DialogDescription>{children}</DialogDescription>
      </DialogContent>
    </Dialog>
  );
};
