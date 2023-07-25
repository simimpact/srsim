import { useContext, useEffect } from "react";
import { LogViewer } from "@/components/Sim/Logging/LogViewer";
import { SimControlContext } from "@/providers/SimControl";

const Debug = () => {
  const { runSimulation } = useContext(SimControlContext);

  useEffect(() => {
    runSimulation();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  return (
    <div id="dev" className="flex w-full grow flex-col gap-4 self-start">
      <div className="bg-accent text-accent-foreground mx-2 flex grow flex-col rounded-md p-2">
        <LogViewer />
      </div>
    </div>
  );
};

export { Debug };
