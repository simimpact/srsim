import React from "react";
import { Alert, AlertDescription, AlertTitle, Button, Progress } from "@ui/components";
import { cn } from "@ui/lib/utils";

type Props = {
  val?: number;
  cancel?: () => void;
  className?: string;
};
export const InProgress = ({ val = 0, cancel, className = "" }: Props) => {
  const c = cn("ml-1", className);
  return (
    <Alert className={c}>
      <AlertTitle className="text-xl mb-2">Simulation in progress...</AlertTitle>
      <AlertDescription>
        <div className="flex flex-col w-full">
          <Progress value={val} />
          {cancel !== undefined ? (
            <Button variant="destructive" className="mt-2 ml-auto" onClick={() => cancel()}>
              Cancel
            </Button>
          ) : null}
        </div>
      </AlertDescription>
    </Alert>
  );
};
