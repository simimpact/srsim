import React, { ReactNode } from "react";
import { Card, CardContent, CardHeader, CardTitle } from "@ui/components";

type CardProps = {
  title: string;
  children: ReactNode;
};

export const GraphCard = (props: CardProps) => {
  return (
    <Card className="dark:bg-slate-800 bg-slate-100">
      <CardHeader className="pb-2 ">
        <CardTitle>{props.title}</CardTitle>
      </CardHeader>
      <CardContent>{props.children}</CardContent>
    </Card>
  );
};
