import React, { ReactNode } from "react";
import { Card, CardContent, CardHeader, CardTitle } from "@ui/components";

type CardProps = {
  title: string;
  children: ReactNode;
};

export const GraphCard = (props: CardProps) => {
  return (
    <div className="pl-1">
      <Card>
        <CardHeader className="pb-2">
          <CardTitle>{props.title}</CardTitle>
        </CardHeader>
        <CardContent>{props.children}</CardContent>
      </Card>
    </div>
  );
};
